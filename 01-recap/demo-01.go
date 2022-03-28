package main

import (
	"fmt"
	"math"
)

//2015
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

//2016
type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

//2022
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

func main() {
	c := Circle{Radius: 12}
	PrintShape(c)
	//2017

	r := Rectangle{Height: 10, Width: 12}
	PrintShape(r)
	//2017

}
