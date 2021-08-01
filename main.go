package main

import (
	"github.com/gorilla/mux"
	"net/http"

	unit "unit.nginx.org/go"
)

var handlers = make(map[string]http.Handler)

// CreateRouter routes declared handlers to their respective paths.
func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	for p, h := range handlers {
		r.PathPrefix(p).Handler(h)
	}

	return r
}

func main() {
	if e := unit.ListenAndServe(":8080", CreateRouter()); e != nil {
		panic(e)
	}
}
