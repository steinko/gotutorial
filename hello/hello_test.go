package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("should say hellow world", func(t *testing.T) {
	  got := Hello("World")
      want := "Hello World"

      if got != want {
        t.Errorf("got %q want %q", got, want)
      }
	})
	
	t.Run("should say hellow stein", func(t *testing.T) {
	  got := Hello("Stein")
      want := "Hello Stein"

      if got != want {
        t.Errorf("got %q want %q", got, want)
      }
	})
	
	t.Run("should say hellow world", func(t *testing.T) {
	  got := Hello("")
      want := "Hello World"

      if got != want {
        t.Errorf("got %q want %q", got, want)
      }
	})
}