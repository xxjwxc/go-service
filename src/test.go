package main

import (
	"fmt"
	"os"
	"time"

	"github.com/xxjwxc/go-service/src/data/config"
	"github.com/xxjwxc/go-service/src/server"
)

func CallBack() {
	//
	fmt.Println("111")
	for {
		ticker := time.NewTicker(1 * time.Second)
		<-ticker.C
		fmt.Println("yyyyy")
	}
	//-----------------------------end
}

func main() {
	if config.OnIsDev() || len(os.Args) == 0 {
		CallBack()
	} else {
		//name, displayName, desc := config.GetServiceConfig()
		server.On(config.GetServiceConfig()).Start(CallBack)
	}
}
