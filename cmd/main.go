package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

type pos struct {
	y int
	x int
}
type player struct {
	pos     pos
	shape   rune
	isAlive bool
	state   int8
	lives   int8
}

type enemy struct {
	pos   pos
	shape rune
	state int8
}

var player1 player
var redEnemy enemy
var purpleEnemy enemy
var blueEnemy enemy
var yellowEnemy enemy
var left pos
var right pos
var up pos
var down pos
var dir pos

var grid [31][28]rune
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
				grid[numLine][i] = '█'
				walls = append(walls, pos{numLine, i})
			case '*':
				grid[numLine][i] = '●'
				food = append(food, pos{numLine, i})
				allowed = append(allowed, pos{numLine, i})
			case '@':
				grid[numLine][i] = '◉'
				powerUp = append(powerUp, pos{numLine, i})
				allowed = append(allowed, pos{numLine, i})
			case '=':
				grid[numLine][i] = '='
				doors = append(doors, pos{numLine, i})
			case ' ':
				grid[numLine][i] = ' '
				empty = append(empty, pos{numLine, i})
				allowed = append(allowed, pos{numLine, i})
			}
		}
		numLine++
	}
}
func start() {

	player1.pos.x = 14
	player1.pos.y = 23
	player1.shape = 'C'
	player1.isAlive = true
	player1.lives = 3
	player1.state = 0

	redEnemy.pos.x = 20
	redEnemy.pos.y = 20
	redEnemy.shape = 'X'
	redEnemy.state = 0

	purpleEnemy.pos.x = 20
	purpleEnemy.pos.y = 20
	purpleEnemy.shape = 'X'
	purpleEnemy.state = 0

	blueEnemy.pos.x = 20
	blueEnemy.pos.y = 20
	blueEnemy.shape = 'X'
	blueEnemy.state = 0

	yellowEnemy.pos.x = 20
	yellowEnemy.pos.y = 20
	yellowEnemy.shape = 'X'
	yellowEnemy.state = 0

	left = pos{0, -1}
	right = pos{0, 1}
	up = pos{-1, 0}
	down = pos{1, 0}
	dir = left

}

func input(wg *sync.WaitGroup) {

	defer wg.Done()
	for {
		time.Sleep(time.Millisecond * 17)
		reader := bufio.NewReader(os.Stdin) //doesnt work need to read input without pressin eneter
		input, _ := reader.ReadByte()

		if input == 'a' && grid[player1.pos.y][player1.pos.x-1] != '█' {
			dir = left
			continue
		}

		if input == 'd' && grid[player1.pos.y][player1.pos.x+1] != '█' {
			dir = right
			continue
		}

		if input == 'w' && grid[player1.pos.y-1][player1.pos.x] != '█' {
			dir = up
			continue
		}

		if input == 's' && grid[player1.pos.y+1][player1.pos.x] != '█' {
			dir = down
			continue
		}
	}

}

func update(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if grid[player1.pos.y+dir.y][player1.pos.x+dir.x] != '█' {

			grid[player1.pos.y][player1.pos.x] = ' '
			player1.pos.x += dir.x
			player1.pos.y += dir.y
		}

		grid[player1.pos.y][player1.pos.x] = player1.shape

		time.Sleep(time.Millisecond * 150)
	}
}

func draw(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		buf := new(bytes.Buffer)
		for i := 0; i < 31; i++ {
			for j := 0; j < 28; j++ {
				buf.WriteString(string(grid[i][j]))
			}
			buf.WriteString("\n")
		}
		buf.WriteString("\033[H\033[1:1H")
		buf.WriteTo(os.Stdout)
		time.Sleep(time.Millisecond * 17)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	state, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error setting raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), state)

	start()
	makeGrid()
	go input(&wg)
	go update(&wg)
	go draw(&wg)
	wg.Wait()
}
