package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/atotto/clipboard"
)

const maxFiles = 10

func createDirectoryIfNotExists() {
	directoryPath := filepath.Join(os.Getenv("HOME"), ".clip-save", "copies")

	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		err := os.MkdirAll(directoryPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		fmt.Println("Directory created:", directoryPath)
	} else if err != nil {
		fmt.Println("Error checking directory:", err)
		return
	}
}

func main() {
	createDirectoryIfNotExists()

	directoryPath := filepath.Join(os.Getenv("HOME"), ".clip-save", "copies")
	prevSelections := make(map[string]bool)
	nextID := 1

	for {

		clipboardContent, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("Error reading clipboard:", err)
		} else {

			if _, found := prevSelections[clipboardContent]; !found {

				fmt.Println("Clipboard content:", clipboardContent)

				newFileName := fmt.Sprintf("clip_%d.txt", nextID)
				newFilePath := filepath.Join(directoryPath, newFileName)

				if err := os.WriteFile(newFilePath, []byte(clipboardContent), 0644); err != nil {
					fmt.Println("Error writing new file:", err)
				} else {
					fmt.Printf("Written clipboard content to new file: %s\n", newFileName)
					prevSelections[clipboardContent] = true

					nextID = (nextID % maxFiles) + 1
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
}
