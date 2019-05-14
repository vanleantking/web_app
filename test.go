package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"

	// "./app/route"
	"strings"

	"./app/controllers"
	"./app/route"
	// "./app/config"
)

const (
	SLASH = "/"
	APP   = "./app/"
)

func main() {
	backend := "backend"
	// frontend := "frontend"

	routes := []string{
		backend + SLASH + "index" + SLASH + "view",
		backend + SLASH + "index" + SLASH + "edit",
		backend + SLASH + "index" + SLASH + "save"}

	request := &http.Request{}
	var response http.ResponseWriter
	router := &route.Router{Request: request, Response: response}

	for _, route := range routes {
		fmt.Println("routesssssssssssssssssssssssssssss, ", route)
		pieces := strings.Split(route, SLASH)
		if len(pieces) > 0 {
			folder := pieces[0]
			if _, err := os.Stat(APP + folder); os.IsNotExist(err) {
				panic(err.Error())
			}
			controller := pieces[1]
			if _, err := os.Stat(APP + folder + SLASH + "controllers" + SLASH + controller + ".go"); os.IsNotExist(err) {
				panic(err.Error())
			}
			if folder == backend {
				files, _ := ioutil.ReadDir(APP + folder + SLASH + "controllers" + SLASH)
				for _, file := range files {
					fmt.Println("file zzzzzzzzzzzzzzzzzzzzzzzzzzzz", file.Name(), APP+folder+SLASH+"controllers"+SLASH)
					switch file.Name() {
					case "index.go":
						fmt.Println("file nameeeeeeeeeeeeeeeeeeeeeee", file.Name())
						http.HandleFunc("/backend/index/view", router.MakeHandler(controllers.ViewHandler))
					}
				}

			}
			// else if folder == frontend {

			// } else {

			// }
			// action := strings.ToTitle(pieces[2]) + "Handler"

		}
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CallFuncByName(myClass interface{}, funcName string, params ...interface{}) (out []reflect.Value, err error) {
	myClassValue := reflect.ValueOf(myClass)
	m := myClassValue.MethodByName(funcName)
	if !m.IsValid() {
		return make([]reflect.Value, 0), fmt.Errorf("Method not found \"%s\"", funcName)
	}
	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}
	out = m.Call(in)
	return
}
