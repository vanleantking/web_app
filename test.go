package main

import (
	"fmt"
	"log"
	"net/http"

	// "./app/route"
	"./app/controllers"
	"./app/route"
	// "./app/config"
)

func main() {
	request := &http.Request{}
	var response http.ResponseWriter
	router := &route.Router{Request: request, Response: response}
	fmt.Println("uuuuuuuuuuuuuuuuuuuuuuuuuu")
	http.HandleFunc("/view/", router.MakeHandler(controllers.ViewHandler))
	http.HandleFunc("/edit/", router.MakeHandler(controllers.EditHandler))
	http.HandleFunc("/save/", router.MakeHandler(controllers.SaveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
