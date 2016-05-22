package filesystem

import (
	"bufio"
	"github.com/unloadedroe/common"
	"io"
	"os"
)

// wrapper for golang api's
type Reader struct {
	path   *Path
	err    error
	offset int

	file   *os.File
	reader *bufio.Reader
	done   bool
}

func NewReader(path *Path) *Reader {
	var err error
	reader := new(Reader)
	reader.path = path
	reader.err = nil
	reader.offset = 0
	reader.file, err = os.Open(path.GetPath())
	if err != nil {
		panic("Could not open file" + path.GetPath())
	}
	reader.reader = bufio.NewReader(reader.file)
	reader.done = false
	return reader
}

func (r Reader) ReadChunk() *common.Chunk {
	chunk := common.NewChunk()
	r.read(chunk)
	return chunk
}

func (r Reader) read(chunk *common.Chunk) {
	if chunk != nil && chunk.Content != nil {
		numRead, err := r.reader.Read(chunk.Content)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if numRead == 0 {
			r.closeFile()
		} else {
			chunk.SetLen(numRead)
			r.offset = r.offset + numRead
		}
	}
}

func (r Reader) closeFile() {
	err := r.file.Close()
	if err != nil {
		panic(err)
	}
	r.reader = nil
	r.done = true
}
