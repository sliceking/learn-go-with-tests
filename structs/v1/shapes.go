package main

import "math"

type Shape interface {
	Area() float64
}

// Perimeter returns the perimeter of a rectangle.
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

type Rectangle struct {
	Height float64
	Width  float64
}

// Area returns the perimeter of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// Area returns the perimeter of a rectangle.
func (c Circle) Area() float64 {
	return  math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base float64
	Height float64
}

// Area returns the perimeter of a rectangle.
func (t Triangle) Area() float64 {
	return  (t.Height * t.Base) / 2
}
