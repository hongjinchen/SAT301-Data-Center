package controller

import (
	"log"
	"net/http"
	"time"

	"aruix.net/chj-fyp/config"
	"github.com/gorilla/mux"
)

func greet(w http.ResponseWriter, r *http.Request) {
	res := Response{Status: 200, Message: "Hello World! Current time" + time.Now().String()}
	res.ToJson(w)
}

func index(w http.ResponseWriter, r *http.Request) {
	message := "Application for Detection of New Coronary Pneumonia Disease by Breath Sound.\nThis is Hongjin Chen's FYP."
	res := Response{Status: 200, Message: message}
	res.ToJson(w)
}

// entry point of router
func Router() *mux.Router {
	router := mux.NewRouter()
	// add middleware
	router.Use(crossOrigin)
	// index page
	router.HandleFunc("/", index).Methods("GET")
	// ping to check if server is running
	router.HandleFunc("/ping", greet).Methods("GET")
	// store static files
	router.PathPrefix("/store/").Handler(http.StripPrefix("/store/", http.FileServer(http.Dir(config.Config.AudioDirectory)))).Methods("GET")
	RecordRouter(router)
	// router.Use(mux.CORSMethodMiddleware(router))
	log.Println("Router created")
	return router
}
