package main

import (
	"fmt"

	"log"
	"github.com/eiannone/keyboard"
)

func main() {
	x:=0
	y:=0
	
	r := newRender(100, 100)
	renderScreenf(r)

	def := newBlock(Cyan, "█")
	play := newBlock(Red, "█")
    defer keyboard.Close()

    for {
         if err := keyboard.Open(); err != nil {
        log.Fatal(err)
    }

	    char, _, err := keyboard.GetKey()
        
	if err != nil {
            log.Fatal(err)
        }

	if char== 'x'{
	break	
	}
	if char== 's'{
	fmt.Print("XD")
	changeChar(r,def,x,y)
	y--
	changeChar(r,play,x,y)
	}
	if char== 'w'{
	changeChar(r,def,x,y)
	y++
	changeChar(r,play,x,y)
	}
	if char== 'd'{
	changeChar(r,def,x,y)
	x++
	changeChar(r,play,x,y)
	}
	if char== 'a'{
	changeChar(r,def,x,y)
	x--
	changeChar(r,play,x,y)
	}
    
	renderScreenf(r)
}


}
