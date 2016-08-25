package main

import (
	"fmt"

	"./server"
)

func CallBack() {
	//
	fmt.Println("111")
	//-----------------------------end
}

func main() {
	server.OnStart(CallBack)
	//CallBack()
}
