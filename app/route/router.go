package route

import (
	"fmt"
	"net/http"
	"regexp"
)

type Router struct {
	Request  *http.Request
	Response http.ResponseWriter
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (router *Router) MakeHandler(fn func(*Router, string)) http.HandlerFunc {
	fmt.Println("tttttttttttttttttttttttttttttttt", router.Response)
	return func(w http.ResponseWriter, r *http.Request) {
		router.Response = w
		router.Request = r
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(router, m[2])
	}
}
