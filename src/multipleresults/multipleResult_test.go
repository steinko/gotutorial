package main

import (
    "testing"
)

func TestSwap(t *testing.T) {
	got1,got2  := Swap("Hello","World")
	if got1 != "World" {
		t.Error("should return: World", got1)
	}
	
	if got2 != "Hello" {
		t.Error("should return: Hello", got2)
	}
}

