package register

import (
	"net/http"
	"text/template"
)

func Register(w http.ResponseWriter, r *http.Request) {
	viewPage := "views/register.html"

	/*	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Hello World!")
	}*/

	t, _ := template.ParseFiles(viewPage)
	t.ExecuteTemplate(w, "register", "Good")
}
