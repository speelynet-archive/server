package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
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

		r.PathPrefix("/components" + url).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")

			if regexp.MustCompile("components" + url + "/?$").MatchString(r.URL.Path) {
				if r.URL.Path[len(r.URL.Path)-1] != '/' {
					http.Redirect(w, r, r.URL.Host+"/components"+url+"/", 302)
				} else {
					http.ServeFile(w, r, "./components"+path+"/index.js")
				}
			} else {
				http.StripPrefix("/components"+url, http.FileServer(http.Dir("./components"+path))).ServeHTTP(w, r)
			}
		})
	}
}
