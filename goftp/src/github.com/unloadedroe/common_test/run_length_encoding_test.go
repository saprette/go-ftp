package common_test

import (
	"github.com/unloadedroe/common"
	"testing"
)

func TestRunLengthEncoding1(t *testing.T) {
	chunk := common.NewChunk()
	chunk.SetContent(&[]int{'g', 'o', 'o', 'g', 'l', 'e'}, 6)
	compressedChunk := common.RunLengthEncoding(chunk)
	// (g, 103) (o, 111) (l, 108) (e, 101)
	if compressedChunk.Content[0] != 1 && compressedChunk.Content[1] != 103 {
		t.Fatal("Error, wrong values.")
	}
}

func TestRunLengthEncoding2(t *testing.T) {
	chunk := common.NewChunk()
	chunk.SetContent(&[]int{'g', 'o', 'o', 'g', 'l', 'e'}, 6)
	compressedChunk := common.RunLengthEncoding(chunk)
	// (g, 103) (o, 111) (l, 108) (e, 101)
	if compressedChunk.Content[2] != 2 && compressedChunk.Content[3] != 111 {
		t.Fatalf("Error, wrong values. %d %d", compressedChunk.Content[2], compressedChunk.Content[3])
	}
}

func TestRunLengthEncoding3(t *testing.T) {
	chunk := common.NewChunk()
	chunk.SetContent(&[]int{'g', 'o', 'o', 'g', 'l', 'e'}, 6)
	compressedChunk := common.RunLengthEncoding(chunk)
	// (g, 103) (o, 111) (l, 108) (e, 101)
	if compressedChunk.Content[4] != 1 && compressedChunk.Content[5] != 103 {
		t.Fatal("Error, wrong values.")
	}
}

func TestRunLengthEncoding4(t *testing.T) {
	chunk := common.NewChunk()
	chunk.SetContent(&[]int{'g', 'o', 'o', 'g', 'l', 'e'}, 6)
	compressedChunk := common.RunLengthEncoding(chunk)
	// (g, 103) (o, 111) (l, 108) (e, 101)
	if compressedChunk.Content[6] != 1 && compressedChunk.Content[7] != 108 {
		t.Fatal("Error, wrong values.")
	}
}

func TestRunLengthEncoding5(t *testing.T) {
	chunk := common.NewChunk()
	chunk.SetContent(&[]int{'g', 'o', 'o', 'g', 'l', 'e'}, 6)
	compressedChunk := common.RunLengthEncoding(chunk)
	// (g, 103) (o, 111) (l, 108) (e, 101)
	if compressedChunk.Content[8] != 1 && compressedChunk.Content[9] != 101 {
		t.Fatal("Error, wrong values.")
	}
}

func TestGetPairs(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'g', 'o', 'l', 'a', 'n', 'g'}, 6)
	listOfPairs := common.GetPairs(compressedChunk)
	if listOfPairs.Len() != 6 {
		t.Fatalf("size does not match!. List of pairs => %d", listOfPairs.Len())
	}
}

func TestGetPairs2(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'g', 'o', 'o', 'g', 'l', 'e'}, 6)
	listOfPairs := common.GetPairs(compressedChunk)
	if listOfPairs.Len() != 5 {
		t.Fatalf("size does not match!. List of pairs => %d", listOfPairs.Len())
	}
}

func TestGetPairs3(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'g', 'o', 'o', 'o', 'l', 'e'}, 6)
	listOfPairs := common.GetPairs(compressedChunk)
	if listOfPairs.Len() != 4 {
		t.Fatalf("size does not match!. List of pairs => %d", listOfPairs.Len())
	}
}

func TestGetPairs4(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'g', 'o', 'o', 'o', 'o', 'e'}, 6)
	listOfPairs := common.GetPairs(compressedChunk)
	if listOfPairs.Len() != 3 {
		t.Fatalf("size does not match!. List of pairs => %d", listOfPairs.Len())
	}
}

func TestGetPairs5(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'g', 'o', 'o', 'o', 'o', 'o'}, 6)
	listOfPairs := common.GetPairs(compressedChunk)
	if listOfPairs.Len() != 2 {
		t.Fatalf("size does not match!. List of pairs => %d", listOfPairs.Len())
	}
}

func TestGetPairs6(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'o', 'o', 'o', 'o', 'o', 'o'}, 6)
	listOfPairs := common.GetPairs(compressedChunk)
	if listOfPairs.Len() != 1 {
		t.Fatalf("size does not match!. List of pairs => %d", listOfPairs.Len())
	}
}

func TestFindBoundsForSequenceOneOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'g', 'o', 'l', 'a', 'n', 'g'}, 6)
	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 1 {
		t.Fatal("offset does not match!.")
	}
}

func TestFindBoundsForSequenceTwoOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'g', 'g', 'l', 'a', 'n', 'g'}, 6)
	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 2 {
		t.Fatal("offset does not match!.")
	}
}

func TestFindBoundsForSequenceThreeOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'g', 'g', 'g', 'a', 'n', 'g'}, 6)
	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 3 {
		t.Fatal("offset does not match!.")
	}
}

func TestFindBoundsForSequenceFourOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'g', 'g', 'g', 'g', 'n', 'g'}, 6)
	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 4 {
		t.Fatal("offset does not match!.")
	}
}

func TestFindBoundsForSequenceFiveOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'g', 'g', 'g', 'g', 'g', 'x'}, 6)
	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 5 {
		t.Fatal("offset does not match!.")
	}
}

func TestFindBoundsForSequenceSixOffset(t *testing.T) {
	compressedChunk := common.NewChunk()
	compressedChunk.SetContent(&[]int{'g', 'g', 'g', 'g', 'g', 'g'}, 6)
	lastOffset := common.FindBoundsForSubChunk(compressedChunk, 0)
	if lastOffset != 6 {
		t.Fatal("offset does not match!.")
	}
}
