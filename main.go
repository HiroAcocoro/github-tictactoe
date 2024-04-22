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

func main() {
	// 27 is the start of the Board
	var renderBoard string = EmptyBoard
	var isCircle bool = true

	for _, move := range Moves {
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

	updatedReadme := []byte(TemplateMsg + renderBoard)
	err := os.WriteFile("README.md", updatedReadme, 0644)
	if err != nil {
		fmt.Println("Failed to write : %v", err)
	}
}
