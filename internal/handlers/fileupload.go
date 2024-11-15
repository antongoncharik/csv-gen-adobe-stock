package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

const API_URL = "https://api.openai.com/v1/chat/completions"

const API_KEY = ""

type Data struct {
	Name     string
	Title    string
	Keywords string
}

var data []Data

func UploadTemplateHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "upload.html", nil); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		return
	}
}

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(100 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	photos := r.MultipartForm.File["photos"]
	photoNames := make([]string, len(photos))

	for i, photoHeader := range photos {
		photoNames[i] = photoHeader.Filename
	}

	textFile, _, err := r.FormFile("text")
	if err != nil {
		http.Error(w, "Failed to get text file", http.StatusBadRequest)
		return
	}
	defer textFile.Close()

	textData, err := io.ReadAll(textFile)
	if err != nil {
		http.Error(w, "Failed to read text file", http.StatusInternalServerError)
		return
	}

	titles := splitLines(string(textData))

	for i := 0; i < len(photoNames); i++ {
		data = append(data, Data{Name: photoNames[i], Title: titles[i], Keywords: "yfguhj,yuio,ghyjkl"})
	}

	// var keywords []string
	var keywords string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		keywords = GetKeywords()
	}()
	wg.Wait()

	fmt.Println(keywords)

	if err := tmpl.ExecuteTemplate(w, "table.html", data); err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		return
	}
}

func splitLines(text string) []string {
	return strings.Split(strings.TrimSpace(text), "\n")
}

func GetKeywords() string {
	title := "pine forest"

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": fmt.Sprintf("Give only 10 keywors from this name: %s", title)},
		},
	})
	if err != nil {
		fmt.Println("Ошибка создания JSON:", err)
		return ""
	}

	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return ""
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка выполнения запроса:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа:", err)
		return ""
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return ""
	}
	fmt.Println(result)
	content := ""
	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		content = choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
		fmt.Println("Ключевые слова:", content)
	} else {
		fmt.Println("Не удалось получить ключевые слова.")
	}

	return content
}
