package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

type pos struct {
	y int
	x int
}
type entity struct {
	pos     pos
	name    string
	shape   rune
	isAlive bool
	state   int8
	lives   int8
}

var grid [31][28]string
var walls []pos
var doors []pos
var food []pos
var powerUp []pos
var allowed []pos
var empty []pos

func makeGrid() {
	file, err := os.Open("map.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numLine int = 0

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {

			switch line[i] {
			case '#':
				grid[numLine][i] = "█"
				walls = append(walls, pos{numLine, i})
			case '*':
				grid[numLine][i] = "●"
				food = append(food, pos{numLine, i})
				allowed = append(allowed, pos{numLine, i})
			case '@':
				grid[numLine][i] = "◉"
				powerUp = append(powerUp, pos{numLine, i})
				allowed = append(allowed, pos{numLine, i})
			case '=':
				grid[numLine][i] = "="
				doors = append(doors, pos{numLine, i})
			case ' ':
				grid[numLine][i] = " "
				empty = append(empty, pos{numLine, i})
				allowed = append(allowed, pos{numLine, i})
			}
		}
		numLine++
	}
}
func start() {
	var player entity

	player.pos.x = 20
	player.pos.y = 20
	player.name = "player"
	player.shape = 'C'
	player.isAlive = true
	player.lives = 3
	player.state = 0
}
func input() {

}
func update() {

}
func draw() {
	buf := new(bytes.Buffer)
	for i := 0; i < 31; i++ {
		for j := 0; j < 28; j++ {
			buf.WriteString(grid[i][j])
		}
		buf.WriteString("\n")
	}
	fmt.Print("\033[H\033[1:1H")
	fmt.Print(buf.String())
	buf.Reset()
}

func main() {

	start()
	makeGrid()
	go input()

	for {
		update()
		draw()
		time.Sleep(time.Millisecond * 17)
	}

}
