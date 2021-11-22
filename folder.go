package main

import "net/http"

type Folder struct {
	URL           string `json:"url"`
	FolderPath    string `json:"folder_path"`
	EnableAuth    bool   `json:"enable_auth"`
	Authorization `json:"authorization"`
}

func (server config) folderHandler(folder Folder) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if folder.EnableAuth {
			if !folder.authorizationHandler(w, r) {
				return
			}
		}
		http.StripPrefix(folder.URL, http.FileServer(http.Dir(folder.FolderPath))).ServeHTTP(w, r)
	}
}
