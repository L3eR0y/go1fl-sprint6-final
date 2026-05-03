package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, "../index.html")
}

func UploaderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Некорректный POST запрос", http.StatusMethodNotAllowed)
		return
	}

	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("myFile")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)

	dst, err := os.Create(time.Now().UTC().Format("02012006_151405") + filepath.Ext(header.Filename))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	isMorse := service.IsMorse(string(content))

	var resultString string

	if isMorse {
		resultString = morse.ToText(string(content))
	} else {
		resultString = morse.ToMorse(string(content))
	}

	_, err = dst.Write([]byte(string(resultString)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write([]byte(resultString))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
