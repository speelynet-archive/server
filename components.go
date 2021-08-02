package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func components(subrouter *mux.Router) {
	for _, v := range []interface{}{"/latest/", "/stable/", []string{"/", "/stable/"}} {
		var url, path string
		switch v.(type) {
		case string:
			url = v.(string)
			path = v.(string)
		case []string:
			url = v.([]string)[0]
			path = v.([]string)[1]
		}

		subrouter.Path(url).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./components"+path+"index.js")
		})
		subrouter.PathPrefix(url).Handler(http.StripPrefix("/components"+url, http.FileServer(http.Dir("./components"+path))))
	}
}
