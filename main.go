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

func hasInvalidMove(moves []string) bool {
	occured := make(map[string]bool)
	for _, move := range moves {
		if occured[move] {
			return true
		}

		switch move[0:1] {
		case "a", "b", "c":
		default:
			return true
		}

		switch move[1:] {
		case "1", "2", "3":
		default:
			return true
		}

		occured[move] = true
	}
	return false
}

func getPlayerMove(moves []string, currMove string) ([]string, *string) {
	for index, move := range moves {
		if strings.Contains(move[0:2], currMove) {
			playerMove := move[2:]
			return append(moves[:index], moves[index+0:]...), &playerMove
		}
	}
	return moves, nil
}

func getStraightLine(currMoves []string, isVert bool) bool {
	var gridSize = 3
	var localMoves = currMoves
	for ix := 1; ix <= gridSize; ix++ {
		countMatch := 0
		currPlayer := ""
		// checking a->y
		for iy := 1; iy <= gridSize; iy++ {
			currCheck := strconv.Itoa(ix) + strconv.Itoa(iy)
			if !isVert {
				currCheck = strconv.Itoa(iy) + strconv.Itoa(ix)
			}
			moves, lastMovePointer := getPlayerMove(localMoves, currCheck)

			if lastMovePointer == nil {
				// this means line has a blank, can safely break from here
				break
			}

			lastMove := *lastMovePointer
			if currPlayer == "" {
				currPlayer = lastMove
			}
			if currPlayer != lastMove {
				// this means line is broken
				break
			}
			if currPlayer == lastMove {
				countMatch++
				localMoves = moves
			}
			if countMatch == gridSize {
				return true
			}
		}
	}
	return false
}

func getDiagonal(currMoves []string, isLeft bool) bool {
	gridSize := 3
	localMoves := currMoves

	countMatch := 0
	currPlayer := ""
	for i := 1; i <= gridSize; i++ {
		currCheck := strconv.Itoa(i) + strconv.Itoa(i)
		if !isLeft {
			currCheck = strconv.Itoa(i) + strconv.Itoa(gridSize+1-i)
		}
		moves, lastMovePointer := getPlayerMove(localMoves, currCheck)

		if lastMovePointer == nil {
			// this means line has a blank, can safely break from here
			break
		}

		lastMove := *lastMovePointer
		if currPlayer == "" {
			currPlayer = lastMove
		}
		if currPlayer != lastMove {
			// this means line is broken
			break
		}
		if currPlayer == lastMove {
			countMatch++
			localMoves = moves
		}
		if countMatch == gridSize {
			return true
		}
	}
	return false
}

func isGameOver(currMoves []string) bool {
	for i := 0; i < 2; i++ {
		if getStraightLine(currMoves, i == 0) {
			return true
		}
		if getDiagonal(currMoves, i == 0) {
			return true
		}
	}
	return false
}

func main() {
	var moves = [...]string{"b2", "a1", "c1", "b3", "a3"}
	// constants
	var emptyBoard string = "||a|b|c|\n|---|---|---|---|\n|1|⬛|⬛|⬛|\n|2|⬛|⬛|⬛|\n|3|⬛|⬛|⬛|"
  var headerMsg string = "# GITHUB IS A GAME ENGINE\n\n"
	var templateMsg string = "### PR a new move inside main.go > moves\n# CURRENT GAME:\n\n"
	var gameOverMsg string = "## GAME OVER!\n\n"
  var isMovesValid bool = !hasInvalidMove(moves[:])
  isGameWon := false

	if isMovesValid {
		// 26 is the start of the Board
		var renderBoard string = emptyBoard
		var isCircle bool = true
		var boardArr = strings.Split(renderBoard, "")
		var currMoves = []string{}

		for _, move := range moves {
			var movePiece string
			if isCircle {
				movePiece = "⭕"
			} else {
				movePiece = "❌"
			}

			// abc notations
			var xRaw = letterToNum(move[0:1])
			var xAxis = xRaw*2 + 1
			// 122 notations
			yRaw, err := strconv.Atoi(move[1:])
			if err != nil {
				fmt.Println("Error converting string to integer:", err)
				return
			}
			var yAxis = yRaw * 10

			boardArr[xAxis+yAxis+17] = movePiece
			var currMove = strconv.Itoa(xRaw) + strconv.Itoa(yRaw) + strconv.FormatBool(isCircle)
			currMoves = append(currMoves, currMove)
			if isGameOver(currMoves) {
        isGameWon = true
				break
			}
			isCircle = !isCircle
		}

		renderBoard = strings.Join(boardArr, "")
		updatedReadme := []byte("")
    if isGameWon {
      updatedReadme = []byte(headerMsg + gameOverMsg + templateMsg + renderBoard)
    } else {
      updatedReadme = []byte(headerMsg + templateMsg + renderBoard)
    }
		err := os.WriteFile("README.md", updatedReadme, 0643)
		if err != nil {
			fmt.Println("Failed to write : %v", err)
		}
	} else {
		// @TODO invalid move
		fmt.Println("Invalid Move")
	}
}
