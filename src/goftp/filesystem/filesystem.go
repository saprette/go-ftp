package filesystem

import (
	"../common"
	"path.go"
)

type FileSystem struct {
	reader reader.Reader
}

var fileSystem *FileSystem

func initFileSystem() {
	if fileSystem == nil {
		fileSystem = new(FileSystem)
	}
}

func (f *FileSystem) Read(filename *Path) *common.Chunck {
	initFileSystem()
	return f.reader.ReadChunck()
}

func (f *FileSystem) Write(filename *Path) {

}

func (f *FileSystem) Delete(filename *Path) {

}

func (f *FileSystem) Move(from *Path, to *Path) {

}
