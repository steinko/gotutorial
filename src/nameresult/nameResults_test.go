package main


import (  "testing"
          "github.com/stretchr/testify/assert"
        )

func TestSplit(t *testing.T) {
	assert := assert.New(t)
	got1,got2  := Split(3)
	assert.Equal(got1,1)
	assert.Equal(got2,2)
}

