package filesystem

import (
	"../common"
	"./path"
	"bufio"
	"io"
	"os"
)

// wrapper for golang api's
type Reader struct {
	path   *path.Path
	err    error
	offset int

	file   *os.File
	reader *bufio.Reader
	done   bool
}

func NewReader(path *path.Path) *Reader {
	reader := new(Reader)
	reader.path = path
	reader.offset = 0
	return reader
}

func (r *Reader) ReadChunck() *common.Chunck {
	chunk := common.NewChunck()
	r.read(chunk)
	return chunk
}

func (r *Reader) read(chunck *common.Chunck) {

	numRead, err := r.reader.Read(chunck.Content)
	if err != nil && err != io.EOF {
		panic(err)
	}
	r.offset = r.offset + numRead
}

func (r *Reader) openFile() (*os.File, error) {
	var file *os.File
	var err error
	if r.file == nil && r.done == false {
		file, err = os.Open(r.path.GetPath())
	} else {
		panic("create a new instance of Reader.")
	}
	return file, err
}

func (r *Reader) closeFile() {
	err := r.file.Close()
	if err != nil {
		panic(err)
	}
	r.reader = nil
	r.done = true
}
