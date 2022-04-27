package model

import (
	"context"
	"fmt"
	"time"

	"aruix.net/chj-fyp/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User repository
var UserRepo *mongo.Collection = Collection("user")

type User struct {
	// UserID      string    `json:"userID" bson:"userID"`
	Username    string    `json:"username" bson:"username"`
	Password    []byte    `json:"password" bson:"password"`
	CreatedDate time.Time `json:"createdDate" bson:"created_date"`
}

func NewUser(username string, password string) User {
	return User{username, service.CryptPassword(password), time.Now()}
}

func (u *User) Save() {
	result, err := UserRepo.InsertOne(context.Background(),
		bson.M{
			"username":    u.Username,
			"password":    u.Password,
			"create_date": u.CreatedDate,
		})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("result: %v\n", result)
}
