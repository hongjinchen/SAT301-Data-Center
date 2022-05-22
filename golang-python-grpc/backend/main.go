package main

import (
	"log"
	"net/http"

	"aruix.net/chj-fyp/config"
	"aruix.net/chj-fyp/controller"
	"aruix.net/chj-fyp/model"
)

func Init() {
	log.Println("init")
	model.DataBase()
}

func main() {
	log.Println("Starting the application")
	Init()
	// wait for the database to be ready, then go on
	model.Wg.Wait()
	log.Println("Initial ready")
	// add router into http path
	http.Handle("/", controller.Router())
	http.ListenAndServe(":"+config.Config.Port, nil)
}
