package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"https://github.com/Wharem85/indexer-Go/tree/main/Indexer/funcZincSearch"
)

func main() {
	nameDb := os.Args

	if len(nameDb) < 2 {
		panic("Database not specific")
	}

	path, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path += "/" + nameDb[1]

	os.Setenv("nameDb", nameDb[1])

	files, err := ioutil.ReadDir(path)

	if err != nil {
		panic(err)
	}
	if len(files) < 1 {
		panic("Files not found")
	}

	var listFiles []string
	var listDirs []string

	for _, file := range files {
		if file.IsDir() {
			listDirs = append(listDirs, file.Name())
		} else {
			listFiles = append(listFiles, file.Name())
		}
	}

	/* if len(listFiles) >= 1 {
		functions.To
	} */
}
