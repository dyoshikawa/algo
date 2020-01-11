package p269

import (
	"github.com/dyoshikawa/algo/util"
	"testing"
)

func TestInvoke(t *testing.T) {
	as := [][]int{
		{1, 2, 2, 4},
		{2, 1, 4},
		{3, 0},
		{4, 1, 3},
	}
	expected := [][]int{
		{0, 1, 0, 1},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 1, 0},
	}
	if res := Invoke(4, as); !util.EqualDoubleArrInt(expected, res) {
		t.Fatal(res)
	}
}
