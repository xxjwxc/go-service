package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"../loginfo"
	"github.com/golang/glog"
	"github.com/kardianos/service"
)

func readFile(filename string) (map[string]string, error) {
	var Confg_map = map[string]string{}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}

	if err := json.Unmarshal(bytes, &Confg_map); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}

	return Confg_map, nil
}

func GetServiceInit() (name, displayName, desc string) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	path = filepath.Dir(path)

	_map, err := readFile(path + "/config.json")
	if err != nil {
		fmt.Println("readFileError: ", err.Error())
		return
	}
	name = _map["service_name"]
	displayName = _map["service_displayname"]
	desc = _map["sercice_desc"]
	return
}

func OnStart(callBack func()) {
	name, displayName, desc := GetServiceInit()
	//	var name = "GoServiceTest"
	//	var displayName = "Go Service Test"
	//	var desc = "This is a test Go service.  It is designed to run well."
	var s, err = service.NewService(name, displayName, desc)
	if err != nil {
		fmt.Printf("%s unable to start: %s", displayName, err)
		return
	}

	fmt.Printf("Service \"%s\" do.\n", displayName)

	if len(os.Args) > 1 {
		var err error
		verb := os.Args[1]
		switch verb {
		case "install":
			{
				err = s.Install()
				if err != nil {
					fmt.Printf("Failed to install: %s\n", err)
					return
				}
				fmt.Printf("Service \"%s\" installed.\n", displayName)
			}
		case "remove":
			{
				err = s.Remove()
				if err != nil {
					fmt.Printf("Failed to remove: %s\n", err)
					return
				}
				fmt.Printf("Service \"%s\" removed.\n", displayName)
			}
		case "run":
			{
				doWork(callBack)
			}
		case "start":
			{
				err = s.Start()
				if err != nil {
					fmt.Printf("Failed to start: %s\n", err)
					return
				}
				fmt.Printf("Service \"%s\" started.\n", displayName)
				glog.V(loginfo.Log_Info).Infof("Service \"%s\" started.\n", displayName)
			}
		case "stop":
			{
				err = s.Stop()
				if err != nil {
					fmt.Printf("Failed to stop: %s\n", err)
					glog.V(loginfo.Log_Error).Infof("Failed to stop: %s\n", err)
					glog.Flush()
					return
				}
				fmt.Printf("Service \"%s\" stopped.\n", displayName)
				glog.V(loginfo.Log_Info).Infof("Service \"%s\" stopped.\n", displayName)
			}
		}
		return
	} else {
		fmt.Print("Failed to read args\n")
	}

	err = s.Run(func() error { // start
		go doWork(callBack)
		return nil
	}, func() error { // stop
		stopWork()
		return nil
	})

	if err != nil {
		s.Error(err.Error())
	}

	glog.Flush()
}

var exit = make(chan struct{})

func doWork(callBack func()) {
	glog.V(loginfo.Log_Info).Info("start Running!")
	go callBack() //go svc()
	for {
		select {
		case <-exit:
			os.Exit(0)
			return
		}
	}
}
func stopWork() {
	glog.V(loginfo.Log_Info).Info("I'm Stopping!")
	exit <- struct{}{}
}

//func Svc() {
//	m := martini.Classic()
//	m.Get("/", func() string {
//		loginfo.Info("tttttttt!")
//		loginfo.Flush()
//		return "Hello world!"
//	})
//	http.ListenAndServe(":8080", m)
//}
