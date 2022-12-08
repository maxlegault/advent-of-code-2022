package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("./input.txt")
	handleError(err)
	var forest [][]rune
	for _, line := range strings.Split(string(bytes), "\n") {
		var row []rune
		for _, tree := range line {
			row = append(row, tree)
		}
		forest = append(forest, row)
	}
	visibleCount := 0 //len(forest)*2 + (len(forest[0])-2)*2
	for x := 0; x < len(forest); x++ {
		for y := 0; y < len(forest[x]); y++ {
			if x == 0 || x == len(forest)-1 || y == 0 || y == len(forest[x])-1 || isTreeVisible(forest, x, y) {
				visibleCount++
			}
		}
	}
	println("Number of visible trees:", visibleCount)
}

func isTreeVisible(forest [][]rune, treeX int, treeY int) bool {
	treeHeight := forest[treeX][treeY]
	visible := true
	for x := 0; x < treeX; x++ {
		if forest[x][treeY] >= treeHeight {
			visible = false
		}
	}
	if visible {
		return true
	}
	visible = true
	for x := treeX + 1; x < len(forest); x++ {
		if forest[x][treeY] >= treeHeight {
			visible = false
		}
	}
	if visible {
		return true
	}
	visible = true
	for y := 0; y < treeY; y++ {
		if forest[treeX][y] >= treeHeight {
			visible = false
		}
	}
	if visible {
		return true
	}
	visible = true
	for y := treeY + 1; y < len(forest[treeX]); y++ {
		if forest[treeX][y] >= treeHeight {
			visible = false
		}
	}
	return visible
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("an error has occurred: %v", err)
	}
}
