package main

import (
	"log"
	"net/http"
)

type File struct {
	URL           string `json:"url"`
	FilePath      string `json:"file_path"`
	EnableAuth    bool   `json:"enable_auth"`
	Authorization `json:"authorization"`
}

func (server config) staticFileHandler(d []byte, file File) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if file.EnableAuth {
			if !file.authorizationHandler(w, r) {
				return
			}
		}
		_, err := w.Write(d)
		if err != nil {
			log.Printf("WARNING: %v\n", err.Error())
		}
	}
}
