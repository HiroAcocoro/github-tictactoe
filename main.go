package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func letterToNum(value string) int {
	letterMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	return letterMap[strings.ToLower(value)]
}

func hasNoDuplicate(moves []string) bool {
	occured := make(map[string]bool)
	for _, move := range moves {
		if occured[move] {
			return false
		}
		occured[move] = true
	}
	return true
}

func main() {
	var moves = [...]string{"c1", "a1", "b3", "b2", "c1"}
	// constants
	var emptyBoard string = "||a|b|c|\n|---|---|---|---|\n|1|⬛|⬛|⬛|\n|2|⬛|⬛|⬛|\n|3|⬛|⬛|⬛|"
	var templateMsg string = "# GITHUB IS A GAME ENGINE\n\n### PR a new move inside main.go > moves\n\n### ToDo\n- Github actions to set and render your moves\n- Move validations\n- Win game scenario\n\n\n# CURRENT GAME:\n\n\n"
	var isMovesValid bool = false

	isMovesValid = hasNoDuplicate(moves[:])

	fmt.Println(isMovesValid)
	// 27 is the start of the Board
	var renderBoard string = emptyBoard
	var isCircle bool = true

	for _, move := range moves {
		var boardArr = strings.Split(renderBoard, "")
		var movePiece string
		if isCircle {
			movePiece = "⭕"
		} else {
			movePiece = "❌"
		}

		// abc notations
		var xAxis = letterToNum(move[0:1])*2 + 1
		// 123 notations
		yAxis, err := strconv.Atoi(move[1:])
		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			return
		}
		yAxis = yAxis * 10

		boardArr[xAxis+yAxis+17] = movePiece
		isCircle = !isCircle
		renderBoard = strings.Join(boardArr, "")
	}

	updatedReadme := []byte(templateMsg + renderBoard)
	err := os.WriteFile("README.md", updatedReadme, 0644)
	if err != nil {
		fmt.Println("Failed to write : %v", err)
	}
}
