package types

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius int
}

func NewCircle(radius int) *Circle {
	circle := Circle{radius}
	return &circle
}

func (c *Circle) Area() float64 {
	area := float64(math.Pi) * math.Pow(float64(c.radius), 2)
	fmt.Printf("Круг: радиус %d Площадь: %f", c.radius, area)
	fmt.Println()
	return area
}

type Triangle struct {
	base   int
	height int
}

func NewTriangle(base int, height int) *Triangle {
	triangle := Triangle{base, height}
	return &triangle
}

func (t *Triangle) Area() float64 {
	area := (float64(t.base) * float64(t.height)) / 2
	fmt.Printf("Треугольник: основание %d, высота %d Площадь: %f", t.base, t.height, area)
	fmt.Println()
	return area
}

type Rectangle struct {
	length int
	width  int
}

func NewRectangle(length int, width int) *Rectangle {
	rectangle := Rectangle{length, width}
	return &rectangle
}

func (r Rectangle) Area() float64 {
	area := float64(r.length) * float64(r.width)
	fmt.Printf("Прямоугольник: ширина %d, высота %d Площадь: %f", r.length, r.width, area)
	fmt.Println()
	return area
}

func CalculateArea(s any) float64 {
	var area float64
	shape, ok := s.(Shape)
	if ok {
		area = shape.Area()
	} else {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
		area = 0
	}
	return area
}
