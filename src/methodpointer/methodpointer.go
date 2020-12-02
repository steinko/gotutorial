package main

type Vertex struct { X, Y float64}

func (v *Vertex) scale( s float64) {
        v.X = v.X*s
        v.Y = v.Y*s
}
