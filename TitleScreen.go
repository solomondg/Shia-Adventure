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

var shiagreyscale_source = []string{

	"00000000111111111110000000000000",
	"00000001111111111110000000000000",
	"00000011111111001110000000000000",
	"00000001111000000011000000000000",
	"00000111100000000001000000000000",
	"00000111100000000001100000000000",
	"00000111100000000001100000000000",
	"00001111100000000001100000000000",
	"00000111100000000001110000000000",
	"00000111011100001110110000000000",
	"00011111110010000001110000000000",
	"00000111101110011110110000000000",
	"00000111010101000010110000000000",
	"00001010000000000000111000000000",
	"00001010000000000000000000000000",
	"00001001000000000000101100000000",
	"00001101000000000000100100000000",
	"00001111000001110000101110000000",
	"00011111100011010000111110000000",
	"00001111100111111100111111000000",
	"00111111101111111101111110000000",
	"00111111111001010111111100000000",
	"01111111111000100011111000000000",
	"01111111111001100111111000000000",
	"01111111111111110111111100000000",
	"00111111111111111111111100000000",
	"00111111111111111111111100000000",
	"00011111011111111111111000000000",
	"00011111011111111111110000000000",
	"00001111001111011101110000000000",
	"00011111000111110001110000000000",
	"00111111000000000001111000000000",
}

func sourcetoarray(input []string) [32][32]byte {
	var array [32][32]byte
	xPos, yPos := 0, 0
	xPos += 1
	yPos += 1
	for xPos, x := range input {
		for yPos, y := range x {
			array[yPos][xPos] = byte(y)
		}
	}
	return array
}

func (b *buffer) arrayBuffer(xPos, yPos, xSize, ySize byte, array [][]byte) {
	for x = xPos; x < xSize; x++ {

	}
}

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
	xSize = 40
	ySize = 160
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

var (
	s     string = fmt.Sprintf("%v %v", g, d)
	line1 string = fmt.Sprintf(" _______  __   __  ___  _______  __  _______    _______  ______   __   __  _______  __    _  _______  __   __  ______    _______ ")
	line2 string = fmt.Sprintf("|       ||  | |  ||   ||   _   ||  ||       |  |   _   ||      | |  | |  ||       ||  |  | ||       ||  | |  ||    _ |  |       |")
	line3 string = fmt.Sprintf("|  _____||  |_|  ||   ||  |_|  ||__||  _____|  |  |_|  ||  _    ||  |_|  ||    ___||   |_| ||_     _||  | |  ||   | ||  |    ___|")
	line4 string = fmt.Sprintf("| |_____ |       ||   ||       |    | |_____   |       || | |   ||       ||   |___ |       |  |   |  |  |_|  ||   |_||_ |   |___ ")
	line5 string = fmt.Sprintf("|_____  ||       ||   ||       |    |_____  |  |       || |_|   ||       ||    ___||  _    |  |   |  |       ||    __  ||    ___|")
	line6 string = fmt.Sprintf(" _____| ||   _   ||   ||   _   |     _____| |  |   _   ||       | |     | |   |___ | | |   |  |   |  |       ||   |  | ||   |___ ")
	line7 string = fmt.Sprintf("|_______||__| |__||___||__| |__|    |_______|  |__| |__||______|   |___|  |_______||_|  |__|  |___|  |_______||___|  |_||_______|")
)

func (b *buffer) titleDraw(topLine byte, starting byte) {
	defer b.stringBuffer((2-topLine)+starting, 8, 'h', titleLines[6-topLine])
	if topLine > 0 {
		b.titleDraw(topLine-1, starting)
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

	var z, x byte
	x = 0
	for z = 0; z < 7; z++ {
		mainbuffer.clearBuffer(bgc)
		mainbuffer.outline(1, 1, xSize, ySize, 1, w, " ")
		mainbuffer.titleDraw(z, x)
		x++
		dl(500)
		mainbuffer.clrDrw(1)
	}

	//	mainbuffer.outline(1, 1, xSize, ySize, 1, w, " ")
	//	mainbuffer.titleDraw(5)
	//	mainbuffer.clearScreen(2)
	//	mainbuffer.drawFrame()
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
