package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		formPage := template.Must(template.New("form.template.html").ParseFiles("form.template.html"))
		if err := formPage.Execute(w, nil); err != nil {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		r.ParseForm()
		username := r.Form.Get("username")
		validationError := struct {
			Message []string
		}{
			Message: []string{},
		}

		if username != "terry" {
			validationError.Message = append(validationError.Message, "username must be terry")
		}

		if username == "" {
			validationError.Message = append(validationError.Message, "username cannot be blank")
		}

		if len(username) < 3 {
			validationError.Message = append(validationError.Message, "username must be at least 3 characters")
		}

		if len(username) > 10 {
			validationError.Message = append(validationError.Message, "username must be at most 10 characters")
		}
		// if there are any errors, re-render the form
		if len(validationError.Message) > 0 {
			formPage := template.Must(template.New("form.template.html").ParseFiles("form.template.html"))
			if err := formPage.Execute(w, validationError); err != nil {
				fmt.Println(err)
			}
			// else redirect to the success page :)
		} else {
			http.Redirect(w, r, "/success", http.StatusPermanentRedirect)
		}
	})

	http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hey, you signed up with no validation errors!")
	})
	fmt.Println("Starting server on port 3000")
	http.ListenAndServe(":3000", nil)
}
