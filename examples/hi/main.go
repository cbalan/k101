package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			log.Fatal(err)
		}

		hostname, err := os.Hostname()
		if err != nil {
			log.Fatal(err)
		}

		count, err := fmt.Fprintf(w, "Hostname: %s\n%s", hostname, string(data))
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Request: %#v, Bytes written: %d\n", r, count)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
