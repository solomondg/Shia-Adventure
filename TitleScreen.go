/* _______  __   __  ___  _______  __  _______    _______  ______   __   __  _______  __    _  _______  __   __  ______    _______
|       ||  | |  ||   ||   _   ||  ||       |  |   _   ||      | |  | |  ||       ||  |  | ||       ||  | |  ||    _ |  |       |
|  _____||  |_|  ||   ||  |_|  ||__||  _____|  |  |_|  ||  _    ||  |_|  ||    ___||   |_| ||_     _||  | |  ||   | ||  |    ___|
| |_____ |       ||   ||       |    | |_____   |       || | |   ||       ||   |___ |       |  |   |  |  |_|  ||   |_||_ |   |___
|_____  ||       ||   ||       |    |_____  |  |       || |_|   ||       ||    ___||  _    |  |   |  |       ||    __  ||    ___|
 _____| ||   _   ||   ||   _   |     _____| |  |   _   ||       | |     | |   |___ | | |   |  |   |  |       ||   |  | ||   |___
|_______||__| |__||___||__| |__|    |_______|  |__| |__||______|   |___|  |_______||_|  |__|  |___|  |_______||___|  |_||_______|
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
	ySize = 144
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

func (b *buffer) clearScreen(rep byte) {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	fmt.Printf("\033[39m\033[49m\n")
	if rep > 0 {
		b.clearScreen(rep - 1)
	}
}

func dl(t float32) {
	time.Sleep(time.Duration(t) * time.Millisecond)
}

func (b *buffer) drawFrame() {
	var x, y byte
	for x = 1; x < xSize; x++ {
		for y = 1; y < ySize; y++ {
			fmt.Printf("%v", b.cells[x][y])
			//	time.Sleep(50 * time.Microsecond)
		}
		fmt.Printf("\n")
		//	time.Sleep(50 * time.Microsecond)
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

func (b *buffer) outline(xVertex, yVertex, width, height, thickness byte, color string, borderchar string) {
	var x, y byte = xVertex, yVertex
	for y = yVertex; y < height-1; y++ { // Draw top line, left-right
		b.cells[x][y] = fmt.Sprintf("%v%v%v", color, borderchar, d)
	}
	b.cells[x][y] = fmt.Sprintf("%v", d) // Prevent color from bleeding

	for x = xVertex; x < width-1; x++ { // Cursor's on the right; draw right line top-bottom
		b.cells[x][y] = fmt.Sprintf("%v%v%v%v", color, borderchar, borderchar, d)
	}
	b.cells[x][y] = fmt.Sprintf("%v%v%v", borderchar, borderchar, d) // Anti-bleed

	for y = height - 1; y > yVertex; y-- { // Now we're at the bottom. Draw bottom line, right to left.
		b.cells[x][y] = fmt.Sprintf("%v%v", color, borderchar)
	}
	b.cells[x][y] = fmt.Sprintf("%v", d) // it's not really necessary, but anti-bleed
	//x, y = 1, 1
	x, y = xVertex, yVertex             //reset "cursor"
	for x = xVertex; x < width-1; x++ { // left line, top-bottom
		b.cells[x][y] = fmt.Sprintf("%v%v%v%v", color, borderchar, borderchar, d)
	}
	b.cells[x][y] = fmt.Sprintf("%v%v%v", color, borderchar, d) //anti-bleed, just out of habit
	if thickness > 1 {
		b.outline(xVertex-1, yVertex-1, width-1, height-1, thickness-1, color, borderchar)
	}
}

func (b *buffer) stringBuffer(xpos, ypos byte, direction rune, inputString string) {
	var x, y byte = xpos, ypos
	//	fmt.Println(inputString)

	for a, z := range inputString {
		if direction == 'v' {
			b.cells[x+byte(a)][y] = string(z)
		} else {
			b.cells[x][y+byte(a)] = string(z)
		}
	}
}

const (
	line1 string = " _______  __   __  ___  _______  __  _______    _______  ______   __   __  _______  __    _  _______  __   __  ______    _______ "
	line2 string = "|       ||  | |  ||   ||   _   ||  ||       |  |   _   ||      | |  | |  ||       ||  |  | ||       ||  | |  ||    _ |  |       |"
	line3 string = "|  _____||  |_|  ||   ||  |_|  ||__||  _____|  |  |_|  ||  _    ||  |_|  ||    ___||   |_| ||_     _||  | |  ||   | ||  |    ___|"
	line4 string = "| |_____ |       ||   ||       |    | |_____   |       || | |   ||       ||   |___ |       |  |   |  |  |_|  ||   |_||_ |   |___ "
	line5 string = "|_____  ||       ||   ||       |    |_____  |  |       || |_|   ||       ||    ___||  _    |  |   |  |       ||    __  ||    ___|"
	line6 string = " _____| ||   _   ||   ||   _   |     _____| |  |   _   ||       | |     | |   |___ | | |   |  |   |  |       ||   |  | ||   |___ "
	line7 string = "|_______||__| |__||___||__| |__|    |_______|  |__| |__||______|   |___|  |_______||_|  |__|  |___|  |_______||___|  |_||_______|"
)

func (b *buffer) titleDraw(top byte) {
	//var x byte
	/*
		for x = 0; x < line; x++ {
			defer b.stringBuffer(x+2, 8, 'h', titleLines[x])
		}
	*/

	b.stringBuffer(2, 8, 'h', titleLines[top])
	if top != 0 {
		b.titleDraw(top - 1)

	}
}

func (b *buffer) clrDrw(rep byte) {
	b.clearScreen(rep)
	b.drawFrame()
}

var (
	bgc        string = fmt.Sprintf("%v ", g)
	mainbuffer buffer
	titleLines [7]string = [7]string{line1, line2, line3, line4, line5, line6, line7}
)

func main() {
	mainbuffer.clearScreen(2)
	mainbuffer.init(bgc)
	mainbuffer.clearBuffer(bgc)
	//mainbuffer.fillCells(10, 22, 8, 27, fmt.Sprintf("%v ", r), fmt.Sprintf("%v%v", d, bgc))
	// mainbuffer.stringBuffer(5, 5, 'h', "ayy")

	//	var z byte
	//	for z = 0; z < 7; z++ {
	//		mainbuffer.clearBuffer(bgc)
	//		mainbuffer.outline(1, 1, xSize, ySize, 1, w, " ")
	//		mainbuffer.titleDraw(z)
	//		dl(100)
	//		mainbuffer.clrDrw(3)
	//	}

	mainbuffer.outline(1, 1, xSize, ySize, 1, w, " ")
	mainbuffer.titleDraw(2)
	mainbuffer.clearScreen(2)
	mainbuffer.drawFrame()
	/*
		bgc = fmt.Sprintf("%v ", c)
		mainbuffer.clearScreen(3)
		mainbuffer.clearBuffer(fmt.Sprintf("%v ", gr))
		mainbuffer.outline(1, 1, xSize, ySize, 1, w, " ")
		mainbuffer.fillCells(7, 13, 20, 40, fmt.Sprintf("%v %v", b, gr), fmt.Sprintf("%v%v ", d, gr))
		mainbuffer.stringBuffer(20, 50, 'v', "lmao")
		mainbuffer.drawFrame()
	*/
}
