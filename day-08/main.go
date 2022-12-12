package main

import (
	"log"
	"math"
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
	scenicScores := calculateScenicScores(forest)
	max := float64(0)
	for _, score := range scenicScores {
		max = math.Max(max, score)
	}
	println("Highest scenic score:", int(max))
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

func calculateScenicScores(forest [][]rune) []float64 {
	var scores []float64
	for x := 0; x < len(forest); x++ {
		for y := 0; y < len(forest[x]); y++ {
			scores = append(scores, calculateScenicScore(forest, x, y))
		}
	}
	return scores
}

func calculateScenicScore(forest [][]rune, x int, y int) float64 {
	currentHeight := forest[x][y]
	left := 0
	for i := x - 1; i >= 0; i-- {
		left++
		if forest[i][y] >= currentHeight {
			break
		}
	}
	right := 0
	for i := x + 1; i < len(forest); i++ {
		right++
		if forest[i][y] >= currentHeight {
			break
		}
	}
	top := 0
	for i := y - 1; i >= 0; i-- {
		top++
		if forest[x][i] >= currentHeight {
			break
		}
	}
	bottom := 0
	for i := y + 1; i < len(forest[x]); i++ {
		bottom++
		if forest[x][i] >= currentHeight {
			break
		}
	}
	return float64(top * right * bottom * left)
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("an error has occurred: %v", err)
	}
}
