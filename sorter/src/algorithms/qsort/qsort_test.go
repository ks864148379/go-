package qsort

import "testing"

func TestQuickSort1(t *testing.T) {
	values := []int{5, 3, 1}
	QuickSort(values)
	if values[0] != 1 || values[1] != 3 || values[2] != 5 {
		t.Error("QuickFailed")
	}
}