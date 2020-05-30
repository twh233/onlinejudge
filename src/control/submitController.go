package control

import (
	"base"
	"fmt"
	"github.com/gorilla/schema"
	"io"
	"net/http"
	"sync"
)

var JudgeResult = make(map[string]base.Result)
var Mutex sync.RWMutex

type JudgeQueue chan base.Submit

const form = `
    <html>
    	<body>
        <form action="#" method="post" name="Submit">
        	<table>
				<tr><td>SubmitId：<input type="text" name="SubmitId"/></td></tr>
				<tr><td>ProblemId：<input type="text" name="ProblemId"/></td></tr>
				<tr><td>ProblemType：<input type="text" name="ProblemType"/></td></tr>
				<tr><td>Language：<input type="text" name="Language"/></td></tr>
				<tr><td>TimeLimit：<input type="text" name="TimeLimit"/></td></tr>
				<tr><td>MemoryLimit：<input type="text" name="MemoryLimit"/></td></tr>
				<tr><td>UserCode：<textarea id="UserCode" name="UserCode" style="width:500px;height:300px;" value=""></textarea></td></tr>
				<tr><td><input type="submit" value="submit"/></td></tr>
            </table>
        </form>
    	</body>
    </html>
`

const form2 = `
    <html>
    	<body>
        <form action="#" method="post" name="Submit">
        	<table>
            	需要查询的题目ID：<input type="text" name="SubmitId" />
           	 	<input type="submit" value="查询"/>
			</table> 	
        </form>
    	</body>
    </html>
`

func (jq JudgeQueue)SubmitHandler(w http.ResponseWriter, r *http.Request) {
	//body,_ := ioutil.ReadAll(r.Body)
	//var submit base.Submit
	//err := json.Unmarshal(body, &submit); if err != nil {
	//	fmt.Println("Get Error:",err)
	//}
	//result := base.Result{}
	//result.Result = base.Waiting
	//Mutex.Lock()
	//JudgeResult[submit.SubmitId] = result
	//Mutex.Unlock()
	//jq <- submit


	//w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		//body,_ := ioutil.ReadAll(r.Body)
		r.ParseForm()
		var submit base.Submit
		if err := schema.NewDecoder().Decode(&submit, r.PostForm); err != nil {
			fmt.Println("error")
		}
		//err := json.Unmarshal(body, &submit); if err != nil {
		//	fmt.Println("Get Error:",err)
		//}
		result := base.Result{}
		result.Result = base.Waiting
		Mutex.Lock()
		JudgeResult[submit.SubmitId] = result
		Mutex.Unlock()
		jq <- submit
	}
}

func GetResultHandler(w http.ResponseWriter, r *http.Request) {
	//body,_ := ioutil.ReadAll(r.Body)
	//var submit base.Submit
	//err := json.Unmarshal(body, &submit); if err != nil {
	//	fmt.Println("Get Error:",err)
	//}
	//Mutex.RLock()
	//result := JudgeResult[submit.SubmitId]
	//Mutex.RUnlock()
	//responseResult,_ := json.Marshal(result)
	//_, _ = w.Write(responseResult)

	switch r.Method {
	case "GET":
		io.WriteString(w, form2)
	case "POST":
		//body,_ := ioutil.ReadAll(r.Body)
		r.ParseForm()
		var submit base.Submit
		if err := schema.NewDecoder().Decode(&submit, r.PostForm); err != nil {
			fmt.Println("error")
		}
		//err := json.Unmarshal(body, &submit); if err != nil {
		//	fmt.Println("Get Error:",err)
		//}
		Mutex.RLock()
		result := JudgeResult[submit.SubmitId]
		Mutex.RUnlock()
		_, _ = w.Write([]byte(base.ReturnResult(result.Result)))
	}
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello~"))
}
