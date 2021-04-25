package controllers

import (
	"github.com/KellyButic/AWS-Assignment2-CC/controllers/login"
	"github.com/KellyButic/AWS-Assignment2-CC/controllers/register"
	"github.com/gorilla/mux"
)

func Route(r *mux.Router) {
	r.HandleFunc("/", login.Login)
	r.HandleFunc("register", register.Register)

	//if r.Method() == "POST" {

	//	return
	//}

}
