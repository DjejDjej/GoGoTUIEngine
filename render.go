package main

import (
	"fmt"
	"strings"
)

type block struct {
	col, char, data string
}

type render struct {
	width, height, maxlen, changes int
	buffer                         []block
	querry                         []action
	querrytodo                     []action
	screen                         string
	screenBuffer                   string
}

type action struct {
	x, y int
	b    block
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
	r.changes = 0
	r.screenBuffer = ""
	r.buffer = make([]block, 0)
	r.querry = make([]action, 0)
	r.querrytodo = make([]action, 0)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r.buffer = append(r.buffer, b)

		}
	}
	return r

}

func newAction(x, y int, b block) *action {
	a := &action{
		x: x,
		y: y,
		b: b,
	}
	return a

}

func buildScreen(r *render) int {

	var builder strings.Builder
	r.screenBuffer = ""
	for i := 0; i <= len(r.buffer)-1; i++ {
		builder.WriteString(r.buffer[i].data + Invisible)
		if (i+1)%r.height == 0 && (i+1) != len(r.buffer) {
			builder.WriteString("\n")

		}
	}
	r.screenBuffer = builder.String()
	return 0
}

func processScreen(r *render) {

	if len(r.querry) == 0 {
		return
	}
	if len(r.screenBuffer) == 0 {

		return
	}

	last := r.querry[0]

	r.buffer[(last.y*r.height)+last.x] = last.b
	split := strings.Split(r.screenBuffer, Invisible)
	split[(last.y*r.height)+last.x] = last.b.data

	r.screenBuffer = strings.Join(split, Invisible)
	r.querry = r.querry[1:]
	r.querrytodo = append(r.querrytodo, last)
}

func renderScreen(r *render) {

	if len(r.screen) == 0  {
	renderFullScreen(r)
	}
	if len(r.querrytodo) == 0 {
		return
	}

	if len(r.querrytodo) < 20 {
		 renderQuerry(r)
		return
	}

	renderFullScreen(r)
	r.querrytodo = r.querrytodo[:0]
}

func renderFullScreen(r *render){
	clearScreen()
	r.screen = r.screenBuffer
	fmt.Print(r.screen)

}

func renderQuerry(r *render) {

	for i := 0; i < len(r.querrytodo); i++ {

		c := r.querrytodo[i]
		fmt.Printf("\033[%d;%dH%s", c.y, c.x, c.b.data)

	}

}

// func changeBlock(r *render, b block, x, y int) {
// 	r.buffer[(y*r.height)+x] = b
// 	split := strings.Split(r.screen, Invisible)
// 	split[0] = b.data
//
// }

func draw(r *render, b block, x, y int) {
	a := newAction(x, y, b)
	r.querry = append(r.querry, *a)

}

func drawLine(r *render, b block, x, y int) {
	a := newAction(x, y, b)
	r.querry = append(r.querry, *a)

}
