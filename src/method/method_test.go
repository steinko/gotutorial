package main

import (  
	      "math"
	      "testing"
          "github.com/stretchr/testify/assert"
        )

func TestMethod (t *testing.T) {
	assert := assert.New(t)
	v := Vertex{3,4}
	assert.Equal(v.abs(), math.Sqrt(3*3 + 4*4))
}