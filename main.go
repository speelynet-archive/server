package main

import (
	"net/http"

	unit "unit.nginx.org/go"
)

func main() {
	if e := unit.ListenAndServe(":8080", http.FileServer(http.Dir("./static/"))); e != nil {
		panic(e)
	}
}
