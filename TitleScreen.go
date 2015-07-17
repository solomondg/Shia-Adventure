package main

import (
	"fmt"
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

/*
 _______  __   __  ___   _______  _______  __     _______  ______   __   __  _______  __    _  _______  __   __  ______    _______
|       ||  | |  ||   | |   _   ||       ||  |   |   _   ||      | |  | |  ||       ||  |  | ||       ||  | |  ||    _ |  |       |
|  _____||  |_|  ||   | |  |_|  ||  _____||__|   |  |_|  ||  _    ||  |_|  ||    ___||   |_| ||_     _||  | |  ||   | ||  |    ___|
| |_____ |       ||   | |       || |_____        |       || | |   ||       ||   |___ |       |  |   |  |  |_|  ||   |_||_ |   |___
|_____  ||       ||   | |       ||_____  |       |       || |_|   ||       ||    ___||  _    |  |   |  |       ||    __  ||    ___|
 _____| ||   _   ||   | |   _   | _____| |       |   _   ||       | |     | |   |___ | | |   |  |   |  |       ||   |  | ||   |___
|_______||__| |__||___| |__| |__||_______|       |__| |__||______|   |___|  |_______||_|  |__|  |___|  |_______||___|  |_||_______|

*/

func makebuffer() {
	buffer := make([][]string, 128)
	for i := 0; i < 128; i++ {
		buffer[i] = make([]string, 32)
	}
}

var ()

func main() {
	fmt.Println("\033[45m")
	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println("\033[49m")
}
