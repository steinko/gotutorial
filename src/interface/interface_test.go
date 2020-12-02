package main

import 
      (
      	"math"
      	"testing"
        "github.com/stretchr/testify/assert"
      )

func TestInterface (t *testing.T) {
	assert := assert.New(t)
	var a Abser
	v := Vertex{2,3}
	a = &v
	assert.Equal(a.abs(),math.Sqrt(2*2 + 3*3))
	
}