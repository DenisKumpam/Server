package main

import (
	"Server22/validate"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func reg(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return
	}
	defer r.Body.Close()

	var fvr validate.Form
	err = json.Unmarshal(b, &fvr)
	if err != nil {
		log.Printf("%s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v\n", fvr)

	errors := fvr.ValidateForm()

	sendErrors(w, errors)

}

func sendErrors(w http.ResponseWriter, errors []error) {
	b, err := json.Marshal(errors)
	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.Write(b)
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
