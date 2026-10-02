package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func GetHtml(w http.ResponseWriter, r *http.Request) {
	html, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandlerForForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Parse Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error reading form: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	stringData, err := service.ParseString(string(fileBytes))
	if err != nil {
		http.Error(w, "Error parse string from file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var nameFile = time.Now().UTC().Format("2006-01-02_15-04-05") + filepath.Ext(header.Filename)
	newFile, err := os.OpenFile(nameFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		http.Error(w, "Error parse string from file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer newFile.Close()

	_, err = newFile.Write([]byte(stringData))
	if err != nil {
		http.Error(w, "Error parse write to new file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newFileBytes, err := io.ReadAll(newFile)
	if err != nil {
		http.Error(w, "Error parse write to new file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(newFileBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
