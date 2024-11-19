package handlers

import (
	"encoding/csv"
	"net/http"
)

func DownloadCSVHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=table.csv")
	w.Header().Set("Content-Type", "text/csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	for _, row := range data {
		if row.Name == "" || row.Title == "" || row.Keywords == "" {
			continue
		}

		if err := writer.Write([]string{row.Name, row.Title, row.Keywords}); err != nil {
			http.Error(w, "Failed to write CSV", http.StatusInternalServerError)
			return
		}
	}
}
