package common

import (
	"container/list"
)

func RunLengthEncoding(chunk *Chunk) *Chunk {
	return GetCompressedChunk(GetPairs(chunk))
}

func GetCompressedChunk(list *list.List) *Chunk {
	compressedChunk := NewChunk()
	newChunkContent := make([]int, list.Len()*2)
	var index int = 0
	for e := list.Front(); e != nil; e = e.Next() {
		pair := e.Value.(*Pair)
		newChunkContent[index] = pair.First().(int)
		index++
		newChunkContent[index] = pair.Second().(int)
		index++
	}
	compressedChunk.SetContent(&newChunkContent, index)
	return compressedChunk
}

func GetPairs(chunk *Chunk) *list.List {
	pairList := list.New()
	for oldOffset := 0; oldOffset != chunk.Len(); {
		character := chunk.Content[oldOffset]
		newOffset := FindBoundsForSubChunk(chunk, oldOffset)
		pairList.PushBack(NewPair(newOffset-oldOffset, character))
		oldOffset = newOffset
	}
	return pairList
}

func FindBoundsForSubChunk(chunk *Chunk, offset int) int {
	contentLength := chunk.Len()
	var index = offset + 1
	for first := chunk.Content[offset]; index < contentLength && first == chunk.Content[index]; index++ {
	}
	return index
}
