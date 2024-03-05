package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"

	tm "github.com/buger/goterm"
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
	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(30|tm.PCT, 20, 0)

	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box, buf.String())

	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box.String(), 40|tm.PCT, 40|tm.PCT))

}

func main() {

	start()
	makeGrid()
	go input()
	tm.Clear() // Clear current screen
	for {
		tm.MoveCursor(1, 1)
		tm.Flush()

		update()
		draw()
		time.Sleep(time.Millisecond * 17)
	}

}
