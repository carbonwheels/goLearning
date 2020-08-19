package objects

// Version 1: Calculates preimeter
func Perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// Version 2: Calulates the area
func Area(width float64, height float64) float64 {
	return width * height
}

// Version 3: Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Version 3: Perimeter returns the perimeter of the rectangle.
func Perimeter3(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Version 3: Area returns the area of the rectangle.
func Area3(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
