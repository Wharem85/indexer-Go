package functzinc

import (
	"bytes"
	"encoging/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// If there's an error, panic
func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// It takes a directory name and a current path as arguments, then it reads the directory and checks if
// there are any files in it. If there are, it converts them to ndjson format and then it checks if
// there are any sub-directories in the current directory. If there are, it calls itself with the name
// of the sub-directory and the current path as arguments
func recursively(nameDir string, curtPath string) {
	curtPath += "/" + nameDir

	files, err := ioutil.ReadDir(curtPath)

	HandleErr(err)

	if len(files) < 1 {
		panic("Files Not found")
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

	if len(listFiles) >= 1 {
		To_ndjson(listFiles, curtPath)
	}

	for _, dir := range listDirs {
		recursively(dir, curtPath)
	}

	if len(listDirs) == 0 {
		return
	}
}

func writeFile(dict1 []byte, dict2 []byte) {
	if _, err := os.Stat(os.Getenv("nameDb") + ".ndjson");
	err == nil {
		f, err := os.OpenFile(os.Getenv("nameDb") + ".ndjson", os.O_APPEND|os.O_WRONLY, 0660)
		HandleErr(err)

		str := string(dict1)
		_, err = fmt.Fprint(f, str, "\n")
		HandleErr(err)

		str2 := string(dict2)
		_, err = fmt.Fprint(f, str2, "\n")
		HandleErr(err)

		defer f.Close()
	}
}

func To_ndjson(namesFiles []string, path string) {
	
}
