package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type formValidationRequest struct {
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func reg(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return
	}
	defer r.Body.Close()

	fvr := &formValidationRequest{}

	err = json.Unmarshal(b, fvr)
	if err != nil {
		log.Printf("%s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("%+v\n", fvr)

}

func main() {

	fs := http.FileServer(http.Dir("ui/"))

	http.Handle("/", fs)
	http.HandleFunc("/validate", reg)

	log.Println("Server is starting on port 7777")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal(err)
	}
}
