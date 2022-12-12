package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Register rune

const MainRegister Register = 'X'

type Instruction string

const AddXInstruction Instruction = "addx"
const NoopInstruction Instruction = "noop"

func NewDevice() *Device {
	return &Device{
		registers: map[Register]int{
			MainRegister: 1,
		},
		currentCycle: 1,
	}
}

type SignalStrength struct {
	cycle int
	value int
}

func (s SignalStrength) Strength() int {
	return s.cycle * s.value
}

type Device struct {
	registers       map[Register]int
	currentCycle    int
	signalStrengths []SignalStrength
}

func (d *Device) ProcessInstruction(instruction Instruction, value int) {
	d.checkSignalStrength()
	if instruction == AddXInstruction {
		d.currentCycle++
		d.checkSignalStrength()
		d.registers[MainRegister] += value
	}
	d.currentCycle++
}

func (d *Device) checkSignalStrength() {
	if math.Mod(float64(d.currentCycle+20), 40) == 0 {
		d.signalStrengths = append(d.signalStrengths, SignalStrength{
			cycle: d.currentCycle,
			value: d.registers[MainRegister],
		})
	}
}

func (d *Device) ReportSignalStrengthAfterNCycles(cycleCount int) int {
	sum := 0
	for _, signalStrength := range d.signalStrengths {
		if signalStrength.cycle <= cycleCount {
			sum += signalStrength.Strength()
		}
	}
	return sum
}

func main() {
	bytes, err := os.ReadFile("./input.txt")
	handleError(err)
	device := NewDevice()
	for _, line := range strings.Split(string(bytes), "\n") {
		parts := strings.SplitN(line, " ", 2)
		instruction := Instruction(parts[0])
		var value int
		if len(parts) > 1 {
			value, err = strconv.Atoi(parts[1])
			handleError(err)
		}
		device.ProcessInstruction(instruction, value)
	}
	fmt.Printf("The sum of all signal strengths is %d after 220 cycles\n", device.ReportSignalStrengthAfterNCycles(220))
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("an error has occurred: %v", err)
	}
}
