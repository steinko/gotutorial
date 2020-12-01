package iteration

import "testing"

func TestIteration(t *testing.T) {
     t.Run("should repeat aaaa letter", func(t *testing.T) {
   		  got := Repeate("a",4)
   		  result := "aaaa"
   		  if got != result {
   			t.Errorf("not repated expected %q but got %q", result,got)
   		  }
      })
     
       t.Run("should repeat bbbbb letter", func(t *testing.T) {
   		  got := Repeate("b",5)
   		  result := "bbbbb"
   		  if got != result {
   			t.Errorf("not repated expected %q but got %q", result,got)
   		  }
      })
}