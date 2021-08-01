package main

import "net/http"

var home = http.StripPrefix("/", http.FileServer(http.Dir("./static/")))
