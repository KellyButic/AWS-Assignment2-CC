package main

import (
	//"github.com/KellyButic/AWS-Assignment2-CC/tree/main/controllers"
	"github.com/KellyButic/AWS-Assignment2-CC/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	controllers.Route(router)
}
