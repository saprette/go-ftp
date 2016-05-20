package filesystem

import (
	"github.com/unloadedroe/common"
)

type FileSystem struct {
	reader *Reader
	writer *Writer
}

func NewFileSystemReadMode(path *Path) *FileSystem {
	fileSystem := new(FileSystem)
	fileSystem.reader = NewReader(path)
	return fileSystem
}

func NewFileSystemWriteMode(path *Path) *FileSystem {
	fileSystem := new(FileSystem)
	fileSystem.writer = NewWriter(path)
	return fileSystem
}

func (f FileSystem) Read() *common.Chunk {
	return f.reader.ReadChunk()
}

func (f FileSystem) Write(chunk *common.Chunk) {
	f.writer.Write(chunk)
}

func (f FileSystem) CloseWriter() {
	f.writer.Done()
}

func Delete(filename *Path) {

}

func Move(from *Path, to *Path) {

}
