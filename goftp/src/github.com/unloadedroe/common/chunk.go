package common

const MAX_CHUNK_SIZE = 1024

type Chunk struct {
	len     int64
	Content []byte
}

func NewChunk() *Chunk {
	chunk := new(Chunk)
	chunk.len = 0
	chunk.Content = make([]byte, MAX_CHUNK_SIZE)
	return chunk
}

func (c *Chunk) GetLen() int64 {
	return c.len
}
