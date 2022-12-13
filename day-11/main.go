package main

import (
	"log"
	"math"
	"sort"
)

type MonkeyID int

type WorryLevel float64

func (l WorryLevel) Float64() float64 {
	return float64(l)
}

type InspectFunc func(level WorryLevel) WorryLevel

const superMod = float64(11 * 17 * 5 * 13 * 19 * 2 * 3 * 7)

type Monkey struct {
	ID            MonkeyID
	Items         []WorryLevel
	receivedItems []WorryLevel
	InspectFunc   InspectFunc
	WorryModulo   float64
	TrueTo        MonkeyID
	FalseTo       MonkeyID
	InspectCount  int
}

func (m *Monkey) Inspect(item WorryLevel, divider float64) (MonkeyID, WorryLevel) {
	m.InspectCount++
	newWorryLevel := math.Mod(math.Floor(m.InspectFunc(item).Float64()/divider), superMod)
	mod := math.Mod(newWorryLevel, m.WorryModulo)
	if mod == 0 {
		return m.TrueTo, WorryLevel(newWorryLevel)
	}
	return m.FalseTo, WorryLevel(newWorryLevel)
}

func (m *Monkey) Receive(level WorryLevel) {
	m.Items = append(m.Items, level)
}

type Squad struct {
	monkeys []*Monkey
}

func (s *Squad) ProcessRounds(count int, divider float64) {
	for i := 0; i < count; i++ {
		s.ProcessRound(divider)
	}
}

func (s *Squad) ProcessRound(divider float64) {
	for _, monkey := range s.monkeys {
		for _, item := range monkey.Items {
			receiver, worryLevel := monkey.Inspect(item, divider)
			s.monkeys[receiver].Receive(worryLevel)
		}
		monkey.Items = nil
	}
}

func (s *Squad) ReportMonkeyBusiness() {
	var inspectCounts []int
	for _, monkey := range s.monkeys {
		println("Monkey", monkey.ID, "inspected items", monkey.InspectCount, "times")
		inspectCounts = append(inspectCounts, monkey.InspectCount)
	}
	sort.Slice(inspectCounts, func(i, j int) bool { return inspectCounts[i] > inspectCounts[j] })
	println("Monkey business is", inspectCounts[0]*inspectCounts[1])
}

func main() {
	println("Part 1")
	squad := newSquad()
	squad.ProcessRounds(20, 3)
	squad.ReportMonkeyBusiness()

	println("Part 2")
	squad = newSquad()
	squad.ProcessRounds(10000, 1)
	squad.ReportMonkeyBusiness()
}

func newSquad() *Squad {
	return &Squad{
		monkeys: []*Monkey{
			{
				ID:          MonkeyID(0),
				Items:       []WorryLevel{98, 97, 98, 55, 56, 72},
				InspectFunc: func(level WorryLevel) WorryLevel { return level * 13 },
				WorryModulo: 11,
				TrueTo:      MonkeyID(4),
				FalseTo:     MonkeyID(7),
			},
			{
				ID:          MonkeyID(1),
				Items:       []WorryLevel{73, 99, 55, 54, 88, 50, 55},
				InspectFunc: func(level WorryLevel) WorryLevel { return level + 4 },
				WorryModulo: 17,
				TrueTo:      MonkeyID(2),
				FalseTo:     MonkeyID(6),
			},
			{
				ID:          MonkeyID(2),
				Items:       []WorryLevel{67, 98},
				InspectFunc: func(level WorryLevel) WorryLevel { return level * 11 },
				WorryModulo: 5,
				TrueTo:      MonkeyID(6),
				FalseTo:     MonkeyID(5),
			},
			{
				ID:          MonkeyID(3),
				Items:       []WorryLevel{82, 91, 92, 53, 99},
				InspectFunc: func(level WorryLevel) WorryLevel { return level + 8 },
				WorryModulo: 13,
				TrueTo:      MonkeyID(1),
				FalseTo:     MonkeyID(2),
			},
			{
				ID:          MonkeyID(4),
				Items:       []WorryLevel{52, 62, 94, 96, 52, 87, 53, 60},
				InspectFunc: func(level WorryLevel) WorryLevel { return level * level },
				WorryModulo: 19,
				TrueTo:      MonkeyID(3),
				FalseTo:     MonkeyID(1),
			},
			{
				ID:          MonkeyID(5),
				Items:       []WorryLevel{94, 80, 84, 79},
				InspectFunc: func(level WorryLevel) WorryLevel { return level + 5 },
				WorryModulo: 2,
				TrueTo:      MonkeyID(7),
				FalseTo:     MonkeyID(0),
			},
			{
				ID:          MonkeyID(6),
				Items:       []WorryLevel{89},
				InspectFunc: func(level WorryLevel) WorryLevel { return level + 1 },
				WorryModulo: 3,
				TrueTo:      MonkeyID(0),
				FalseTo:     MonkeyID(5),
			},
			{
				ID:          MonkeyID(7),
				Items:       []WorryLevel{70, 59, 63},
				InspectFunc: func(level WorryLevel) WorryLevel { return level + 3 },
				WorryModulo: 7,
				TrueTo:      MonkeyID(4),
				FalseTo:     MonkeyID(3),
			},
		},
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("an error has occurred: %v", err)
	}
}
