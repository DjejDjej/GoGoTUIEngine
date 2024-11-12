package main

import "time"

//	"github.com/eiannone/keyboard"




func main() {
	
r := newRender(25, 25)
go renderDrawLoop(r)
go renderProcessLoop(r)
b:= newBlock(Red,"x")
	
draw(r,b,0,2)
time.Sleep(10 * time.Second)


}



