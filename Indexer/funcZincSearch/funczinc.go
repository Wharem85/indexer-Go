package functzinc

import (
	"bytes"
	"encoding/json"
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
		createJsno(listFiles, curtPath)
	}

	for _, dir := range listDirs {
		recursively(dir, curtPath)
	}

	if len(listDirs) == 0 {
		return
	}
}

// It opens a file, writes to it, and closes it
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

	} else {
		f, err := os.Create(os.Getenv("nameDb") + ".ndjson")
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

// It takes a list of files, and a path, and then it creates a json file with the content of all the
// files
func createJsno(namesFiles []string, path string) {
	splitIndex := strings.Split(path, "/")

	var nameIndex string

	if len(splitIndex) >= 2 {
		nameIndex1 := splitIndex[len(splitIndex)-2]
		nameIndex1 = strings.TrimPrefix(nameIndex1, "_")
		nameIndex = nameIndex1 + "." + splitIndex[len(splitIndex)-1]
	} else {
		nameIndex = splitIndex[len(splitIndex)-1]
		nameIndex = strings.TrimPrefix(nameIndex, "_")
	}

	var cont int64 = 0

	for _, nameFile := range namesFiles {
		myFile, err := os.Stat(path + "/" + nameFile)
		if err != nil {
			fmt.Println("File not exist")
		}
		cont += myFile.Size()
	}


	dict1 := map[string]map[string]string{
		"index": {
			"_index": os.Getenv("nameDb"),
		},
	}

	toJson, err := json.Marshal(dict1)
	HandleErr(err)

	dict2 := make(map[string]string)

	for _, name := range namesFiles {

		content, err := ioutil.ReadFile(path + "/" + name)
		HandleErr(err)

		strContent := string(content)

		dict2[nameIndex+"."+name] = strContent
	}

	toJson2, err := json.Marshal(dict2)
	HandleErr(err)

	writeFile(toJson, toJson2)
}


