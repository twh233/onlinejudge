package main

import (
	"base"
	"control"
	"net/http"
	"sandbox"
)

var jq control.JudgeQueue = make(chan base.Submit, 100)

func main() {
	http.HandleFunc("/", control.SayHello)
	http.HandleFunc("/submit", jq.SubmitHandler)
	http.HandleFunc("/result", control.GetResultHandler)
	go judge()
	http.ListenAndServe(":8888", nil)
}

func judge(){
	for{
		submit := <-jq
		result := base.Result{}
		result.Result = base.Judging
		control.Mutex.Lock()
		control.JudgeResult[submit.SubmitId] = result
		control.Mutex.Unlock()
		sandbox2 := sandbox.StdSandbox{}
		sandbox2.TimeLimit = submit.TimeLimit
		sandbox2.MemoryLimit = submit.MemoryLimit

		result, _ = sandbox2.Run(submit)

		control.Mutex.Lock()
		control.JudgeResult[submit.SubmitId] = result
		control.Mutex.Unlock()
	}
}
