package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.width + rect.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.radius
}

/*
第一题
定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter()方法。
*/
func main() {
	var rect Rectangle = Rectangle{10, 5}
	fmt.Printf("rectangle area: %f\n", rect.Area())
	fmt.Printf("rectangle perimeter: %f\n", rect.Perimeter())
	var circle Circle = Circle{10}
	fmt.Printf("circle area: %f\n", circle.Area())
	fmt.Printf("circle perimeter: %f\n", circle.Perimeter())
}
