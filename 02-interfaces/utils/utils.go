package utils

import "fmt"

type ShapeWithArea interface {
	Area() float32
}

func PrintArea(sa ShapeWithArea) {
	fmt.Println("Area = ", sa.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func PrintPerimeter(sp ShapeWithPerimeter) {
	fmt.Println("Perimeter = ", sp.Perimeter())
}

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintShape(x Shape) {
	PrintArea(x)
	PrintPerimeter(x)
}
