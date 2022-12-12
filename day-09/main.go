package main

import (
	"fmt"
	"log"
	"math"
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
	Tails      []Point
	TailVisits map[string]struct{}
}

func NewRope(tailCount int) *Rope {
	return &Rope{
		Tails:      make([]Point, tailCount),
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
		r.moveTailsIfRequired()
	}
}

func (r *Rope) moveTailsIfRequired() {
	for i := range r.Tails {
		r.moveTailIfRequired(i)
	}
}

func (r *Rope) moveTailIfRequired(position int) {
	previous := r.Head
	if position > 0 {
		previous = r.Tails[position-1]
	}
	gapX := previous.X - r.Tails[position].X
	gapY := previous.Y - r.Tails[position].Y
	if math.Abs(float64(gapX)) > 1 && math.Abs(float64(gapY)) > 1 {
		r.Tails[position].X += gapX / 2
		r.Tails[position].Y += gapY / 2
	} else if gapX < -1 {
		r.Tails[position].X = previous.X + 1
		r.Tails[position].Y = previous.Y
	} else if gapX > 1 {
		r.Tails[position].X = previous.X - 1
		r.Tails[position].Y = previous.Y
	} else if gapY < -1 {
		r.Tails[position].X = previous.X
		r.Tails[position].Y = previous.Y + 1
	} else if gapY > 1 {
		r.Tails[position].X = previous.X
		r.Tails[position].Y = previous.Y - 1
	}
	if position == len(r.Tails)-1 {
		r.TailVisits[r.Tails[position].String()] = struct{}{}
	}
}

func main() {
	bytes, err := os.ReadFile("./input.txt")
	handleError(err)
	ropeWithTwoKnots := NewRope(1)
	ropeWithTenKnots := NewRope(9)
	for _, line := range strings.Split(string(bytes), "\n") {
		parts := strings.SplitN(line, " ", 2)
		direction := Direction(parts[0])
		distance, err := strconv.Atoi(parts[1])
		handleError(err)
		ropeWithTwoKnots.Move(direction, distance)
		ropeWithTenKnots.Move(direction, distance)
	}
	fmt.Printf("The tail of the rope with 2 knots visited %d different positions at least once\n", len(ropeWithTwoKnots.TailVisits))
	fmt.Printf("The tail of the rope with 10 knots visited %d different positions at least once\n", len(ropeWithTenKnots.TailVisits))
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("an error has occurred: %v", err)
	}
}
