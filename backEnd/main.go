package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var gameGrid [][]string
var frogGrid [][]string
var lives = 3
var score = 0

func main() {
	go runServer()
	gameGrid = buildGrid()
	frogGrid = buildGrid()
	addWinningPlatforms()
	for {
		for round := 0; lives > 0; round++ {
			moveCars(round)
			moveAnimals(round)
			if round%40 == 0 {
				addCars()
			}
			if round%80 == 0 {
				addWaterAnimals()
			}
			//printGrid(gameGrid)
			time.Sleep(25 * time.Millisecond)
		}
	}
}

func buildGrid() [][]string {
	var tempGrid [][]string
	for z := 0; z < 13; z++ {
		var x []string
		for i := 0; i < 36; i++ { // Space between each item in the array - means there is definitely space for the frog to go
			x = append(x, " ")
		}
		tempGrid = append(tempGrid, x)
	}
	return tempGrid
}

func addWinningPlatforms() {
	for i := 3; i < 34; i++ {
		if (i+1)%6 == 0 {
		} else {
			gameGrid[0][i] = "x"
		}
	}
}

func printGrid(grid [][]string) {
	for i := 0; i < 13; i++ {
		printLn := "["
		for x := 3; x < 34; x++ {
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
			gameGrid[i][2] = strconv.Itoa(i)
		} else if chance == 1 {
			gameGrid[i][33] = strconv.Itoa(i)
		}
	}
}

func addWaterAnimals() {
	for i := 5; i > 0; i-- {
		chance := rand.Intn(2)
		fake := rand.Intn(5)
		if i%2 == 0 && chance == 1 {
			gameGrid[i][0] = strconv.Itoa(i)
			gameGrid[i][1] = strconv.Itoa(i)
			gameGrid[i][2] = strconv.Itoa(i)
		} else if chance == 1 && (fake != 1 || i == 1) {
			gameGrid[i][34] = strconv.Itoa(i)
			gameGrid[i][35] = strconv.Itoa(i)
		} else if chance == 1 && fake == 1 && i != 1 {
			gameGrid[i][34] = "-1"
			gameGrid[i][35] = "-1"
		}
	}
}

