package main

import (  "testing"
          "github.com/stretchr/testify/assert"
        )


func TestStrutureliterates( t *testing.T) {
	 assert := assert.New(t)
	 t.Run("should deliver a array", func(t *testing.T) {
	    a:= array()
	    assert.Equal(a[0],"Hello")
	    assert.Equal(a[1],"World")
	 })
	 	t.Run("should deliver prime", func(t *testing.T) {
	 			prime := primes()
	 			assert.Equal(prime[0],1)
	            assert.Equal(prime[1],3)
	            assert.Equal(prime[2],5)
	            assert.Equal(prime[3],7)
	            assert.Equal(prime[4],11)
	            assert.Equal(prime[5],13)
	 		})
}