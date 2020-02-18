package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/deploy-uat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8881", nil))
}
