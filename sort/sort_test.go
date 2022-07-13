package sort

import (
	"testing"
)

func Test_quickSort(t *testing.T) {
	a := []int{1, 3, 2}
	quickSort(a)
	t.Log(a)
}
