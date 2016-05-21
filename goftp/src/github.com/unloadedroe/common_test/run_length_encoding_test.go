package common_test

import (
	"github.com/unloadedroe/common"
	"testing"
)

func TestFindBoundsForSequenceOneOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.Content = []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	compressedChunk.SetLen(int64(6))

	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 1 {
		t.Fatal("offset does not match!.")
	}
}

func TestFindBoundsForSequenceTwoOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.Content = []byte{'g', 'g', 'l', 'a', 'n', 'g'}
	compressedChunk.SetLen(int64(6))

	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 2 {
		t.Fatal("offset does not match!.")
	}
}

func TestFindBoundsForSequenceThreeOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.Content = []byte{'g', 'g', 'g', 'a', 'n', 'g'}
	compressedChunk.SetLen(int64(6))

	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 3 {
		t.Fatal("offset does not match!.")
	}
}

func TestFindBoundsForSequenceFourOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.Content = []byte{'g', 'g', 'g', 'g', 'n', 'g'}
	compressedChunk.SetLen(int64(6))

	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 4 {
		t.Fatal("offset does not match!.")
	}
}

func TestFindBoundsForSequenceFiveOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.Content = []byte{'g', 'g', 'g', 'g', 'g', 'x'}
	compressedChunk.SetLen(int64(6))

	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 5 {
		t.Fatal("offset does not match!.")
	}
}

func TestFindBoundsForSequenceSixOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.Content = []byte{'g', 'g', 'g', 'g', 'g', 'g'}
	compressedChunk.SetLen(int64(6))

	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 6 {
		t.Fatal("offset does not match!.")
	}
}
