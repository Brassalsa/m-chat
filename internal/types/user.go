package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Name     string             `json:"name" bson:"name"`
	Password string             `bson:"password"`
}

type Username struct {
	Username string `json:"username" bson:"username"`
}

type Email struct {
	Email string `json:"email" bson:"email"`
}

type Id struct {
	Id primitive.ObjectID `json:"_id" bson:"_id"`
}
