package main

import (
	"hw05_shapes/types"
)

func main() {
	c := types.NewCircle(8)
	r := types.NewRectangle(2, 5)
	t := types.NewTriangle(5, 10)
	a := "shape"
	types.CalculateArea(c)
	types.CalculateArea(r)
	types.CalculateArea(t)
	types.CalculateArea(a)
}
