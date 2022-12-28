package main

import (
	"fmt"
	"strconv"
	"time"
)

var gameGrid [][]string
var frogGrid [][]string
var lives = 3

func main() {
	gameGrid = buildGrid()
	frogGrid = buildGrid()
	for {
		for round := 0; lives > 0; round++ {
			if round%10 == 0 {
				moveCars()
			}
			if round%40 == 0 {
				addCars()
			}
			printGrid(gameGrid)
			time.Sleep(25 * time.Millisecond)
		}
	}
}

func buildGrid() [][]string {
	var tempGrid [][]string
	for z := 0; z < 13; z++ {
		var x []string
		for i := 0; i < 30; i++ { // Space between each item in the array - means the bullet can not collide with anything
			x = append(x, " ")
		}
		tempGrid = append(tempGrid, x)
	}
	return tempGrid
}

func printGrid(grid [][]string) {
	for i := 0; i < 13; i++ {
		printLn := "["
		for x := 0; x < 30; x++ {
			printLn += grid[i][x]
		}
		fmt.Println(printLn + "]")
	}
	fmt.Println("-------------SPLIT--------------")
}

func addCars() {
	for i := 11; i > 6; i-- {
		if i%2 == 0 {
			gameGrid[i][0] = strconv.Itoa(i)
		} else {
			gameGrid[i][29] = strconv.Itoa(i)
		}
	}
}

func moveCars() {
	for i := 11; i > 6; i-- {
		if i%2 == 0 {
			moveCarRight(i)
		} else {
			moveCarLeft(i)
		}
	}
}

func moveCarLeft(pos int) {
	for i := 1; i < 29; i++ {
		gameGrid[pos][i] = gameGrid[pos][i+1]
		gameGrid[pos][i+1] = " "
	}
	if gameGrid[pos][0] != " " {
		gameGrid[pos][0] = " "
	}
}

func moveCarRight(pos int) {
	for i := 28; i > 0; i-- {
		gameGrid[pos][i] = gameGrid[pos][i-1]
		gameGrid[pos][i-1] = " "
	}
	if gameGrid[pos][29] != " " {
		gameGrid[pos][29] = " "
	}
}

func moveFrogCheck(direction string, y, x int) bool {
	if direction == "left" {

	}
	return true
}

func moveFrogLeft() {
	for i := 1; i < 13; i++ {
		for x := 0; x < 30; x++ {
			if gameGrid[i][x] == "f" {
				moveFrogCheck("left", i, x)
			}
		}
	}
}

func moveFrogRight() {

}

func moveFrogUp() {

}
