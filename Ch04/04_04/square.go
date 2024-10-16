package main

import (
	"fmt"
	"log"
)

// Square is a square
type Square struct {
	x, y, length int
}

// NewSquare returns a new square
func NewSquare(x int, y int, length int) (*Square, error) {
	if length < 1 {
		return nil, fmt.Errorf("length must be > 0")
	}

	square := Square{
		x: x,
		y: y,
		length: length,
	}

	return &square, nil
}

// Move moves the square
func (s *Square) Move(dx int, dy int) {
	s.x += dx
	s.y += dy
}

// Area returns the square are
func (s *Square) Area() int {
	return s.length * s.length
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
