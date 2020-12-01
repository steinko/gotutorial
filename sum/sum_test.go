package array

import "testing"

func TestSum(t *testing.T) {
	t.Run("should deliver 15", func(t *testing.T) {
	  array := [5]int{1,2,3,4,5}
	  result := Sum(array)
	  expected := 15
	  if (result != expected) {
		t.Errorf("got %d want %d given, %v", result, expected, array)
	  }
	})
	
	t.Run("should deliver 16", func(t *testing.T) {
	  array := [5]int{2,2,3,4,5}
	  result := Sum(array)
	  expected := 16
	  if (result != expected) {
		t.Errorf("got %d want %d given, %v", result, expected, array)
	  }
	})
}