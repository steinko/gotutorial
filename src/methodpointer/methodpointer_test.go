package main

import (  
	      "testing"
          "github.com/stretchr/testify/assert"
        )

func TestMethodPointer (t *testing.T) { 
	 assert := assert.New(t)
     v := Vertex{ 1,2}
     v.scale(2)
     assert.Equal(v.X,2.0)
     assert.Equal(v.Y,4.0)
}
