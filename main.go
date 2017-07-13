package main

import (
	"test_glog/glog"
	"flag"
	"time"
	"os"
	"path/filepath"
	"strings"
	"fmt"
	"os/exec"
)

func main() {

	flag.Parse()

	// 创建一个log目录，与exe文件同级别
	dir := getCurPath()
	dir = strings.Join([]string{dir, "/log"}, "")
	os.Mkdir(dir, os.ModeDir)

	// 设置日志大小，日志路径
	glog.MaxSize = 1024 * 1024
	glog.Cheat(dir)
	// 设置关闭程序的时候，flush
	defer glog.Flush()

	glog.Infoln("123")
	glog.Errorln("e123")
	glog.Warningln("w123")

	// 多线程写日志
	func1 := func(tag string){
		for i:=0; i< 100000; i++ {
			glog.Infoln(tag, "123")
			glog.Errorln(tag, "e123")
			glog.Warningln(tag, "w123")
		}
	}
	go func1("func1")
	go func1("func2")
	go func1("func3")
	go func1("func4")
	go func1("func5")

	// 等待协程函数完毕
	time.Sleep(10 * time.Second)
}

func getCurPath() string {
	fmt.Println(os.Args[0])
	file, _ := exec.LookPath(os.Args[0])

	if len(file) == 0 {
		file = os.Args[0]
	}

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)

	//将全路径用\\分割，得到4段，①E: ②golang ③test ④a.exe
	splitstring := strings.Split(path, "\\")

	//size为4
	size := len(splitstring)

	//将全路径用最后一段(④a.exe)进行分割，得到2段，①E:\\golang\\test\\ ②a.exe
	splitstring = strings.Split(path, splitstring[size-1])

	//将①(E:\\golang\\test\\)中的\\替换为/，最终得到结果E:/golang/test/
	rst := strings.Replace(splitstring[0], "\\", "/", size-1)
	return rst
}

