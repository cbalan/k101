package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/the-data", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			log.Fatal(err)
		}
		count, err := fmt.Fprintf(w, "%s", string(data))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Request: %#v, Bytes written: %d\n", r, count)
	})

	log.Fatal(http.ListenAndServe(":31001", nil))
}