func moveAnimals(round int) {
	for i := 5; i > 0; i-- {
		if round%18 == 0 && i == 5 {
			moveCarLeft(i)
		} else if round%15 == 0 && i == 4 {
			moveCarRight(i)
		} else if round%12 == 0 && i == 3 {
			moveCarLeft(i)
		} else if round%10 == 0 && i == 2 {
			moveCarRight(i)
		} else if round%9 == 0 && i == 1 {
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
	}
}

func moveCarLeft(pos int) {
	for i := 2; i < 35; i++ {
		gameGrid[pos][i] = gameGrid[pos][i+1]
		gameGrid[pos][i+1] = " "
		if pos > 6 && frogGrid[pos][i] == "f" && gameGrid[pos][i] != " " {
			frogDeathCars()
		} else if pos < 6 && frogGrid[pos][i+1] == "f" && gameGrid[pos][i] != " " {
			moveFrogLeft()
		}
	}
	if frogGrid[pos][3] == "f" && gameGrid[pos][3] == " " {
		frogDeathCars()
	}
	if gameGrid[pos][2] != " " {
		gameGrid[pos][2] = " "
	}
}

func moveCarRight(pos int) {
	for i := 35; i > 0; i-- {
		gameGrid[pos][i] = gameGrid[pos][i-1]
		gameGrid[pos][i-1] = " "
		if pos > 6 && frogGrid[pos][i] == "f" && gameGrid[pos][i] != " " {
			frogDeathCars()
		} else if pos < 6 && frogGrid[pos][i-1] == "f" && gameGrid[pos][i] != " " {
			moveFrogRight()
		}
	}
	if frogGrid[pos][32] == "f" && gameGrid[pos][32] == " " {
		frogDeathCars()
	}
	if gameGrid[pos][33] != " " {
		gameGrid[pos][33] = " "
	}
}

func moveFrogCheckCars(direction string, y, x int) bool {
	if direction == "left" {
		if gameGrid[y][x-1] == " " {
			return true
		}
	} else if direction == "right" {
		if gameGrid[y][x+1] == " " {
			return true
		}
	} else if direction == "up" {
		if gameGrid[y-1][x] == " " {
			return true
		}
	} else if direction == "down" {
		if gameGrid[y+1][x] == " " {
			return true
		}
	}
	return false
}

func moveFrogLeft() {
	for i := 1; i < 13; i++ {
		for x := 4; x < 33; x++ {
			if i > 5 {
				local := moveFrogCheckCars("left", i, x)
				if frogGrid[i][x] == "f" && local {
					frogGrid[i][x-1] = frogGrid[i][x]
					frogGrid[i][x] = " "
					return
				} else if frogGrid[i][x] == "f" && !local {
					frogGrid[i][x-1] = "d"
					frogGrid[i][x] = " "
					frogDeathCars()
					return
				}
			} else if i > 0 {
				if gameGrid[i][x-1] == " " && frogGrid[i][x] == "f" {
					frogGrid[i][x-1] = "d"
					frogGrid[i][x] = " "
					frogDeathCars()
					return
				} else if gameGrid[i][x-1] != " " && frogGrid[i][x] == "f" {
					frogGrid[i][x-1] = frogGrid[i][x]
					frogGrid[i][x] = " "
					return
				}
			}
		}
	}
}

func moveFrogRight() {
	for i := 1; i < 13; i++ {
		for x := 3; x < 32; x++ {
			if i > 5 {
				local := moveFrogCheckCars("right", i, x)
				if frogGrid[i][x] == "f" && local {
					frogGrid[i][x+1] = frogGrid[i][x]
					frogGrid[i][x] = " "
					return
				} else if frogGrid[i][x] == "f" && !local {
					frogGrid[i][x+1] = "d"
					frogGrid[i][x] = " "
					frogDeathCars()
					return
				}
			} else if i > 0 {
				if gameGrid[i][x+1] == " " && frogGrid[i][x] == "f" {
					frogGrid[i][x+1] = "d"
					frogGrid[i][x] = " "
					frogDeathCars()
					return
				} else if gameGrid[i][x+1] != " " && frogGrid[i][x] == "f" {
					frogGrid[i][x+1] = frogGrid[i][x]
					frogGrid[i][x] = " "
					return
				}
			}
		}
	}
}

func moveFrogUp() {
	var local bool
	for i := 0; i < 13; i++ {
		for x := 3; x < 34; x++ {
			if i > 6 {
				local = moveFrogCheckCars("up", i, x)
				if frogGrid[i][x] == "f" && local {
					frogGrid[i-1][x] = frogGrid[i][x]
					frogGrid[i][x] = " "
					score += 10
					return
				} else if frogGrid[i][x] == "f" && !local {
					frogGrid[i-1][x] = "d"
					frogGrid[i][x] = " "
					frogDeathCars()
					return
				}
			} else if i > 1 {
				local = moveFrogCheckCars("up", i, x)
				if frogGrid[i][x] == "f" && !local && gameGrid[i-1][x] != "-1" {
					frogGrid[i-1][x] = frogGrid[i][x]
					frogGrid[i][x] = " "
					score += 10
					return
				} else if frogGrid[i][x] == "f" && (local || gameGrid[i-1][x] == "-1") {
					frogGrid[i-1][x] = "d"
					frogGrid[i][x] = " "
					frogDeathCars()
					return
				}
			} else if i == 1 {
				local = moveFrogCheckCars("up", i, x)
				if frogGrid[i][x] == "f" && local {
					frogGrid[i][x] = " "
					gameGrid[0][x] = "bf"
					frogGrid[0][x] = " "
					frogGrid[12][14] = "f"
					score += 200
					winCheck()
					return

				}
			}
		}
	}
}

func moveFrogDown() {
	var local bool
	for i := 0; i < 12; i++ {
		for x := 3; x < 34; x++ {
			if i > 4 {
				local = moveFrogCheckCars("down", i, x)
				if frogGrid[i][x] == "f" && local {
					frogGrid[i+1][x] = frogGrid[i][x]
					frogGrid[i][x] = " "
					return
				} else if frogGrid[i][x] == "f" && !local {
					frogGrid[i+1][x] = "d"
					frogGrid[i][x] = " "
					frogDeathCars()
					return
				}
			} else if i > 0 {
				local = moveFrogCheckCars("down", i, x)
				if frogGrid[i][x] == "f" && !local && gameGrid[i+1][x] != "-1" {
					frogGrid[i+1][x] = frogGrid[i][x]
					frogGrid[i][x] = " "
					return
				} else if frogGrid[i][x] == "f" && (local || gameGrid[i+1][x] == "-1") {
					frogGrid[i+1][x] = "d"
					frogGrid[i][x] = " "
					frogDeathCars()
					return
				}
			}
		}
	}
}

func winCheck() {
	var total = 0
	for i := 5; i < 33; i += 6 {
		if gameGrid[0][i] == "bf" {
			total += 1
		}
	}
	if total == 5 {
		score += 1000
		gameGrid = buildGrid()
		frogGrid = buildGrid()
		addWinningPlatforms()
		frogGrid[12][14] = "f"
	}
}

func frogDeathCars() {
	lives -= 1
	//time.Sleep(2 * time.Second)
	frogGrid = buildGrid()
	addWinningPlatforms()
	frogGrid[12][14] = "f"
}

func runServer() {
	http.HandleFunc("/state", getState)
	http.HandleFunc("/frogState", getFrogState)
	http.HandleFunc("/moveFrog", getNewFrog)
	http.HandleFunc("/info", getInfo)
	http.HandleFunc("/reset", resetCheck)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getState(w http.ResponseWriter, r *http.Request) {
	var localGrid [][]string
	for i := 0; i < 13; i++ {
		var tempLine []string
		for x := 3; x < 33; x++ {
			tempLine = append(tempLine, gameGrid[i][x])
		}
		localGrid = append(localGrid, tempLine)
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(localGrid)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getFrogState(w http.ResponseWriter, r *http.Request) {
	var localGrid [][]string
	for i := 0; i < 13; i++ {
		var tempLine []string
		for x := 3; x < 33; x++ {
			tempLine = append(tempLine, frogGrid[i][x])
		}
		localGrid = append(localGrid, tempLine)
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(localGrid)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNewFrog(w http.ResponseWriter, r *http.Request) {
	frogDirection := r.URL.Query()["direction"]
	if frogDirection[0] == "right" {
		moveFrogRight()
	} else if frogDirection[0] == "left" {
		moveFrogLeft()
	} else if frogDirection[0] == "up" {
		moveFrogUp()
	} else if frogDirection[0] == "down" {
		moveFrogDown()
	}
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	localLives := strconv.Itoa(lives)
	localScore := strconv.Itoa(score)
	write := [2]string{localLives, localScore}
	err := json.NewEncoder(w).Encode(write)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func resetCheck(w http.ResponseWriter, r *http.Request) {
	reset := r.URL.Query()["reset"]

	if reset[0] == "yes" {
		lives = 3
		score = 0
		gameGrid = buildGrid()
		frogGrid = buildGrid()
		addWinningPlatforms()
		frogGrid[12][14] = "f"
	}
}
