package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Measurable interface {
	Perimeter() float64
}

type Geometry interface {
	Shape
	Measurable
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64{
	return 2*(r.width + r.height)
}

func describeShape(g Geometry){
	fmt.Println(g.Area())
	fmt.Println(g.Perimeter())
}

func main(){
	rect := Rectangle{width: 5, height: 4}

	describeShape(rect)
}