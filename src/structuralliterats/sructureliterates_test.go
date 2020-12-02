package main

import (  "testing"
          "github.com/stretchr/testify/assert"
        )


func TestStrutureliterates( t *testing.T) {
	 assert := assert.New(t)
	 var ( v1 = Vertex{1, 2}
     v2 = Vertex{X:3}
     v3 = Vertex{})
   //  p = &Vertex{} )
     assert.Equal(v1.X,1)
	 assert.Equal(v1.Y,2)
	 assert.Equal(v2.X,3)
	 assert.Equal(v2.Y,0)
	 assert.Equal(v3.X,0)
	 assert.Equal(v3.Y,0)
	// assert.Equal(p.X,0)
	// assert.Equal(p.Y,0)
}