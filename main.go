package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// array of files types
var removeFiles = []string{"desktop.ini", "Thumbs.db", "LICENSE.txt"}

func main() {

	var lista, _ = GetWebPageResponse("https://www.toptal.com/developers/gitignore/api/visualstudiocode,go,macos")

	// creating another array with element from lista which first character dont start with #
	var removeExtensions = make([]string, 0)
	for _, element := range lista {

		// check element start with two following characters *. and element lenght > 1 then add to lista2 element without *
		if len(element) > 2 {
			if element[0:2] == "*." {
				removeExtensions = append(removeExtensions, element[1:])
			}
		}
	}

	fmt.Println(removeExtensions)

	DirectoryListRecursiveRemove("/Users/alfredomartel/Downloads", removeFiles, removeExtensions)
}

// Directory list recursive that remove especific filesname and file extensions and folders
func DirectoryListRecursiveRemove(dir string, removeFiles []string, removeExtensions []string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			DirectoryListRecursiveRemove(filepath.Join(dir, file.Name()), removeFiles, removeExtensions)
		} else {
			for _, removeFile := range removeFiles {

				if file.Name() == removeFile {
					fmt.Println("Removing: ", filepath.Join(dir, file.Name()))
					err = os.Remove(filepath.Join(dir, file.Name()))
					if err != nil {
						fmt.Println("Error: ", err, file.Name())
					}
				}
			}
			for _, removeExtension := range removeExtensions {
				if filepath.Ext(file.Name()) == removeExtension {
					fmt.Println("Removing: ", filepath.Join(dir, file.Name()))
					err = os.Remove(filepath.Join(dir, file.Name()))
					if err != nil {
						fmt.Println("Error r.ext ", err, file.Name())
					}
				}
			}
		}
	}
}

// Request web page and return to array with all lines
func GetWebPageResponse(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(body), "\n"), nil
}
