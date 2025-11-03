package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("contact.html")
		if err != nil {
			http.Error(w, "contact.html not found", http.StatusInternalServerError)
			return
		}

		if r.Method == http.MethodPost {
			name := r.FormValue("name")
			email := r.FormValue("email")
			message := r.FormValue("message")

			fmt.Fprintf(w, "Thank you, %s!\n", name)
			fmt.Fprintf(w, "Email: %s\n", email)
			fmt.Fprintf(w, "Message: %s\n", message)
			return
		}

		tmpl.Execute(w, nil)
	})

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
