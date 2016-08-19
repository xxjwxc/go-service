package loginfo

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/golang/glog"
)

const ( //
	Log_Error   = iota //打印 Error 及以上级别
	Log_warning        //打印 warning 及以上级别
	Log_Info           //默认的返回值，为0，自增 //打印 Info 及以上级别
)

var Confg_map = map[string]string{}

func init() {
	onInit()
	flag.Parse()
}

func onInit() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	path = filepath.Dir(path)

	_map, err := readFile(path + "/config.json")
	if err != nil {
		fmt.Println("readFile: ", err.Error())
		return
	}

	flag.Set("v", _map["log_v"])         //初始化命令行参数
	flag.Set("log_dir", _map["log_dir"]) //初始化命令行参数 默认目录
	n, _ := strconv.ParseUint(string(_map["log_maxSize"]), 11, 64)
	glog.MaxSize = 1024 * 1024 * n
}

//获取端口号
func GetServerPort() (strPort string) {
	strPort = string(Confg_map["port"])
	return
}

//获取端口号
func GetServerHttpsPort() (strPort string) {
	strPort = string(Confg_map["https_port"])
	return
}

//获取mysql连接字符串
func GetMysqlUrl() (strMysqlUrl string) {
	strMysqlUrl = string(Confg_map["mysql_url"])
	return
}

func readFile(filename string) (map[string]string, error) {
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

func SetFlag(f int) {
	str2 := fmt.Sprintf("%d", f)
	flag.Set("v", str2) //初始化命令行参数
}

func SetOutPutDir(outDir string) {
	flag.Set("log_dir", outDir) //初始化命令行参数
}

//func Info(args ...interface{}) {
//	if LogFlag <= Log_Info {
//		glog.Info(args)
//	}
//}

//func Warning(args ...interface{}) {
//	if LogFlag <= Log_warning {
//		glog.Warning(args)
//	}
//}

//func Error(args ...interface{}) {
//	if LogFlag <= Log_Error {
//		glog.Error(args)
//	}
//}

//func Infof(format string, args ...interface{}) {
//	if LogFlag <= Log_Info {
//		glog.Infof(format, args)
//	}
//}

//func Warningf(format string, args ...interface{}) {
//	if LogFlag <= Log_warning {
//		glog.Warningf(format, args)
//	}
//}

//func Errorf(format string, args ...interface{}) {
//	if LogFlag <= Log_Error {
//		glog.Errorf(format, args)
//	}
//}

//刷新
func Flush() {
	glog.Flush()
}
