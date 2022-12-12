package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	DirectionUp    Direction = "U"
	DirectionRight Direction = "R"
	DirectionDown  Direction = "D"
	DirectionLeft  Direction = "L"
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type Rope struct {
	Head       Point
	Tail       Point
	TailVisits map[string]struct{}
}

func NewRope() *Rope {
	return &Rope{
		TailVisits: map[string]struct{}{},
	}
}

func (r *Rope) Move(direction Direction, distance int) {
	for i := 0; i < distance; i++ {
		switch direction {
		case DirectionLeft:
			r.Head.X -= 1
		case DirectionRight:
			r.Head.X += 1
		case DirectionUp:
			r.Head.Y += 1
		case DirectionDown:
			r.Head.Y -= 1
		default:
			log.Panicf("Unknown direction %q", string(direction))
		}
		r.moveTailIfRequired()
	}
}

func (r *Rope) moveTailIfRequired() {
	gapX := r.Head.X - r.Tail.X
	gapY := r.Head.Y - r.Tail.Y
	if gapX == -2 {
		r.Tail.X = r.Head.X + 1
		r.Tail.Y = r.Head.Y
	} else if gapX == 2 {
		r.Tail.X = r.Head.X - 1
		r.Tail.Y = r.Head.Y
	} else if gapY == -2 {
		r.Tail.X = r.Head.X
		r.Tail.Y = r.Head.Y + 1
	} else if gapY == 2 {
		r.Tail.X = r.Head.X
		r.Tail.Y = r.Head.Y - 1
	}
	r.TailVisits[r.Tail.String()] = struct{}{}
}

func main() {
	bytes, err := os.ReadFile("./input.txt")
	handleError(err)
	rope := NewRope()
	for _, line := range strings.Split(string(bytes), "\n") {
		parts := strings.SplitN(line, " ", 2)
		direction := Direction(parts[0])
		distance, err := strconv.Atoi(parts[1])
		handleError(err)
		rope.Move(direction, distance)
	}
	fmt.Printf("The tail visited %d different positions at least once\n", len(rope.TailVisits))
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("an error has occurred: %v", err)
	}
}
