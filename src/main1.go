package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"syscall"
	"base"
)

var time_limit_s int
var language int
var memory_limit_MB int


var readFile string
var outFile string
var writeFile string


var running_time int64
var isSpj = false

func set_limit(resouce int, rlim_cur uint64, rlim_max uint64) int {
	var limit syscall.Rlimit
	//syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit)
	//fmt.Println(int(limit.Cur))
	//fmt.Println(int(limit.Max))

	limit.Cur = rlim_cur
	limit.Max = rlim_max
	syscall.Setrlimit(resouce, &limit)

	//syscall.Getrlimit(syscall.RLIMIT_NOFILE, &limit)
	//fmt.Println(int(limit.Cur))
	//fmt.Println(int(limit.Max))
	return 0
}

func main() {
	flag.IntVar(&time_limit_s, "time_limit_s", 3, "")
	flag.IntVar(&language, "language", 2, "")
	flag.IntVar(&memory_limit_MB, "memory_limit_MB", 100, "")
	flag.Parse()

	var path = "/gowork/src/data/judge/"
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return
	}
	for _, fi := range rd {
		if fi.Name() == "spj.cpp" {
			isSpj = true
			break
		}
	}
	for _, fi := range rd {
		length := len(fi.Name())
		if length <= 2 {
			continue
		}
		if fi.Name()[:length - 3] + ".in" != fi.Name() {
			continue
		}
		readFile = path + fi.Name()
		outFile = readFile[:len(readFile) - 3] + ".out"
		writeFile = path[:len(path) - 6] + "out/" + fi.Name()[:length - 3] + ".user"
		running_time = time.Now().UnixNano() / 1000000
		fmt.Println(running_time)

		set_limit(syscall.RLIMIT_CPU, uint64(time_limit_s), uint64(time_limit_s))
		set_limit(syscall.RLIMIT_CORE, uint64(0), uint64(0))

		//test
		//judgePE(readFile, outFile, writeFile)

		//if language == 1 { // c++
		//	cmd := exec.Command("g++", "Main.cpp", "-o", "Main")
		//	cmd.Output()
		//
		//	cmd = exec.Command("./Main")
		//	out,err := cmd.Output()
		//	if err != nil {
		//		fmt.Println("c++ Command fail:", err)
		//		return
		//	}
		//	fmt.Println(string(out))
		//} else if language == 2 { // java
		//	cmd := exec.Command("javac", "Main.java")
		//	cmd.Output()
		//
		//	cmd = exec.Command("java","Main", "tom")
		//	out,err := cmd.Output()
		//	if err != nil {
		//		fmt.Println("java Command fail:", err)
		//		return
		//	}
		//	fmt.Println(string(out))
		//} else { // py
		//	cmd := exec.Command("python", "Main.py")
		//	out,err := cmd.Output()
		//	if err != nil {
		//		fmt.Println("py Command fail:", err)
		//		return
		//	}
		//	fmt.Println(string(out))
		//}
		fmt.Println(base.Normal)
		go func(){
			//running_time = time.Now().UnixNano() / 1000000
			//set_limit(1, 1)
		}()
	}

}