package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path

		w.Write([]byte(url))
	}
	http.HandleFunc("/", handler)

	fmt.Println("Starting tddo server...")

	swag, _ := ioutil.ReadFile("swag.txt")
	fmt.Println(string(swag))

	http.ListenAndServe(":8000", nil)
}
