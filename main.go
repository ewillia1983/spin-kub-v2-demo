package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Handling %+v\n", r)

	feature := os.Getenv("NEW_FEATURE")

	var content string
	if feature == "true" {
		content = "/content/newFeature.html"
	} else {
		content = "/content/index.html"
	}

	data, err := ioutil.ReadFile(content)

	if err != nil {
		fmt.Printf("Couldn't read index.html: %v", err)
		os.Exit(1)
	}

	io.WriteString(w, string(data[:]))
}

func main() {
	http.HandleFunc("/", index)
	port := ":8000"
	fmt.Printf("Starting to service on port %s\n", port)
	http.ListenAndServe(port, nil)
}
