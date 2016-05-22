package common_test

import (
	"github.com/unloadedroe/common"
	"testing"
)

func TestCreatePairFirst(t *testing.T) {
	pair := common.NewPair(1, 2)
	first := pair.First().(int)
	if first != 1 {
		t.Fatal("Incorrect value.")
	}
}

func TestCreatePairSecond(t *testing.T) {
	pair := common.NewPair(1, 2)
	first := pair.Second().(int)
	if first != 2 {
		t.Fatal("Incorrect value.")
	}
}
