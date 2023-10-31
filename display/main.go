package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"path/filepath"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func readFileAsString(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func readFilesInDirectory(directoryPath string) ([]string, error) {
	fileList, err := os.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}

	var fileContents []string

	for _, fileInfo := range fileList {
		if !fileInfo.IsDir() {
			filePath := directoryPath + string(os.PathSeparator) + fileInfo.Name()
			content, err := readFileAsString(filePath)
			if err != nil {
				return nil, err
			}
			fileContents = append(fileContents, content)
		}
	}

	return fileContents, nil
}

func populateContent()(*widget.List){

	directoryPath := filepath.Join(os.Getenv("HOME"), ".clip-save", "copies")

	data, _ := readFilesInDirectory(directoryPath)

	return widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(string(data[i]))
		})
}


func main() {


	myApp := app.New()
	myWindow := myApp.NewWindow("< clip-save >")

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	myWindow.SetCloseIntercept(func() {
		myWindow.Hide()
	})

	fmt.Println("Press ESC to quit")
	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)
		if event.Key == keyboard.KeyCtrlV {

			fmt.Println("ctr+v pressed...")
			list := populateContent()
			myWindow.SetContent(list)
			myWindow.ShowAndRun()
		}
		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}
