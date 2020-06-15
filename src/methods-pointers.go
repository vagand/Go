package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (s *Vertex) Scale(f float64) {
	s.X = s.X * f
	s.Y = s.Y * f
}

func (s Vertex) Scale2(f float64) {
	s.X = s.X * f
	s.Y = s.Y * f
}

func (s Vertex) Print() {
	fmt.Println(s.X, s.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale2(10)
	v.Print()
	v.Scale(10)
	v.Print()
	fmt.Println(v.Abs())
}
