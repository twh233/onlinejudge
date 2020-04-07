package control

import (
	"base"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var JudgeResult = make(map[string]base.Result)
var Mutex sync.RWMutex

type JudgeQueue chan base.Submit

func (jq JudgeQueue)SubmitHandler(w http.ResponseWriter, r *http.Request) {
	body,_ := ioutil.ReadAll(r.Body)
	var submit base.Submit
	err := json.Unmarshal(body, &submit); if err != nil {
		fmt.Println("Get Error:",err)
	}
	result := base.Result{}
	result.Result = base.Waiting
	Mutex.Lock()
	JudgeResult[submit.SubmitId] = result
	Mutex.Unlock()
	jq <- submit
}

func GetResultHandler(w http.ResponseWriter, r *http.Request) {
	body,_ := ioutil.ReadAll(r.Body)
	var submit base.Submit
	err := json.Unmarshal(body, &submit); if err != nil {
		fmt.Println("Get Error:",err)
	}
	Mutex.RLock()
	result := JudgeResult[submit.SubmitId]
	Mutex.RUnlock()
	responseResult,_ := json.Marshal(result)
	_, _ = w.Write(responseResult)
}