package main

import (  "testing"
          "github.com/stretchr/testify/assert"
        )

func TestFunction( t *testing.T) {
	 assert := assert.New(t)
	 assert.Equal(add(2,3),5)
}