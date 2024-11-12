package main

import "fmt"

func clearScreen() {
	fmt.Print("\033c")
	fmt.Print("\033[?25l")

}

func renderDrawLoop(r *render ) {

	if buildScreen(r) == 0 {
		for {
			renderScreen(r)

		}

	}
}

func renderProcessLoop(r *render) {

	for {
		processScreen(r)
	}
}
