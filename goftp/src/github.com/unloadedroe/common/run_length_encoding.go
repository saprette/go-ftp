package common

func RunLengthEncoding(chunk *Chunk) *Chunk {
	compressedChunk := NewChunk()

	return compressedChunk
}

func FindBoundsForSubChunk(chunk *Chunk, offset int) int {
	contentLength := int(chunk.GetLen())
	var index = offset + 1
	for first := chunk.Content[offset]; index < contentLength && first == chunk.Content[index]; index++ {
	}
	return index
}
