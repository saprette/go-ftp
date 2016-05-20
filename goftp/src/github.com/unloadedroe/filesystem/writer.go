package filesystem

import (
	"github.com/unloadedroe/common"
	"bufio"
	"os"
)

// wrapper for golang api's
type Writer struct {
	path   *Path
	err    error
	offset int

	file   *os.File
	writer *bufio.Writer
}

func NewWriter(path *Path) *Writer {
	var err error
	writer := new(Writer)
	writer.path = path
	writer.err = nil
	writer.offset = 0
	writer.file, err = os.Create(path.GetPath())
	if err != nil {
		panic("Could not create file" + path.GetPath())
	}
	writer.writer = bufio.NewWriter(writer.file)
	return writer
}

func (w Writer) Write(chunk *common.Chunk) {
	n := chunk.GetLen()
	if _, err := w.writer.Write(chunk.Content[:n]); err != nil {
		panic(err)
    }
}


func (w Writer) Done() {
	if w.writer != nil {
		if err := w.writer.Flush(); err != nil {
	        panic(err)
	    }
		if err := w.file.Close(); err != nil {
			panic(err)
		}
		w.writer = nil
	}
}
