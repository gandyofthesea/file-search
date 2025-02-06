package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScanForWordMatch(files []string, word string) {
	matchCount := 0
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		reader := bufio.NewScanner(f)

		for reader.Scan() {
			if strings.Contains(reader.Text(), word) {
				matchCount++
				fmt.Println("\033[32m" + file + "\u001B[0m \n")
				break
			}
		}
		_ = f.Close()
	}
	fmt.Printf("\033[32m Found %v matches. \u001B[0m \n", matchCount)
}
