package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

//
type Cfg_Base struct {
	Serial_number       string `json:"serial_number,omitempty"`       //对应JSON的serial_number,如果为空置则忽略字段
	Service_name        string `json:"service_name,omitempty"`        //对应JSON的 service_name,如果为空置则忽略字段
	Service_displayname string `json:"service_displayname,omitempty"` //对应JSON的 service_displayname,如果为空置则忽略字段
	Sercice_desc        string `json:"sercice_desc,omitempty"`        //对应JSON的 sercice_desc,如果为空置则忽略字段
	IsDev               bool   `json:"is_dev,omitempty"`              //是否是开发版本
}

var _map = Config{}

func init() {
	onInit()
	flag.Parse()
}

func onInit() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)

	// if len(path) > 0 {
	// 	path += "/"
	// }
	path = filepath.Dir(path)
	err := InitFile(path + "/config.json")
	if err != nil {
		fmt.Println("InitFile: ", err.Error())
		return
	}
}

func InitFile(filename string) error {
	if IsRunTesting() {
		bytes := []byte(test_file)
		if err := json.Unmarshal(bytes, &_map); err != nil {
			fmt.Println("Unmarshal: ", err.Error())
			return err
		}

		return nil
	} else {
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("ReadFile: ", err.Error())
			return err
		}

		if err := json.Unmarshal(bytes, &_map); err != nil {
			fmt.Println("Unmarshal: ", err.Error())
			return err
		}

		return nil
	}

}

//获取service配置信息
func GetServiceConfig() (name, displayName, desc string) {
	name = _map.Service_name
	displayName = _map.Service_displayname
	desc = _map.Sercice_desc
	return
}

//是否是开发版本
func OnIsDev() bool {
	return _map.IsDev
}
