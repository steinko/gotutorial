package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("should say hello world", func(t *testing.T) {
	  got := Hello("World","English")
      want := "Hello World"

      if got != want {
        t.Errorf("got %q want %q", got, want)
      }
	})
	
	t.Run("should say hello stein", func(t *testing.T) {
	  got := Hello("Stein","Enlish")
      want := "Hello Stein"

      if got != want {
        t.Errorf("got %q want %q", got, want)
      }
	})
	
	t.Run("should say hellow world", func(t *testing.T) {
	  got := Hello("","English")
      want := "Hello World"

      if got != want {
        t.Errorf("got %q want %q", got, want)
      }
	})
	
	t.Run("should say Hola Mundo", func(t *testing.T) {
	  got := Hello("Mundo","Spanish")
      want := "Hola Mundo"

      if got != want {
        t.Errorf("got %q want %q", got, want)
      }
	})
}