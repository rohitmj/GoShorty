package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rohitmj/goshorty/utils"
)

var urlStore = make(map[string]string)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || !strings.HasPrefix(req.URL, "http") {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortID := utils.GenerateShortID()
	urlStore[shortID] = req.URL

	resp := ShortenResponse{
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortID),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func ResolveURL(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	originalURL, ok := urlStore[path]
	if !ok {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}
