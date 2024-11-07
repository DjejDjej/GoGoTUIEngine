package main

import (
	"fmt"
	"strings"
)

type block struct {
	col, char, data string
}

type render struct {
	width, height, maxlen int
	buffer                []block
	screen                string
}

func clearScreen() {
	fmt.Print("\033c")
    	fmt.Print("\033[?25l")

}

func newBlock(col, char string) block {
	b := block{
		col:  col,
		char: char,
		data: "",
	}

	b.data = col + char + Reset

	return b
}

func newRender(width, height int) *render {
	defer fmt.Print("\033[?25h")
	r := &render{
		width:  width,
		height: height,
	}
	b := newBlock(Cyan, "█")
	r.maxlen = r.width * r.height
	r.screen = ""
	r.buffer = make([]block, 0)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r.buffer = append(r.buffer, b)

		}
	}
	return r

}

func renderScreenf(r *render) {
	r.screen = ""
	for i := 0; i <= len(r.buffer)-1; i++ {
		if (i+1) % r.height == 0 && (i + 1) != len(r.buffer) {
			r.screen += "\n"
		}

		r.screen += r.buffer[i].data + Invisible
	}
	clearScreen()

			
	fmt.Print(r.screen)
}





func changeChar(r *render, b block, x, y int) {

	r.buffer[(y*r.height) + x] = b
	fmt.Printf("\033[%d;%dH%s", y, x, b.data)
}


func changeBlock(r *render, b block, x, y int) {
r.buffer[(y*r.height) + x] = b

split := strings.Split(r.screen,Invisible)
split[0] = b.data
r.screen = strings.Join(split,Invisible)

}
