package shapes

import "math"

func (c Rectangle) Perimeter() float64 {
	return 2 * (c.Width + c.Height)
}

func (c Rectangle) Area() float64 {
	return c.Width * c.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
