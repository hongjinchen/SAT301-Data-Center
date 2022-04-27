package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"aruix.net/chj-fyp/config"
	"aruix.net/chj-fyp/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func RecordRouter(router *mux.Router) {
	router.HandleFunc("/Post/query/addQuery", addRecord).Methods("POST")
	router.HandleFunc("/Get/query/getPercentageDetectionResults", getNumber).Methods("GET")
	router.HandleFunc("/Get/query/getNumberDetectionByRegion", getNumberByRegion).Methods("GET")
	router.HandleFunc("/Get/query/getAll", getAll).Methods("GET")

}
func getAll(w http.ResponseWriter, req *http.Request) {
	// these bson format will work well but not in record struct
	// var results []bson.D
	// var results []bson.M

	// load data into record struct
	var results []model.Record
	c, err := model.RecordRepo.Find(req.Context(), bson.M{})
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	err = c.All(req.Context(), &results)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	r := Response{"success", results, 200}
	r.ToJson(w)
}

func getNumberByRegion(res http.ResponseWriter, req *http.Request) {
	r := Response{"success", model.MetaData.AddressCount, 200}
	r.ToJson(res)
}

func getNumber(res http.ResponseWriter, req *http.Request) {
	type rr struct {
		TrueNumber  int64 `json:"TrueNumber"`
		FalseNumber int64 `json:"FalseNumber"`
	}
	r := Response{"success", rr{model.MetaData.PositiveResult, model.MetaData.NegativeResult}, 200}
	r.ToJson(res)
}

func addRecord(w http.ResponseWriter, req *http.Request) {
	// get Form data
	err := req.ParseMultipartForm(64 << 20)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	// get addressInfo from multipart form
	addressInfo := req.MultipartForm.Value["addressInfo"][0]
	record := model.NewRecord(addressInfo)

	// Only need filename ext, not name
	filename := req.MultipartForm.File["audioFile"][0].Filename
	ext := filename[strings.LastIndex(filename, "."):]
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	// get audioFile from multipart form
	file, err := req.MultipartForm.File["audioFile"][0].Open()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	// read file in byte array to local file
	cont, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	// add ext name into filename
	record.AudioFile += ext
	record.Save()
	// path := fmt.Sprintf("store/%v%v", record.AudioFile, ext)
	p := path.Join(config.Config.AudioDirectory, record.AudioFile)
	ioutil.WriteFile(p, cont, 0644)
	// log.Printf("[record save]\tfile name is: %v, data size is %v\n", record.AudioFile+ext, len(f))
	r := Response{"upload success", record, 200}
	r.ToJson(w)
}
