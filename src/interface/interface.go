package main

import (
	"math"
)

type Abser interface {
	abs() float64
} 


type Vertex struct {
	X, Y float64
	
}

func (v *Vertex) abs() float64 {
      return math.Sqrt(v.X*v.X + v.Y*v.Y)
}