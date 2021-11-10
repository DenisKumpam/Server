package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

var path = filepath.FromSlash("")


func reg(w http.ResponseWriter, req *http.Request)  {
	err := req.ParseForm()

	if err != nil {
		http.Error(w,
			fmt.Sprintf("cannot parse form. Error: %s", err.Error()),
			http.StatusInternalServerError,
		)
		return
	}

	if len(req.Form) == 0 {
		return
	}
	for _, val := range req.Form {
		if len(val) == 0 {
			return
		}
	}

	fmt.Fprint(w, "Activation link has been sent")

	log.Printf(
		"first-name: %s\t email: %s\t password: %s\n",
		req.Form.Get("user_name"),
		req.Form.Get("user_mail"),
		req.Form.Get("password"),
	)

}

func main() {

	fs := http.FileServer(http.Dir("ui/"))

	http.Handle("/", fs)
	http.HandleFunc("/form", reg)

	log.Println("Server is starting on port 63342")
	err := http.ListenAndServe(":63342", nil)
	if err != nil {
		log.Fatal(err)
	}
}
