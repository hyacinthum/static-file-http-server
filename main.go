package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

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
	run()
}

func run() {
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
		http.HandleFunc(file.URL, staticFileHandler(d, file))
	}
	for _, folder := range server.StaticFolder {
		http.HandleFunc(folder.URL, folderHandler(folder))
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
