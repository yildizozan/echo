package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, req *http.Request) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "%s\n", string(dump))
	log.Printf("%s\n", string(dump))
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
