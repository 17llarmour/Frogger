import time as t
import requests
import json
from random import *
from pygame import *


def newGame():
    url = "http://localhost/reset" + "?reset=" + "yes"
    r = requests.post(url)
    url = "http://localhost/reset" + "?reset=" + "no"
    r = requests.post(url)


def moveFrog(direction):
    url = "http://localhost/moveFrog?direction=" + direction
    frog = requests.post(url)


def getInfo():
    global lives, score
    r = requests.get("http://localhost/info")
    info = json.loads(r.text)
    lives, score = int(info[0]), int(info[1])


def getStates():
    gameGridTemp = requests.get("http://localhost/state")
    frogGridTemp = requests.get("http://localhost/frogState")
    gameGrid = json.loads(gameGridTemp.text)
    frogGrid = json.loads(frogGridTemp.text)
    # printGrid(gameGrid)
    # printGrid(frogGrid)
    drawing(gameGrid, frogGrid)


def printGrid(board):
    for i in range(13):
        print(board[i])
    print("----------SPLIT-----------")


def drawing(gameGrid, frogGrid):
    global lives, score
    screen.fill((0, 0, 0))
    draw.rect(screen, (0, 255, 0), (0, 900, 1800, 60))
    draw.rect(screen, (0, 255, 0), (0, 840, 1800, 60))
    draw.rect(screen, (0, 255, 0), (0, 480, 1800, 60))
    draw.rect(screen, (0, 0, 255), (0, 180, 1800, 300))
    drawGrid(gameGrid)
    drawGrid(frogGrid)
    writeScreen(lives, score)
    display.flip()


def drawGrid(grid):
    for y in range(13):
        for x in range(30):
            frogImage = None
            if grid[y][x] == "x":
                draw.rect(screen, (0, 255, 0), (x * 60, y * 60 + 120, 60, 60))
            else:
                frogImage = assetCollection[grid[y][x]]
            if frogImage != None:
                screen.blit(frogImage, (x * 60, y * 60 + 120))


def writeScreen(lives, score):
    fontObj = font.SysFont("Comic Sans MS", 50)
    img = fontObj.render("Score = " + str(score), True, (0, 0, 255))
    screen.blit(img, (0, 0))
    img = fontObj.render("Lives", True, (0, 0, 255))
    screen.blit(img, (1400, 0))
    x = 0
    for i in range(lives):
        img = image.load("frogForward.png").convert()
        screen.blit(img, (1560 + (i * 60 + x), 0))
        x += 20


def endScreen():
    screen.fill((255, 255, 255))
    fontObj = font.SysFont("Comic Sans MS", 120)
    img = fontObj.render("Your final score was " + str(score), True, (0, 0, 0))
    screen.blit(img, (0, 480))
    display.flip()
    t.sleep(3)


def loadAssets():
    assetCollection = {
        "f": image.load("frogForward.png").convert(),
        "bf": image.load("frogBackwards.png").convert(),
        "-1": image.load("fakeBug.png"),
        "1": image.load("log.png").convert(),
        "2": image.load("log.png").convert(),
        "3": image.load("bug1.png").convert(),
        "4": image.load("log.png").convert(),
        "5": image.load("bug1.png").convert(),
        "7": image.load("car5.png").convert(),
        "8": image.load("car4.png").convert(),
        "9": image.load("car3.png").convert(),
        "10": image.load("car2.png").convert(),
        "11": image.load("car1.png").convert(),
        " ": None
    }
    return assetCollection


if __name__ == '__main__':
    print("running client")
    score = 0
    lives = 3
    newGame()
    init()
    width = 1800
    height = 960
    screen = display.set_mode((width, height))
    assetCollection = loadAssets()
    endProgram = False
    while not endProgram:
        # t.sleep(0.25)
        for e in event.get():
            if e.type == QUIT:
                endProgram = True
            if e.type == KEYDOWN:

                if e.key == K_LEFT:
                    moveFrog("left")
                if e.key == K_RIGHT:
                    moveFrog("right")
                if e.key == K_UP or e.key == K_SPACE:
                    moveFrog("up")
                if e.key == K_DOWN:
                    moveFrog("down")

        getStates()
        getInfo()

        if lives == 0:
            endProgram = True

    endScreen()
    print("Closing Client")
