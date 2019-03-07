package main

import (
	"data/config"
	"fmt"
	"os"

	"./server"
)

func CallBack() {
	//
	fmt.Println("111")
	//-----------------------------end
}

func main() {
	if config.OnIsDev() || len(os.Args) == 0 {
		CallBack()
	} else {
		server.OnStart(CallBack)
	}
}
