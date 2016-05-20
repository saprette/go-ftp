package filesystem

import (
	"testing"
	"strconv"
	"github.com/unloadedroe/filesystem"
)

/*
GOPATH=GoFTP/goftp

Run test:
	go test github.com/unloadedroe/filesystem_test
*/

func TestReadFiveChuncks(t *testing.T) {
	var size int = 1024
	path := filesystem.NewPath("./test/file_" + strconv.Itoa(size) + ".test")
	fileSystem := filesystem.NewFileSystemReadMode(path);
	var sum int64 = 0
	for i := 1; i <= size; i++ {
		sum += fileSystem.Read().GetLen()
	}
	if  sum != int64(size)*1024 {
		t.Fatal("size does not match!.")
	}
}