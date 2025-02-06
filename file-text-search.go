package main

import (
	"flag"
)

func main() {
	pathArg, wordArg, extArg := ReadFlags()
	SetExtension(*extArg)
	files := GetJsonFiles(*pathArg)
	ScanForWordMatch(files, *wordArg)
}

func ReadFlags() (*string, *string, *string) {
	pathArg := flag.String("path", GetPath(), "This is the file path to search.")
	wordArg := flag.String("search", "", "This is the search term you are searching for.")
	extensionArg := flag.String("ext", ".json", "This is the extension for the type of text file to search.")
	flag.Parse()

	if len(*pathArg) == 0 {
		panic("No path given")
	}
	if len(*wordArg) == 0 {
		panic("No search term provided")
	}
	if len(*extensionArg) == 0 {
		panic("No file extension provided")
	}

	return pathArg, wordArg, extensionArg
}
