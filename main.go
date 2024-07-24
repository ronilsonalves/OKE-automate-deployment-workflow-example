package main

import (
	"log"
	HTTP "net/http"
)

func main() {
	h := func(w HTTP.ResponseWriter, r *HTTP.Request) {
		_, err := w.Write([]byte("Hello, world!"))
		if err != nil {
			return
		}
	}
	HTTP.HandleFunc("/", h)
	log.Println("Server is running on port 8080")
	if err := HTTP.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error in HTTP server: ", err.Error())
	}
}
