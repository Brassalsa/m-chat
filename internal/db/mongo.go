package db

import (
	"context"
	"errors"
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
func (m *MongoDb) Connect(ctx context.Context) error {
	// Set up MongoDB connection options
	clientOptions := options.Client().ApplyURI(m.Url)
	m.Ctx = ctx
	// Create a new MongoDB client
	client, err := mongo.Connect(m.Ctx, clientOptions)
	if err != nil {
		return err
	}

	// Ping the MongoDB server to verify the connection
	err = client.Ping(m.Ctx, nil)
	if err != nil {
		return err
	}

	m.client = client
	m.database = client.Database(m.dbName)

	return nil
}

// close connection
func (m *MongoDb) Close() error {
	return m.client.Disconnect(m.Ctx)
}

func (m *MongoDb) SetIndex(coll string, field string, unique bool) error {

	if c := m.IsCollection(coll); !c {
		return fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	iMd := mongo.IndexModel{
		Keys: bson.D{
			{Key: field, Value: 1},
		},
		Options: options.Index().SetUnique(unique),
	}
	collection := m.client.Database(m.dbName).Collection(coll)
	if _, err := collection.Indexes().CreateOne(m.Ctx, iMd); err != nil {
		return err
	}
	return nil
}

func (db *MongoDb) AddCollection(coll []string) {
	db.collections = append(db.collections, coll...)
}

// add data to collection
func (m *MongoDb) InsertTo(coll string, payload interface{}) error {
	if c := m.IsCollection(coll); !c {
		return fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	doc, err := toDoc(payload)
	if err != nil {
		return err
	}
	mColl := m.database.Collection(coll)
	_, err = mColl.InsertOne(m.Ctx, doc)
	return err
}

// delete data from collection
func (m *MongoDb) DeleteFrom(coll string, filter interface{}) error {
	if c := m.IsCollection(coll); !c {
		return fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	mColl := m.database.Collection(coll)

	_, err := mColl.DeleteOne(m.Ctx, filter)
	return err
}

// find data
func (m *MongoDb) FindOne(coll string, filter interface{}, decodeTo interface{}) error {

	if c := m.IsCollection(coll); !c {
		return fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	doc, err := toDoc(filter)
	if err != nil {
		return err
	}

	mColl := m.database.Collection(coll)
	res := mColl.FindOne(m.Ctx, doc)

	if err := res.Decode(decodeTo); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

	return nil
}

func (m *MongoDb) FindMany(coll string, filter interface{}, results interface{}) error {

	if c := m.IsCollection(coll); !c {
		return fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	doc, err := toDoc(filter)
	if err != nil {
		return err
	}

	mColl := m.database.Collection(coll)
	cur, err := mColl.Find(m.Ctx, doc)

	if err != nil {
		return err
	}

	defer cur.Close(m.Ctx)

	if err := cur.All(m.Ctx, results); err != nil {
		return err
	}

	return nil
}

func (m *MongoDb) ToId(s string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(s)
}

// update data by id
func (m *MongoDb) Update(coll string, filter interface{}, updateParam interface{}) error {
	if c := m.IsCollection(coll); !c {
		return fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	type Id struct {
		Id primitive.ObjectID `bson:"_id"`
	}
	id := new(Id)

	err := m.FindOne(coll, filter, &id)
	if err != nil {
		return err
	}

	uDoc, err := toDoc(updateParam)
	if err != nil {
		return err
	}

	mColl := m.database.Collection(coll)
	uRes, err := mColl.UpdateByID(m.Ctx, id.Id, bson.D{{
		Key:   "$set",
		Value: uDoc,
	}})

	if err != nil {
		return err
	}

	if uRes.UpsertedID != nil {
		return errors.New("updateion failed")
	}
	return nil
}

// get all the data
func (m *MongoDb) GetAll(coll string, results interface{}) error {
	if c := m.IsCollection(coll); !c {
		return fmt.Errorf(`"%s" collection doesnot exists`, coll)
	}

	mColl := m.database.Collection(coll)
	cursor, err := mColl.Find(m.Ctx, bson.D{})

	if err != nil {
		return err
	}

	defer cursor.Close(m.Ctx)

	if err := cursor.All(m.Ctx, results); err != nil {
		return err
	}

	return nil
}

func (m *MongoDb) Drop() error {
	return m.database.Drop(m.Ctx)
}

func (m *MongoDb) IsCollection(coll string) bool {
	for _, c := range m.collections {
		if c == coll {
			return true
		}
	}

	return false
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

// queries
func (m *MongoDb) Or(payload ...interface{}) interface{} {
	return bson.D{
		{Key: "$or", Value: payload},
	}
}

func (m *MongoDb) And(payload ...interface{}) interface{} {
	return bson.D{
		{Key: "$and", Value: payload},
	}
}

func (m *MongoDb) Not(payload ...interface{}) interface{} {
	return bson.D{
		{Key: "$not", Value: payload},
	}
}
