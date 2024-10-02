package structsmethodsinterfaces

import "math"

// interface declaration which requires method return float
type Shape interface {
	Area() float64
}

// structs declared
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) (ans float64) {
	ans = 2 * (rectangle.Width + rectangle.Height)
	return
}

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }

// overloading function calling arguments is not legal in Golang, instead you
// can declare functions of the same name in different packages.
// but that's overkill here as well

// methods:
// converted func Area() to method Area() which can be called on
// different variable types.

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
