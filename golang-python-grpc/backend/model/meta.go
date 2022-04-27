package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// singleton pattern, there should only one instance of metaData
type metaData struct {
	// number of positive and negative result
	PositiveResult int64 `json:"positive_result" bson:"positive_result"`
	NegativeResult int64 `json:"negative_result" bson:"negative_result"`

	// address number
	AddressCount map[string]int64 `json:"area_count" bson:"area_count"`
}

// User repository
var MetaDataRepo *mongo.Collection = Collection("meta_data")

var MetaData = getMetadata(context.Background())

func getMetadata(ctx context.Context) *metaData {
	var MetaData *metaData
	MetaDataRepo.FindOne(ctx, bson.M{}).Decode(&MetaData)
	if MetaData == nil {
		MetaData = &metaData{PositiveResult: 0, NegativeResult: 0, AddressCount: make(map[string]int64)}
		MetaDataRepo.InsertOne(ctx, MetaData)
	}
	return MetaData
}

// use record to Update metadata and Save metaData into database
func UpdateMetaData(ctx context.Context, record *Record) {
	operation0 := bson.M{}
	// update the number of positive and negative result
	if record.Result {
		MetaData.PositiveResult++
		operation0["$inc"] = bson.M{"positive_result": 1}
	} else {
		MetaData.NegativeResult++
		operation0["$inc"] = bson.M{"negative_result": 1}
	}
	MetaDataRepo.UpdateOne(ctx, bson.M{}, operation0)

	filter := bson.M{}
	operation1 := bson.M{}
	// update the number of address
	if _, ok := MetaData.AddressCount[record.AddressInfo]; ok {
		MetaData.AddressCount[record.AddressInfo]++
		operation1 = bson.M{"$inc": bson.M{"area_count." + record.AddressInfo: 1}}
	} else {
		MetaData.AddressCount[record.AddressInfo] = 1
		operation1 = bson.M{"$set": bson.M{"area_count." + record.AddressInfo: 1}}
	}
	// update the whole metaData
	MetaDataRepo.UpdateOne(ctx, filter, operation1)
}
