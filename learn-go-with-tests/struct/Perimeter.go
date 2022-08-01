package _struct

import "math"


type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func Perimeter(rectangle Rectangle) float64 {
	return  2 * (rectangle.Width + rectangle.Height)
}
type Trangle struct {
	Base float64
	Height float64
}

func (c Trangle)Area() float64 {
	return c.Base * c.Height * 0.5
}
type Shape interface {
	Area() float64
}