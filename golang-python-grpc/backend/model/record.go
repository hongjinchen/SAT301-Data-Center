package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// User repository
var RecordRepo *mongo.Collection = Collection("record")

type Record struct {
	// ID        string `json:"id" bson:"_id"`
	AddressInfo string    `json:"addressInfo" bson:"address_info"`
	AudioFile   string    `json:"audioFile" bson:"audio_file"`
	CreatedDate time.Time `json:"createdDate" bson:"created_date"`
	Result      bool      `json:"result" bson:"result"`
}

func NewRecord(addressInfo string) *Record {
	rr := &Record{AddressInfo: addressInfo, AudioFile: uuid.New().String(), CreatedDate: time.Now()}
	log.Printf("New Record object: %+v\n", rr)
	return rr
}

func (r *Record) Save() {
	re, err := RecordRepo.InsertOne(context.Background(), bson.M{
		"address_info": r.AddressInfo,
		"audio_file":   r.AudioFile,
		"created_date": r.CreatedDate,
		"result":       r.Result,
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	// use recore to update metadata and save metaData into database
	UpdateMetaData(context.Background(), r)
	log.Printf("record save result: %v\n", re)
}

func (r *Record) FindRecord(addressInfo string) *Record {
	var record Record
	err := RecordRepo.FindOne(context.Background(), bson.M{"address_info": addressInfo}).Decode(&record)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return &record
}

// query the number of true or false result of records
// func CountResultRecord(result bool) int64 {
// 	number, err := RecordRepo.CountDocuments(context.Background(), bson.M{"result": result})
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	return number
// }
