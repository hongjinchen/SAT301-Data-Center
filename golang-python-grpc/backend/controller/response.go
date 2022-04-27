package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data" omitempty:"true"`
	Status  int         `json:"status"`
}

func (r *Response) ToJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	rb, err := json.Marshal(r)
	if err != nil {
		log.Panicf("Error: %s", err)
	}
	w.Write(rb)
}
