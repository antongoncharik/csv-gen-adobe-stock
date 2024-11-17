package main

import (
	"fmt"

	"github.com/antongoncharik/csv-gen-adobe-stock/internal/handlers"
)

func main() {
	str := handlers.GetKeywords()

	fmt.Println("Response:", str)

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
