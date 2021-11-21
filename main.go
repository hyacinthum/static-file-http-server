package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type File struct {
	URL      string `json:"url"`
	FilePath string `json:"file_path"`
}

type Folder struct {
	URL        string `json:"url"`
	FolderPath string `json:"folder_path"`
}

type Config struct {
	Port           int      `json:"port"`
	EnableHTTPS    bool     `json:"enable_https"`
	PublicKeyPath  string   `json:"public_key_path"`
	PrivateKeyPath string   `json:"private_key_path"`
	StaticFiles    []File   `json:"static_files"`
	StaticFolder   []Folder `json:"static_folder"`
}

var server Config

func main() {
	d, err := os.ReadFile("config.json")
	if err != nil {
		initialization(err)
		os.Exit(0)
	}
	err = json.Unmarshal(d, &server)
	if err != nil {
		log.Println("Configuration file parsing error.")
		log.Fatalln("Delete config.json to regenerate the default configuration file.")
	}
	for _, file := range server.StaticFiles {
		d, err = os.ReadFile(file.FilePath)
		if err != nil {
			d = make([]byte, 0)
		}
		http.HandleFunc(file.URL, func(d []byte) func(http.ResponseWriter, *http.Request) {
			return func(w http.ResponseWriter, r *http.Request) {
				_, err = w.Write(d)
				if err != nil {
					log.Printf("WARNING: %v\n", err.Error())
				}
			}
		}(d))
	}
	for _, folder := range server.StaticFolder {
		fmt.Println(folder.URL, folder.FolderPath)
		http.Handle(folder.URL, http.StripPrefix(folder.URL, http.FileServer(http.Dir(folder.FolderPath))))
	}
	log.Printf("Server is Listening: 0.0.0.0:%v\n", server.Port)
	if !server.EnableHTTPS {
		err = http.ListenAndServe(fmt.Sprintf(":%v", server.Port), nil)
	} else {
		err = http.ListenAndServeTLS(fmt.Sprintf(":%v", server.Port), server.PublicKeyPath, server.PrivateKeyPath, nil)
	}
	if err != nil {
		log.Fatalf("ERROR: %v\n", err.Error())
	}
}
