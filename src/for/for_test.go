package main

import (  "testing"
          "github.com/stretchr/testify/assert"
        )

func TestStrutureliterates( t *testing.T) {
	 assert := assert.New(t)
	 
	 
	 	t.Run("should deliver from 1 to 10", func(t *testing.T) {
	 		
	 			cool := cool()
	 			assert.Equal(cool[0],1)
	            assert.Equal(cool[1],2)
	            assert.Equal(cool[2],3)
	            assert.Equal(cool[3],4)
	            assert.Equal(cool[4],5)
	            assert.Equal(cool[5],6)
	 		})
	 	
}