package main

import (
	_ "embed"
	"log"
	"os"
)

//go:embed static/config.json
var sampleConfig []byte

//go:embed static/sample/files/sample1.html
var sampleFile1 []byte

//go:embed static/sample/files/sample2.html
var sampleFile2 []byte

//go:embed static/sample/folder/sample3.html
var sampleFile3 []byte

//go:embed static/sample/folder/sample4.html
var sampleFile4 []byte

func initialization(err error) {
	log.Printf("Loading configuration file error: %v\n", err.Error())
	err = os.WriteFile("config.json", sampleConfig, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to generate sample configuration file: %v\n", err.Error())
	}
	log.Println("The sample configuration file is regenerated.")
	log.Println("Generating sample files...")
	_ = os.Mkdir("sample", os.ModePerm)
	_ = os.Mkdir("sample/files", os.ModePerm)
	_ = os.Mkdir("sample/folder", os.ModePerm)
	err = os.WriteFile("sample/files/sample1.html", sampleFile1, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to generating sample file: %v\n", err.Error())
	}
	err = os.WriteFile("sample/files/sample2.html", sampleFile2, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to generating sample file: %v\n", err.Error())
	}
	err = os.WriteFile("sample/folder/sample3.html", sampleFile3, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to generating sample file: %v\n", err.Error())
	}
	err = os.WriteFile("sample/folder/sample4.html", sampleFile4, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to generating sample file: %v\n", err.Error())
	}
	log.Println("The sample files are generated successfully, please modify the configuration file and restart the program.")
}