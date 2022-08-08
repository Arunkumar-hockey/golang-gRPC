package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-gRPC/server/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type User struct {
	ID       string `bson:"id"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	PhoneNo  string `bson:"phoneno"`
	Password string `bson:"password"`
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func (user *User) Create() (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	emailCount, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if emailCount > 0 {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("email ID already registered"),
		)
	}

	res, err := userCollection.InsertOne(ctx, user)
	fmt.Println("Result is", res.InsertedID)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error on create user"),
		)
	}
	return user, nil
}
