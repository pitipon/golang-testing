package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Init
	elements := []int{9, 7, 6, 4, 2, 5, 8, 4, 0, 8, 7}
	fmt.Println(elements)
	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 9 {
		t.Error("First element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("Last element should be 0")
	}
	fmt.Println(elements)

}
