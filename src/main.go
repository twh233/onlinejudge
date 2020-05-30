package main

import (
	"base"
	"control"
	"net/http"
	"sandbox"
	"github.com/gorilla/mux"
)



var jq control.JudgeQueue = make(chan base.Submit, 100)

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", control.SayHello)
	r.HandleFunc("/submit", jq.SubmitHandler)
	r.HandleFunc("/result", control.GetResultHandler)
	//r.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir(base.PUBLIC_PATH))))

	go judge()

	http.ListenAndServe(":8888", r)
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
