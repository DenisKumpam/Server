package main

import (
	"fmt"
	"log"
	"net/http"
)

func reg(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w,
			fmt.Sprintf("cannot parse form. Error: %s", err.Error()),
			http.StatusInternalServerError,
		)
		return
	}
}



func main() {

	fs := http.FileServer(http.Dir("ui/"))

	http.Handle("/", fs)
	http.HandleFunc("/form", reg)

	log.Println("Server is starting on port 7777")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal(err)
	}
}
