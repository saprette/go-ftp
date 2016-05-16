package common

const MAX_CHUNK_SIZE = 1024

type Chunck struct {
	len     int64
	Content []byte
}

func NewChunck() *Chunck {
	chunck := new(Chunck)
	chunck.len = 0
	chunck.Content = make([]byte, MAX_CHUNK_SIZE)
	return chunck
}

func (c *Chunck) GetLen() int64 {
	return c.len
}
