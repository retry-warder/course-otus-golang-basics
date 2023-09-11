package types

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius int
	area   float64
}

func NewCircle(radius int) *Circle {
	circle := Circle{radius, float64(math.Pi) * math.Pow(float64(radius), 2)}
	return &circle
}

func (c *Circle) Area() float64 {
	return c.area
}

func (c *Circle) String() string {
	tmplts := "Круг: радиус %d Площадь: %f"
	return fmt.Sprintf(tmplts, c.radius, c.area)
}

type Triangle struct {
	base   int
	height int
	area   float64
}

func NewTriangle(base int, height int) *Triangle {
	triangle := Triangle{base, height, (float64(base) * float64(height)) / 2}
	return &triangle
}

func (t *Triangle) Area() float64 {
	return t.area
}

func (t *Triangle) String() string {
	tmplts := "Треугольник: основание %d, высота %d Площадь: %f"
	return fmt.Sprintf(tmplts, t.base, t.height, t.area)
}

type Rectangle struct {
	length int
	width  int
	area   float64
}

func NewRectangle(length int, width int) *Rectangle {
	rectangle := Rectangle{length, width, float64(length) * float64(width)}
	return &rectangle
}

func (r Rectangle) Area() float64 {
	return r.area
}

func (r Rectangle) String() string {
	tmplts := "Прямоугольник: ширина %d, высота %d Площадь: %f"
	return fmt.Sprintf(tmplts, r.length, r.width, r.area)
}

func CalculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if ok {
		fmt.Println(shape)
		return shape.Area(), nil
	}
	fmt.Println("Ошибка: переданный объект не является фигурой.")
	return 0.0, errors.New("ошибка: переданный объект не является фигурой")
}
