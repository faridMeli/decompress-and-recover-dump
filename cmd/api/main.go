package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"

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

var wg sync.WaitGroup

func main() {
	// if len(os.Args) != 2 {
	// 	log.Fatal("Invalid Arguments")
	// }
	var dump string = "Brick"
	var lines []data.DataCompressed

	files := listDirByReadDir("/Users/farahmed/Downloads/Layouts")
	uncompressAndClearDirectory(&files)
	list := make(chan data.DataCompressed)

	go readChannel(list, &lines)

	for _, file := range files {
		log.Println("início - de: " + file)
		wg.Add(1)
		go mapJSONByDumpType(file, dump, list)
	}

	wg.Wait()
	close(list)

	executor := getExecutor(dump, lines)

	if executor == nil {
		log.Fatal("Failed")
	} else {
		executor.RevoverDump()
	}

}

func uncompressGzFile(name string) (string, error) {

	if !strings.Contains(name, ".gz") {
		return "", nil
	}
	// Open compressed file
	gzipFile, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Create a gzip reader on top of the file reader
	// Again, it could be any type reader though
	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer gzipReader.Close()

	// Uncompress to a writer. We'll use a file writer
	newName := strings.TrimSuffix(name, ".gz")
	outfileWriter, err := os.Create(newName)

	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer outfileWriter.Close()

	// Copy contents of gzipped file to output file
	_, err = io.Copy(outfileWriter, gzipReader)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	err = os.Remove(name)
	if err != nil {
		log.Fatal(err)
	}
	return newName, nil
}

func uncompressAndClearDirectory(directory *[]string) {
	for _, file := range *directory {
		name, err := uncompressGzFile(file)
		if err != nil {
			log.Fatal("Failed")
		}
		if name != "" {
			*directory = append(*directory, name)
		}
	}
}

func readChannel(list chan data.DataCompressed, lines *[]data.DataCompressed) {
	for data := range list {
		*lines = append(*lines, data)
	}
}

func getExecutor(dump string, lines []data.DataCompressed) executors.Executor {
	switch dump {
	case "Shortcut":
		return shortcutDump.NewShortcutExecutor(lines)
	case "Brick":
		return brickDump.NewBrickExecutor(lines)
	case "Layout":
		return layoutDump.NewLayoutExecutor(lines)
	case "Page":
		return pageDump.NewPageExecutor(lines)
	case "Filter":
		return filterDump.NewFilterExecutor(lines)
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

func mapJSONByDumpType(source, dump string, list chan data.DataCompressed) {
	// 2. Read the JSON file into the struct array
	sourceFile, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the function
	defer sourceFile.Close()

	var decoder *json.Decoder = json.NewDecoder(sourceFile)
	var ranking data.DataCompressed

	for decoder.More() {
		if err := decoder.Decode(&ranking); err != nil {
			log.Fatal(err)
		} else {
			//*list = append(*list, ranking)
			list <- ranking
		}
	}
	wg.Done()
}
