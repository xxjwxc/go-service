package main

import (
	"fmt"

	_ "./server"
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
