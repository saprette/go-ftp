package common

const MAX_CHUNK_SIZE int = 1024

type Chunk struct {
	len     int
	Content []int
}

func NewChunk() *Chunk {
	chunk := new(Chunk)
	chunk.len = 0
	chunk.Content = make([]int, MAX_CHUNK_SIZE)
	return chunk
}

func (c *Chunk) Len() int {
	return c.len
}

func (c *Chunk) SetLen(length int) {
	c.len = length
}

func (c *Chunk) SetContent(content *[]int, size int) {
	c.Content = *content
	c.len = size
}

func (c *Chunk) MaxCapacity() int {
	return MAX_CHUNK_SIZE
}
