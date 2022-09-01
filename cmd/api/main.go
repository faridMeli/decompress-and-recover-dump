package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	//"fmt"
	"log"
	"os"

	"github.com/faridMeli/decompress-and-recover-dump/internal/data"
	"github.com/faridMeli/decompress-and-recover-dump/internal/executors"
	"github.com/faridMeli/decompress-and-recover-dump/internal/executors/brickDump"
	"github.com/faridMeli/decompress-and-recover-dump/internal/executors/shortcutDump"
)

func main() {
	var dump string = "Brick"
	for _, file := range listDirByReadDir("/Users/farahmed/Downloads/Bricks") {

		log.Println("início - de: " + file)
		if executor, err := mapJSONByDumpType(file, dump); err != nil {
			log.Fatal(err)
		} else {
			executor.RollbackDump()
			log.Println("fin - de: " + file)
		}
	}
}

func listDirByReadDir(path string) []string {
	var directory []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() || !strings.Contains(file.Name(), ".json") {
			continue
		} else {
			directory = append(directory, fmt.Sprint(path, "/", file.Name()))
		}
	}

	return directory
}

func mapJSONByDumpType(source, dump string) (executors.Executor, error) {
	// 2. Read the JSON file into the struct array
	sourceFile, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	// remember to close the file at the end of the function
	defer sourceFile.Close()

	var decoder *json.Decoder = json.NewDecoder(sourceFile)
	var ranking data.DataCompressed
	var list []data.DataCompressed

	for decoder.More() {
		if err := decoder.Decode(&ranking); err != nil {
			return nil, err
		} else {
			list = append(list, ranking)
		}
	}

	switch dump {
	case "Shortcut":
		return shortcutDump.NewShortcutExecutor(list), nil
	case "Brick":
		return brickDump.NewBrickExecutor(list), nil
	default:
		return nil, nil
	}
}
