package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

type Response struct {
	Candidates    []Candidate   `json:"candidates"`
	UsageMetadata UsageMetadata `json:"usageMetadata"`
	ModelVersion  string        `json:"modelVersion"`
}

type Candidate struct {
	Content      Content `json:"content"`
	FinishReason string  `json:"finishReason"`
	AvgLogprobs  float64 `json:"avgLogprobs"`
}

type Content struct {
	Parts []Part `json:"parts"`
	Role  string `json:"role"`
}

type Part struct {
	Text string `json:"text"`
}

type UsageMetadata struct {
	PromptTokenCount     int `json:"promptTokenCount"`
	CandidatesTokenCount int `json:"candidatesTokenCount"`
	TotalTokenCount      int `json:"totalTokenCount"`
}

type Data struct {
	Name     string
	Title    string
	Keywords string
}

const API_URL = "https://api.openai.com/v1/chat/completions"

const API_KEY = ""

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
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key=AIzaSyBIIj2oEdkCUI1_tE22Ox2hyUOA_hpJJq8"

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": "Write 45 keywords from this title and each keywor consists from one word only: pine forest"},
				},
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return ""
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating request:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return ""
	}

	data := Response{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error parsing JSON:", err)
	}

	re := regexp.MustCompile(`\d+\.\s*([A-Za-z]+)`)
	matches := re.FindAllStringSubmatch(data.Candidates[0].Content.Parts[0].Text, -1)

	var words []string
	for _, match := range matches {
		words = append(words, match[1])
	}

	return strings.Join(words, ",")
}
