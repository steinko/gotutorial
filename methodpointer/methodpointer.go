package main

import (
     "fmt"
)

type Vertex struct { X, Y float64}

func (v *Vertex) scale( s float64) {
        v.X = v.X*s
        v.Y = v.Y*s
}

func main () { 
     v := Vertex{ 1,2}
     v.scale(2)
     fmt.Println(v.X ,v.Y)
}



