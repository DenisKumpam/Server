package main

import (
	"Server22/validate"
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

	var fvr validate.FormD
	err = json.Unmarshal(b, &fvr)
	if err != nil {
		log.Printf("%s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v\n", fvr)
	fmt.Fprintf(w, "{\"status\": \"%s\", \"errors\": [\"nope\"]}", "successfulle")

	errors := validate.ValidateFormD(&fvr)
	sendErrors(w, errors)

}

func sendErrors(w http.ResponseWriter, errors validate.ValidationErrs) {
	b, err := json.Marshal(errors)
	if err != nil {
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError,
		)
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
