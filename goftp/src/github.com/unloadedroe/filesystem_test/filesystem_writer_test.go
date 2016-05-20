package filesystem

import (
	"testing"
	"os"
	"log"
	"strconv"
	"github.com/unloadedroe/filesystem"
)

/*
GOPATH=GoFTP/goftp

Run test:
	go test github.com/unloadedroe/filesystem_test
*/


func getFileSize(filepath *filesystem.Path) int64 {
	file, err := os.Open(filepath.GetPath()) 
	if err != nil {
	    log.Fatal(err)
	}
	fi, err := file.Stat()
	if err != nil {
	    log.Fatal(err)
	}
	return fi.Size()
}


func TestWriteFiveChunk(t *testing.T) {
	var size int = 1024
	inputFilePath := filesystem.NewPath("./test/file_" + strconv.Itoa(size) + ".test")
	outputFilePath := filesystem.NewPath("./test/file_" + strconv.Itoa(size) + "_copy.test")
	fileSystemReadMode := filesystem.NewFileSystemReadMode(inputFilePath);
	fileSystemWriteMode := filesystem.NewFileSystemWriteMode(outputFilePath);
	for i := 0; i < size; i++ {
		chunk := fileSystemReadMode.Read()
		fileSystemWriteMode.Write(chunk)
	}
	fileSystemWriteMode.CloseWriter()
	if getFileSize(outputFilePath) != int64(size)*1024 {
		t.Fatal("size does not match!.")
	}
}