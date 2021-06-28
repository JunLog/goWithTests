package structs

import "math"

// 接口
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// 方法
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// 方法
func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}
