package handlers

import (
	"net/http"
)

func TableTamplateHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "table.html", nil); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		return
	}
}
