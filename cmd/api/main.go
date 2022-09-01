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
	"github.com/faridMeli/decompress-and-recover-dump/internal/executors/filterDump"
	"github.com/faridMeli/decompress-and-recover-dump/internal/executors/layoutDump"
	"github.com/faridMeli/decompress-and-recover-dump/internal/executors/pageDump"
	"github.com/faridMeli/decompress-and-recover-dump/internal/executors/shortcutDump"
)

func main() {
	var dump string = "Brick"
	var list []data.DataCompressed
	for _, file := range listDirByReadDir("/Users/farahmed/Downloads/Bricks") {

		log.Println("início - de: " + file)
		if err := mapJSONByDumpType(file, dump, &list); err != nil {
			log.Fatal(err)
		}
	}
	executor := getExecutor(dump, list)
	if executor == nil {
		log.Fatal("Failed")
	} else {
		executor.RevoverDump()
	}

}

func getExecutor(dump string, list []data.DataCompressed) executors.Executor {
	switch dump {
	case "Shortcut":
		return shortcutDump.NewShortcutExecutor(list)
	case "Brick":
		return brickDump.NewBrickExecutor(list)
	case "Layout":
		return layoutDump.NewLayoutExecutor(list)
	case "Page":
		return pageDump.NewPageExecutor(list)
	case "Filter":
		return filterDump.NewFilterExecutor(list)
	default:
		return nil
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

func mapJSONByDumpType(source, dump string, list *[]data.DataCompressed) error {
	// 2. Read the JSON file into the struct array
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	// remember to close the file at the end of the function
	defer sourceFile.Close()

	var decoder *json.Decoder = json.NewDecoder(sourceFile)
	var ranking data.DataCompressed

	for decoder.More() {
		if err := decoder.Decode(&ranking); err != nil {
			return err
		} else {
			*list = append(*list, ranking)
		}
	}

	return nil
}
