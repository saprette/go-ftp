package filesystem

import (
	"github.com/unloadedroe/common"
)

type FileSystem struct {
	reader *Reader
}

var fileSystem *FileSystem

func initFileSystem(path *Path) {
	if fileSystem == nil {
		fileSystem = new(FileSystem)
        fileSystem.reader = NewReader(path)
	}
}

func Read(filename *Path) *common.Chunck {
	initFileSystem(filename)
	return fileSystem.reader.ReadChunck()
}

func Write(filename *Path) {

}

func Delete(filename *Path) {

}

func Move(from *Path, to *Path) {

}
