package main

import (
	"net/http"
)

func main() {
	port := "10010"
	r := NewRouter()
	println("Starting Golang structure demo project at " + port)
	panic(http.ListenAndServe(":"+port, r))
}
