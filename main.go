package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// array of files types
var removeFiles = []string{"desktop.ini", "Thumbs.db", "LICENSE.txt"}

// array of file extensions
var removeExtensions = []string{".bak", ".localized", ".DS_Store"}

// path of root directory
func main() {
	DirectoryListRecursiveRemove("/Users/alfredomartel/Downloads", removeFiles, removeExtensions)
}

// Directory list recursive that remove especific filesname and file extensions and folders
func DirectoryListRecursiveRemove(dir string, removeFiles []string, removeExtensions []string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("DIRECTORIO: ", err)
	}
	for _, file := range files {
		if file.IsDir() {
			DirectoryListRecursiveRemove(filepath.Join(dir, file.Name()), removeFiles, removeExtensions)
		} else {
			for _, removeFile := range removeFiles {
				fmt.Println(file.Name())
				if file.Name() == removeFile {
					err = os.Remove(filepath.Join(dir, file.Name()))
					if err != nil {
						fmt.Println("Error: ", err, file.Name())
					}
				}
			}
			for _, removeExtension := range removeExtensions {
				fmt.Println(file.Name())
				if filepath.Ext(file.Name()) == removeExtension {

					err = os.Remove(filepath.Join(dir, file.Name()))
					if err != nil {
						fmt.Println("Error r.ext ", err, file.Name())
					}
				}
			}
		}
	}
}
