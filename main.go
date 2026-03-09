package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func handler(w http.ResponseWriter, req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatalln(err)
	}

	decoded, err := url.QueryUnescape(string(dump))
	if err != nil {
		decoded = string(dump) // fallback to raw dump
	}

	fmt.Fprintf(w, "%s\n", decoded)
	log.Printf("%s\n", decoded)
}

func main() {
	log.Println("Dont't Ozan's echo server was here")
	log.Println("Echo server starting...")
	log.Println("Server has been started :8080")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
