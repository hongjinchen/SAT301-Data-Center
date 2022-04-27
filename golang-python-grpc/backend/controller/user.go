package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"aruix.net/chj-fyp/model"
	"aruix.net/chj-fyp/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type userForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserRouter(router *mux.Router) {
	router.HandleFunc("/users", greet).Methods("GET")
	router.HandleFunc("/users", greet).Methods("POST")
	r := &Response{Message: "ok"}
	router.HandleFunc("/users/register", func(res http.ResponseWriter, req *http.Request) {
		var uf userForm
		d := json.NewDecoder(req.Body)
		err := d.Decode(&uf)
		if err != nil {
			r.Message = "register decoder json not ok"
		}
		fmt.Printf("Request json: %+v\n", uf)
		u := model.NewUser(uf.Username, uf.Password)
		u.Save()
		json.NewEncoder(res).Encode(r)
	}).Methods("POST")

	router.HandleFunc("/users/login", func(res http.ResponseWriter, req *http.Request) {
		var uf userForm
		d := json.NewDecoder(req.Body)
		err := d.Decode(&uf)
		if err != nil {
			r.Message = fmt.Sprint(err)
		}
		// Find user
		var u model.User
		err = model.UserRepo.FindOne(context.Background(), bson.M{"username": uf.Username}).Decode(&u)
		if err != nil {
			r.Message = fmt.Sprint(err)
		}
		// Check password
		if service.CheckPwd(u.Password, []byte(uf.Password)) {
			r.Message = "Login Successfully"
		} else {
			r.Message = "Login Failed, password incorrect"
		}
		json.NewEncoder(res).Encode(r)
	}).Methods("POST")
}
