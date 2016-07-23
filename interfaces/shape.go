package interfaces

type Shape interface {
	Area() int
}

type Square struct {
	side int
}

type Rectangle struct {
	length int
	breadth int
}

type Hybrid struct {
	shape1 Shape
	shape2 Shape
}

func (square Square) Area() int {
	return square.side * square.side
}

func (rectangle Rectangle) Area() int {
	return rectangle.length*rectangle.breadth
}

func (hybrid Hybrid) Area() int {
	return hybrid.shape1.Area() + hybrid.shape2.Area()
}