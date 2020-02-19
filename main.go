package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/deploy-uat", func(w http.ResponseWriter, r *http.Request) {
		triggerJob(w, "")
	})

	http.HandleFunc("/deploy-prod", func(w http.ResponseWriter, r *http.Request) {
		triggerJob(w, "")
	})

	log.Fatal(http.ListenAndServe(":8882", nil))
}

func triggerJob(w http.ResponseWriter, url string) {

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Fprintf(w, string(body))
}
