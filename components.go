package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func components(r *mux.Router) {
	for _, v := range []interface{}{"/latest", "/stable", []string{"", "/stable"}} {
		var url, path string
		switch v.(type) {
		case string:
			url = v.(string)
			path = v.(string)
		case []string:
			url = v.([]string)[0]
			path = v.([]string)[1]
		}

		r.Path("/components" + url + "/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./components"+path+"/index.js")
		})
		r.PathPrefix("/components" + url + "/").Handler(http.StripPrefix("/components"+url+"/", http.FileServer(http.Dir("./components"+path+"/"))))
		r.Path("/components" + url).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./components"+path+"/index.js")
		})
	}
}
