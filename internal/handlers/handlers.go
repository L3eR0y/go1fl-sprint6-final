package handlers

import (
	"io"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	indexFile, err := os.Open("../index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer indexFile.Close()

	content, err := io.ReadAll(indexFile)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func UploaderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseMultipartForm(10 << 20)
		file, _, err := r.FormFile("myFile")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		content, err := io.ReadAll(file)

		dst, err := os.Create(time.Now().UTC().Format("02012006_151405") + ".txt")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		isMorse, _ := regexp.MatchString(`^[.\-\s/]+$`, string(content))
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
		w.Write([]byte(resultString))

		return
	}

	http.Error(w, "Некорректный POST запрос", http.StatusBadRequest)
}
