package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func errorPage(errorType string, code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	t, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusText(http.StatusInternalServerError))
	}
	data := struct {
		Err  string
		Code int
	}{
		Err:  errorType,
		Code: code,
	}
	err = t.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusText(http.StatusInternalServerError))
	}
}
