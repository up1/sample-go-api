package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! app platform %s\n", r.URL.Path)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	bindAddr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server listening at %s 🚀\n", bindAddr)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}