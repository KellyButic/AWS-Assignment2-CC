package controllers

import (
	"github.com/KellyButic/AWS-Assignment2-CC/controllers/login"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router) {
	r.HandleFunc("/", login.Login)

	//if r.Method() == "POST" {

	//	return
	//}

}
