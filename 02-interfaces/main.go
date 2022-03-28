package main

import (
	"interfaces-demo/circle"
	"interfaces-demo/rectangle"
	"interfaces-demo/utils"
)

func main() {
	c := circle.Circle{Radius: 12}
	r := rectangle.Rectangle{Height: 10, Width: 12}
	utils.PrintShape(c)
	utils.PrintShape(r)
}
