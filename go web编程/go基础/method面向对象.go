package main

import (
	"fmt"
	"math"
)

// 矩形
type Rectangle struct {
	height, width float64
}

// 圆形
type Circle struct {
	radius float64
}

func main() {
	countArea()
}

func (r Rectangle) area() float64 {
	// Rectangle为接收器的method
	rectangleArea := r.height * r.width
	return rectangleArea
}

func (r Circle) area() float64 {
	// Circle为接收器的method
	radius := r.radius
	circleArea := radius * radius * math.Pi
	return circleArea
}

func countArea() {
	var r1 Rectangle = Rectangle{10, 20}
	fmt.Println(r1.area())

	r2 := Rectangle{15, 25}
	fmt.Println(r2.area())

	var r3 Rectangle
	r3.width = 24
	r3.height = 12
	fmt.Println(r3.area())

	c1 := Circle{10}
	fmt.Println(c1.area())
}
