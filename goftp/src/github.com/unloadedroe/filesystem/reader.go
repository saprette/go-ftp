package filesystem

import (
	"github.com/unloadedroe/common"
	"bufio"
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

func (r Reader) ReadChunck() *common.Chunck {
	chunk := common.NewChunck()
	r.read(chunk)
	return chunk
}

func (r Reader) read(chunck *common.Chunck) {
	if chunck != nil && chunck.Content != nil{
		numRead, err := r.reader.Read(chunck.Content)
		if err != nil && err != io.EOF {
			panic(err)
		}
	    if numRead == 0 {
	        r.closeFile()
	    } else {
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
