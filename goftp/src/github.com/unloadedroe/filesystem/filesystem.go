package filesystem

import (
	"github.com/unloadedroe/common"
)

type FileSystem struct {
	reader *Reader
}


func NewFileSystem(path *Path) *FileSystem {
	fileSystem := new(FileSystem)
	fileSystem.reader = NewReader(path)
	return fileSystem
}

func (f FileSystem) Read(filename *Path) *common.Chunk {
	return f.reader.ReadChunk()
}

func Write(filename *Path) {

}

func Delete(filename *Path) {

}

func Move(from *Path, to *Path) {

}
