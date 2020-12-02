package main

import (  "testing"
          "github.com/stretchr/testify/assert"
        )

func TestStructponters(t *testing.T) {
  assert := assert.New(t)	
  v := Vertex{1, 2}
  p := &v
  p.X = 1e9
  assert.Equal(p.X, int(1e9))
  assert.Equal(p.Y,2)
}