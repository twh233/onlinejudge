package main

import (
	"fmt"
	"syscall"
)

const (
	AC    = 0
	TLE   = 1
	MLE   = 2
	RE    = 3
	OLE   = 4
	WA    = 5
	PE    = 6
	ERROR = 7
	SC    = 8
)
var time_use = 0
var last_ret = AC

func check(pid int) int {
	var wstat syscall.WaitStatus
	var use syscall.Rusage
	_, err := syscall.Wait4(pid, &wstat, 0, &use)
	if err != nil {
		fmt.Printf("error waiting 4 %d: %v\n", pid, err)
		panic("wait4")
	}
	return 0
}

func return_ans(ret int, useTime int, useMemory int, running_time int64) int {

	time_use += useTime
	last_ret = ret
	return 0
}