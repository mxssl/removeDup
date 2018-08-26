package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCase(t *testing.T) {
	input := []int{10, 20, 30, 40, 50, 20, 30, 40, 50}
	expected := []int{10, 20, 30, 40, 50}

	output := removeDuplicate(input)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}
