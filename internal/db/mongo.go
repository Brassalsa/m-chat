package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	Url         string
	dbName      string
	collections []string

	Ctx      context.Context
	client   *mongo.Client
	database *mongo.Database
}

func NewMongoDb(url, dbName string) *MongoDb {
	return &MongoDb{
		Url:    url,
		dbName: dbName,
	}
}

// connect to db
func (db *MongoDb) Connect(ctx context.Context) error {
	// Set up MongoDB connection options
	clientOptions := options.Client().ApplyURI(db.Url)
	db.Ctx = ctx
	// Create a new MongoDB client
	client, err := mongo.Connect(db.Ctx, clientOptions)
	if err != nil {
		return err
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(db.Ctx, nil)
	if err != nil {
		return err
	}

	db.client = client
	db.database = client.Database(db.dbName)

	return nil
}

// close connection
func (db *MongoDb) Close() error {
	return db.client.Disconnect(db.Ctx)
}

func (db *MongoDb) SetIndex(coll string, field string, unique bool) error {

	if c := db.IsCollection(coll); !c {
		return fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	iMd := mongo.IndexModel{
		Keys: bson.D{
			{Key: field, Value: 1},
		},
		Options: options.Index().SetUnique(unique),
	}
	collection := db.client.Database(db.dbName).Collection(coll)
	if _, err := collection.Indexes().CreateOne(db.Ctx, iMd); err != nil {
		return err
	}
	return nil
}

func (db *MongoDb) AddCollection(coll []string) {
	db.collections = append(db.collections, coll...)
}

// add data to collection
func (db *MongoDb) InsertTo(coll string, payload interface{}) (*mongo.InsertOneResult, error) {
	if c := db.IsCollection(coll); !c {
		return nil, fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	doc, err := toDoc(payload)
	if err != nil {
		return nil, err
	}
	mColl := db.database.Collection(coll)
	return mColl.InsertOne(db.Ctx, doc)
}

// delete data from collection
func (db *MongoDb) DeleteFrom(coll string, filter interface{}) error {
	if c := db.IsCollection(coll); !c {
		return fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	mColl := db.database.Collection(coll)

	_, err := mColl.DeleteOne(db.Ctx, filter)
	return err
}

// find data
func (db *MongoDb) FindOne(coll string, filter interface{}) (*mongo.SingleResult, error) {

	if c := db.IsCollection(coll); !c {
		return nil, fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	doc, err := toDoc(filter)
	if err != nil {
		return nil, err
	}

	mColl := db.database.Collection(coll)
	res := mColl.FindOne(db.Ctx, doc)
	return res, nil
}

func (db *MongoDb) ToId(s string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(s)
}

// update data by id
func (db *MongoDb) Update(coll string, filter interface{}, updateParam interface{}) (*mongo.UpdateResult, error) {
	if c := db.IsCollection(coll); !c {
		return nil, fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	type Id struct {
		Id primitive.ObjectID `bson:"_id"`
	}
	id := new(Id)

	res, err := db.FindOne(coll, filter)
	if err != nil {
		return nil, err
	}

	if err = db.DecodeTo(res, &id); err != nil {
		return nil, err
	}

	uDoc, err := toDoc(updateParam)
	if err != nil {
		return nil, err
	}

	mColl := db.database.Collection(coll)
	uRes, err := mColl.UpdateByID(db.Ctx, id.Id, bson.D{{
		Key:   "$set",
		Value: uDoc,
	}})

	if err != nil {
		return nil, err
	}
	return uRes, nil
}

// get all the data
func (db *MongoDb) GetAll(coll string, results interface{}) error {
	if c := db.IsCollection(coll); !c {
		return fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	mColl := db.database.Collection(coll)
	cursor, err := mColl.Find(db.Ctx, bson.D{})

	if err != nil {
		return err
	}

	defer cursor.Close(db.Ctx)

	if err := cursor.All(db.Ctx, results); err != nil {
		return err
	}

	return nil
}

func (db *MongoDb) Drop() error {
	return db.database.Drop(db.Ctx)
}

func (db *MongoDb) IsCollection(coll string) bool {
	for _, c := range db.collections {
		if c == coll {
			return true
		}
	}

	return false
}

func (db *MongoDb) DecodeTo(res *mongo.SingleResult, ty interface{}) error {
	return res.Decode(ty)
}

func toDoc(v interface{}) (*bson.D, error) {
	data, err := bson.Marshal(v)

	if err != nil {
		return nil, err
	}
	doc := new(bson.D)
	err = bson.Unmarshal(data, &doc)
	return doc, err
}
