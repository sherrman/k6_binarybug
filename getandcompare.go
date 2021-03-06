package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getBinary(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
	filename := qp.Get("f")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		w.WriteHeader(500)
	}

	w.Write(data)
}

func compareBinary(w http.ResponseWriter, r *http.Request) {
	qp := r.URL.Query()
	filename := qp.Get("f")
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		w.WriteHeader(500)
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body %v", err)
		w.WriteHeader(500)
	}

	fmt.Printf("%v: %d\n", filename, len(body))

	if !bytes.Equal(data, body) {
		ioutil.WriteFile(fmt.Sprintf("%v_compare", filename), body, 0644)
		w.WriteHeader(400)
		w.Write([]byte("Uploaded data doesn't match"))
		return
	}
	w.Write([]byte("Matches!"))
}

func main() {
	http.HandleFunc("/binary", getBinary)
	http.HandleFunc("/compare", compareBinary)
	fmt.Println("Listening on *:9999")
	log.Fatal(http.ListenAndServe(":9999", nil))
}
