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
	"time"
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
	c  = b_cyan
	m  = b_magenta
	a  = b_aqua
	w  = b_white
	d  = b_default
)

const (
	xSize = 32
	ySize = 108
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
	fmt.Printf("\033[39m\033[49m\n")
}

func dl() {
	time.Sleep(500 * time.Millisecond)
}

func (b *buffer) drawFrame() {
	var x, y byte
	for x = 1; x < xSize; x++ {
		for y = 1; y < ySize; y++ {
			fmt.Printf("%v", b.cells[x][y])
			time.Sleep(1 * time.Millisecond)
		}
		fmt.Printf("\n")
		time.Sleep(1 * time.Millisecond)
	}

}

func (b *buffer) fillCells(xStarting, xEnding, yStarting, yEnding byte, fillChar, endChar string) {
	var x, y byte = xStarting, yStarting
	b.cells[x][y] = endChar
	for x = xStarting; x < xEnding+1; x++ {
		b.cells[x][y] = endChar
		for y = yStarting; y < yEnding; y++ {
			b.cells[x][y] = fillChar
		}
		b.cells[x][y] = endChar
	}
	//b.cells[x][y] = endChar
}

func (b *buffer) outline(width, height byte, color string, borderchar string) {
	var x, y byte = 1, 1
	for y = 1; y < height-1; y++ { // Draw top line, left-right
		b.cells[x][y] = fmt.Sprintf("%v%v%v", color, borderchar, d)
	}
	b.cells[x][y] = fmt.Sprintf("%v%v", d) // Prevent color from bleeding

	for x = 0; x < width-1; x++ { // Cursor's on the right; draw right line top-bottom
		b.cells[x][y] = fmt.Sprintf("%v%v%v%v", color, borderchar, borderchar, d)
	}
	b.cells[x][y] = fmt.Sprintf("%v%v%v", borderchar, borderchar, d) // Anti-bleed

	for y = height - 1; y > 0; y-- { // Now we're at the bottom. Draw bottom line, right to left.
		b.cells[x][y] = fmt.Sprintf("%v%v", color, borderchar)
	}
	b.cells[x][y] = fmt.Sprintf("%v%v", d) // it's not really necessary, but anti-bleed
	//x, y = 1, 1
	x, y = 1, 1                   //reset "cursor"
	for x = 1; x < width-1; x++ { // left line, top-bottom
		b.cells[x][y] = fmt.Sprintf("%v%v%v%v", color, borderchar, borderchar, d)
	}
	b.cells[x][y] = fmt.Sprintf("%v%v%v", color, borderchar, d) //anti-bleed, just out of habit
}

const (
	line1 string = ""
	line2 string = ""
	line3 string = ""
	line4 string = ""
	line5 string = ""
	line6 string = ""
	line7 string = ""
	line8 string = ""
)

func main() {
	var bgc string = fmt.Sprintf("%v ", c)
	var mainbuffer buffer
	mainbuffer.clearScreen()
	mainbuffer.init(bgc)
	mainbuffer.clearScreen()
	mainbuffer.clearBuffer(bgc)
	bgc = fmt.Sprintf("%v ", a)
	//	mainbuffer.fillCells(8, 18, 9, 18, fmt.Sprintf("%v%v", gr, bgc), fmt.Sprintf("%v%v", d, bgc))
	//	mainbuffer.fillCells(2, 8, 18, 23, fmt.Sprintf("%v%v", w, bgc), fmt.Sprintf("%v%v", d, bgc))
	mainbuffer.fillCells(3, 5, 3, 5, fmt.Sprintf("%v%v", gr, bgc), fmt.Sprintf("%v%v", d, bgc))
	mainbuffer.fillCells(10, 22, 8, 27, fmt.Sprintf("%v%v", r, bgc), fmt.Sprintf("%v%v", d, bgc))
	mainbuffer.outline(xSize, ySize, w, " ")
	mainbuffer.drawFrame()
}
