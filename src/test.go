package main

import (
	"fmt"
	"os"

	"github.com/xie1xiao1jun/go-service/src/data/config"
	"github.com/xie1xiao1jun/go-service/src/server"
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
		//name, displayName, desc := config.GetServiceConfig()
		server.OnInit(config.GetServiceConfig())
		server.OnStart(CallBack)
	}
}
