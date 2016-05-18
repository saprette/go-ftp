package filesystem

import "testing"

/*
GOPATH=GoFTP/goftp

Run test:
	go test github.com/unloadedroe/filesystem
*/

func TestReadOneChunck(t *testing.T) {
	path := NewPath("./test/file_1.test")
	fileSystem := NewFileSystem(path);
	chunk := fileSystem.Read(path)
	if len(chunk.Content) != 1024 {
		t.Fatal("size does not match!.")
	}
}

func TestReadTwoChuncks(t *testing.T) {
	path := NewPath("./test/file_2.test")
	fileSystem := NewFileSystem(path);
	chunkOne := fileSystem.Read(path)
	chunkTwo := fileSystem.Read(path)
	var totalSizeOfChuncks = len(chunkOne.Content) + len(chunkTwo.Content);
	if  totalSizeOfChuncks != 1024*2 {
		t.Fatal("size does not match!.")
	}
}

func TestReadThreeChuncks(t *testing.T) {
	path := NewPath("./test/file_3.test")
	fileSystem := NewFileSystem(path);
	sum := 0
	for i := 1; i <= 3; i++ {
		sum += len(fileSystem.Read(path).Content)
	}
	if  sum != 1024*3 {
		t.Fatal("size does not match!.")
	}
}

func TestReadFourChuncks(t *testing.T) {
	path := NewPath("./test/file_4.test")
	fileSystem := NewFileSystem(path);
	sum := 0
	for i := 1; i <= 4; i++ {
		sum += len(fileSystem.Read(path).Content)
	}
	if  sum != 1024*4 {
		t.Fatal("size does not match!.")
	}
}

func TestReadFiveChuncks(t *testing.T) {
	path := NewPath("./test/file_5.test")
	fileSystem := NewFileSystem(path);
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += len(fileSystem.Read(path).Content)
	}
	if  sum != 1024*5 {
		t.Fatal("size does not match!.")
	}
}