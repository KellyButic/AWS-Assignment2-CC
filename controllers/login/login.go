package login

import (
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	viewPage := "views/login.html"

	/*	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Hello World!")
	}*/

	t, _ := template.ParseFiles(viewPage)
	t.ExecuteTemplate(w, "login", "Good")
}
