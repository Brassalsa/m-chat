package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Name     string             `json:"name" bson:"name"`
	Password string             `bson:"password"`
}

type Friend struct {
	User1Id primitive.ObjectID `json:"user_1_id" bson:"user_1_id"`
	User2Id primitive.ObjectID `json:"user_2_id" bson:"user_2_id"`
}

type FriendRequest struct {
	To   primitive.ObjectID `json:"to" bson:"to"`
	From primitive.ObjectID `json:"from" bson:"from"`
}

type Messages struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	ChatId  primitive.ObjectID `json:"chat_id" bson:"chat_id"`
	Content string             `json:"content" bson:"content"`
	UserId  primitive.ObjectID `json:"user_id" bson:"user_id"`
}
