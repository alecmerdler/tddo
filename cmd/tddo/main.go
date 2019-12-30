package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// TODO(alecmerdler): Command-line flags for these
const (
	typedocBinary = "../../node_modules/.bin/typedoc"
	typedocOutput = "../../docs"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path[1:]

		// TODO(alecmerdler): Download source code
		os.MkdirAll(url, 700)
		output, _ := os.Create(filepath.Join(typedocOutput, url+".zip"))
		download, _ := http.Get("http:" + url + "/archive.zip")
		// TODO(alecmerdler): Run `typedoc` against source code (and save output somewhere...?)
		io.Copy(output, download.Body)
		fmt.Println("Zip downloaded for " + url)
		// TODO(alecmerdler): Serve output docs as HTML

		// cmd := fmt.Sprintf("")

		w.Write([]byte(url))
	}
	http.HandleFunc("/", handler)

	fmt.Println("Starting tddo server...")

	swag, _ := ioutil.ReadFile("swag.txt")
	fmt.Println(string(swag))

	http.ListenAndServe(":8000", nil)
}
