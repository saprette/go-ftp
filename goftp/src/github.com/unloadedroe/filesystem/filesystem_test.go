package filesystem

import (
	"testing"
)

/*
GOPATH=GoFTP/goftp

Run test:
	go test github.com/unloadedroe/filesystem
*/

func TestReadOneChunck(t *testing.T) {
	path := NewPath("./test/file_1.test")
	if len(Read(path).Content) != 1024 {
		t.Fatal("size does not match!.")
	}
}

func TestReadTwoChuncks(t *testing.T) {
	path := NewPath("./test/file_2.test")

	/*
		TODO: this does not work, add method to check if file is open or not.
	*/

	/*var totalSizeOfChuncks = len(Read(path).Content);
	totalSizeOfChuncks = totalSizeOfChuncks + len(Read(path).Content);
	if  totalSizeOfChuncks != 1024*2 {
		t.Fatal("size does not match!.")
	}*/
}