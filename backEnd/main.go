package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var gameGrid [][]string
var frogGrid [][]string
var lives = 3
var animalsPlaced = 0

func main() {
	gameGrid = buildGrid()
	frogGrid = buildGrid()
	for {
		for round := 0; lives > 0; round++ {
			//if round%20 == 0 {
			//	moveCars()
			//}
			moveCars(round)
			moveAnimals(round)
			if round%60 == 0 {
				addCars()
			}
			if animalsPlaced < 3 && round%5 == 0 { //round%80 == 0 {
				addWaterAnimals()
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
		chance := rand.Intn(3)
		if i%2 == 0 && chance == 1 {
			gameGrid[i][0] = strconv.Itoa(i)
		} else if chance == 1 {
			gameGrid[i][29] = strconv.Itoa(i)
		}
	}
}

func addWaterAnimals() {
	for i := 5; i > 0; i-- {
		chance := rand.Intn(3)
		chance = 1
		if i%2 == 0 && (chance == 1 || animalsPlaced < 3) {
			gameGrid[i][0] = strconv.Itoa(i)
			animalsPlaced += 1
		} else if chance == 1 || animalsPlaced < 2 {
			gameGrid[i][29] = strconv.Itoa(i)
			animalsPlaced += 1
		}
	}
	if animalsPlaced == 2 {
		animalsPlaced = 0
	}
}

func moveAnimals(round int) {
	for i := 5; i > 0; i-- {
		if round%20 == 0 && i == 5 {
			moveCarLeft(i)
		} else if round%15 == 0 && i == 4 {
			moveCarRight(i)
		} else if round%10 == 0 && i == 3 {
			moveCarLeft(i)
		} else if round%8 == 0 && i == 2 {
			moveCarRight(i)
		} else if round%5 == 0 && i == 1 {
			moveCarLeft(i)
		}
	}
}

func moveCars(round int) {
	for i := 11; i > 6; i-- {
		if round%20 == 0 && i == 11 {
			moveCarLeft(i)
		} else if round%15 == 0 && i == 10 {
			moveCarRight(i)
		} else if round%10 == 0 && i == 9 {
			moveCarLeft(i)
		} else if round%8 == 0 && i == 8 {
			moveCarRight(i)
		} else if round%5 == 0 && i == 7 {
			moveCarLeft(i)
		}
		//if i%2 == 0 {
		//	moveCarRight(i)
		//} else {
		//	moveCarLeft(i)
		//}
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
		if gameGrid[y][x-1] == " " {
			return true
		}
	}
	if direction == "right" {
		if gameGrid[y][x+1] == " " {
			return true
		}
	}
	if direction == "up" {
		if gameGrid[y-1][x] == " " {
			return true
		}
	}
	return true
}

func moveFrogLeft() {
	for i := 1; i < 13; i++ {
		for x := 1; x < 30; x++ {
			if i > 5 {
				if frogGrid[i][x] == "f" && moveFrogCheck("left", i, x) {
					frogGrid[i][x-1] = frogGrid[i][x]
					frogGrid[i][x] = " "
				}
			}
		}
	}
}

func moveFrogRight() {
	for i := 1; i < 13; i++ {
		for x := 0; x < 29; x++ {
			if i > 5 {
				if frogGrid[i][x] == "f" && moveFrogCheck("right", i, x) {
					frogGrid[i][x+1] = frogGrid[i][x]
					frogGrid[i][x] = " "
				}
			}
		}
	}
}

func moveFrogUp() {
	for i := 1; i < 13; i++ {
		for x := 0; x < 30; x++ {
			if i > 5 {
				if frogGrid[i][x] == "f" && moveFrogCheck("up", i, x) {
					frogGrid[i-1][x] = frogGrid[i][x]
					frogGrid[i][x] = " "
				}
			} else {

			}
		}
	}
}
