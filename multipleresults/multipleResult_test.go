package main

import (
    "testing"
)

func TestSwap(t *testing.T) {
	got , get := Swap("Hello","World2")
	if got != "World Hello" {
		t.Errorf("should return: World Hello", got)
	}

}

