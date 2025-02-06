package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var FileExtension string

func GetJsonFiles(path string) []string {
	matches, err := filepath.Glob(path + "/*" + FileExtension)

	if err != nil {
		fmt.Println(err)
	}
	return matches
}

func GetPath() string {
	ex, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	fmt.Println("starting in: " + ex)
	return ex
}

func SetExtension(extension string) {
	FileExtension = extension
}
