package main

import (
	"fmt"
	"math"
)

type Vertex struct { X, Y float64}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func method () {
	v := Vertex{1,2}
	fmt.Println(v.abs())
}