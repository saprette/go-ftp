package filesystem

import "testing"

/*
GOPATH=GoFTP/goftp

Run test:
	go test github.com/unloadedroe/filesystem
*/


func TestWriteOneChunck(t *testing.T) {
	fileSystemReadMode := NewFileSystemReadMode(NewPath("./test/file_1.test"));
	chunk := fileSystemReadMode.Read()
	if len(chunk.Content) != 1024 {
		t.Fatal("size does not match!.")
	}
	fileSystemWriteMode := NewFileSystemWriteMode(NewPath("./test/file_1_copy.test"));
	fileSystemWriteMode.Write(chunk, true)
}