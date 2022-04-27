package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
)

var Config = loadconfig()

type config struct {
	Port           string `json:"port"`
	AudioDirectory string `json:"audioDirectory"`
	DB             struct {
		MongoURL string `json:"mongoURL"`
		DBName   string `json:"dbName"`
	}
}

// load config file into system
func loadconfig() *config {
	var myconfig config
	btxt, err := ioutil.ReadFile(path.Join("config", "cfg.json"))
	if err != nil {
		log.Fatalf("read config file error: %v", err)
	}
	json.Unmarshal(btxt, &myconfig)
	return &myconfig
}
