package main

import (
	//"fmt"
	"log"
	"net/http"
)

func main() {
	servemux := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: servemux,
	}

	//servemux.Handle("/assets/logo.png", http.FileServer(http.Dir(".")))
	servemux.Handle("/", http.FileServer(http.Dir(".")))

	log.Fatal(server.ListenAndServe())
	//add error handling heres

}
