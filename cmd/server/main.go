package main

import (
	"github.com/antongoncharik/csv-gen-adobe-stock/internal/handlers"
)

func main() {
	handlers.GetKeywords()
	select {}
	// http.HandleFunc("/", handlers.UploadTemplateHandler)
	// http.HandleFunc("/upload", handlers.UploadFileHandler)
	// http.HandleFunc("/table", handlers.TableTamplateHandler)
	// http.HandleFunc("/download", handlers.DownloadCSVHandler)

	// templatesDir := filepath.Join("templates")
	// handlers.LoadTemplates(templatesDir)

	// log.Println("Starting server at :8080")
	// err := http.ListenAndServe(":8080", nil)
	// log.Fatal(err)
}
