package main

import (
	"fmt"

	"github.com/retry-warder/course-otus-golang-basics/hw05_shapes/types"
)

func CalcArea(s any) {
	var area float64
	area, err := types.CalculateArea(s)
	if err == nil {
		if area == 0 {
			fmt.Println("Ошибка: неверно рассчитана площадь фигуры.")
		} else {
			fmt.Println(s)
		}
	} else {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	}
}

func main() {
	c := types.NewCircle(8)
	r := types.NewRectangle(2, 5)
	t := types.NewTriangle(5, 10)
	a := "shape"
	CalcArea(c)
	CalcArea(r)
	CalcArea(t)
	CalcArea(a)
}
