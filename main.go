package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const viewPath = "views"

func main() {

	templates := map[string]*template.Template{}

	viewPaths, err := os.ReadDir(viewPath)

	if err != nil {
		log.Fatal(err)
	}

	for _, view := range viewPaths {
		filePath := viewPath + "/" + view.Name()
		templates[view.Name()] = template.Must(template.New(view.Name()).ParseFiles(filePath))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		formPage := templates["signup_form.html"]
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

		form := &SignupForm{
			Username: r.Form.Get("username"),
		}

		form.Validate()

		// if there are any errors, re-render the form
		if len(form.Errors) > 0 {
			formPage := templates["signup_form.html"]
			if err := formPage.Execute(w, form); err != nil {
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
