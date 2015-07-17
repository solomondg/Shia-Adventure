/* _______  __   __  ___   _______  _______  __     _______  ______   __   __  _______  __    _  _______  __   __  ______    _______
|       ||  | |  ||   | |   _   ||       ||  |   |   _   ||      | |  | |  ||       ||  |  | ||       ||  | |  ||    _ |  |       |
|  _____||  |_|  ||   | |  |_|  ||  _____||__|   |  |_|  ||  _    ||  |_|  ||    ___||   |_| ||_     _||  | |  ||   | ||  |    ___|
| |_____ |       ||   | |       || |_____        |       || | |   ||       ||   |___ |       |  |   |  |  |_|  ||   |_||_ |   |___
|_____  ||       ||   | |       ||_____  |       |       || |_|   ||       ||    ___||  _    |  |   |  |       ||    __  ||    ___|
 _____| ||   _   ||   | |   _   | _____| |       |   _   ||       | |     | |   |___ | | |   |  |   |  |       ||   |  | ||   |___
|_______||__| |__||___| |__| |__||_______|       |__| |__||______|   |___|  |_______||_|  |__|  |___|  |_______||___|  |_||_______|

*/

package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	b_black   string = "\033[40m"
	b_red     string = "\033[41m"
	b_green   string = "\033[42m"
	b_gold    string = "\033[43m"
	b_cyan    string = "\033[44m"
	b_magenta string = "\033[45m"
	b_aqua    string = "\033[46m"
	b_white   string = "\033[47m"
	b_default string = "\033[49m"
)

var (
	b  = b_black
	r  = b_red
	gr = b_green
	g  = b_gold
	m  = b_magenta
	a  = b_aqua
	w  = b_white
	d  = b_default
)

const (
	xSize = 8
	ySize = 8
)

type buffer struct {
	cells [xSize][ySize]string
}

func (b *buffer) init(fillChar string) {
	var x, y byte
	for x = 1; x < xSize; x++ {
		for y = 1; y < ySize; y++ {
			b.cells[x][y] = fillChar
		}
	}
}

func (b *buffer) clearBuffer(fillChar string) {
	var x, y byte
	for x = 1; x < xSize; x++ {
		for y = 1; y < ySize; y++ {
			b.cells[x][y] = fillChar
		}
	}
}

func (b *buffer) clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Printf("\033[39m\033[49m")
}

func (b *buffer) drawFrame() {
	var x, y byte
	for x = 1; x < xSize; x++ {
		for y = 1; y < ySize; y++ {
			fmt.Printf("%v", b.cells[x][y])
		}
		fmt.Printf("\n")
	}

}

func (b *buffer) fillCells(xStarting, xEnding, yStarting, yEnding byte, fillChar, endChar string) {
	var x, y byte = xStarting, yStarting
	b.cells[x][y] = endChar
	for x = xStarting; x < xEnding; x++ {
		b.cells[x][y] = endChar
		for y = yStarting; y < yEnding; y++ {
			b.cells[x][y] = fillChar
		}
		b.cells[x][y] = endChar
	}
	b.cells[x][y] = endChar
}

func main() {
	var mainbuffer buffer
	mainbuffer.clearScreen()
	mainbuffer.init(" ")
	mainbuffer.clearScreen()
	mainbuffer.clearBuffer(" ")
	mainbuffer.fillCells(2, 6, 4, 6, fmt.Sprintf("%v", r), fmt.Sprintf("%v", d))
	mainbuffer.drawFrame()
}
