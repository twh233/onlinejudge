package sandbox_test

import (
	"base"
	"fmt"
	"os"
	"path/filepath"
	"sandbox"
	"strings"
	"testing"
)

func TestAplusBAC(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A1/ac")
	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A1"
	submit.UserCode = code
	submit.SubmitId = "0"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:", result)
	if result.Result != base.Accepted {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestAplusBWA(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A1/wa")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A1"
	submit.UserCode = code
	submit.SubmitId = "1"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:", result)
	if result.Result != base.WrongAnswer {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestAplusBRE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A1/re")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A1"
	submit.UserCode = code
	submit.SubmitId = "2"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:", result)
	if result.Result != base.RuntimeError {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestAplusBDG(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A1/dg")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A1"
	submit.UserCode = code
	submit.SubmitId = "3"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:", result)
	if result.Result != base.Danger {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestAplusBPE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A1/pe")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A1"
	submit.UserCode = code
	submit.SubmitId = "4"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:", result)
	if result.Result != base.PresentationError {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestAplusBOLE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A1/ole")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A1"
	submit.UserCode = code
	submit.SubmitId = "25"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:", result)
	if result.Result != base.OutputLimitExceeded {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestAplusBTLE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A1/tle")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A1"
	submit.UserCode = code
	submit.SubmitId = "26"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:", result)
	if result.Result != base.TimeLimitExceeded {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestAplusBMLE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A1/mle")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A1"
	submit.UserCode = code
	submit.SubmitId = "27"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:", result)
	if result.Result != base.MemoryLimitExceeded {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestASubBAC(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A2/ac")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A2"
	submit.UserCode = code
	submit.SubmitId = "5"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	if result.Result != base.Accepted {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestASubBOLE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A2/ole")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A2"
	submit.UserCode = code
	submit.SubmitId = "6"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	if result.Result != base.OutputLimitExceeded {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestKMPAC(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A3/ac")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A3"
	submit.UserCode = code
	submit.SubmitId = "7"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	if result.Result != base.Accepted {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result)
	}
	//time.Sleep(time.Second * 3)
}

func TestKMPOLE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A3/ole")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A3"
	submit.UserCode = code
	submit.SubmitId = "8"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	if result.Result != base.OutputLimitExceeded {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

/*func TestKMPTLE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A3/tle")

	submit := control.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A3"
	submit.UserCode = code
	submit.SubmitId = "9"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	if result.Result != base.TimeLimitExceeded {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}*/

func TestGreedAC(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A4/ac")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A4"
	submit.UserCode = code
	submit.SubmitId = "10"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	if result.Result != base.Accepted {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestGreedCE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A4/ce")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A4"
	submit.UserCode = code
	submit.SubmitId = "11"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	if result.Result != base.CompilationError {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestGreedWA(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A4/wa")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A4"
	submit.UserCode = code
	submit.SubmitId = "12"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	if result.Result != base.WrongAnswer {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestGreedTLE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A4/tle")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A4"
	submit.UserCode = code
	submit.SubmitId = "13"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	if result.Result != base.TimeLimitExceeded {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestGreedMLE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A4/mle")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A4"
	submit.UserCode = code
	submit.SubmitId = "14"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	if result.Result != base.MemoryLimitExceeded {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestGreedRE(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/A4/re")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.ICPC
	submit.ProblemId = "A4"
	submit.UserCode = code
	submit.SubmitId = "15"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	if result.Result != base.RuntimeError {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestSPJ40(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/SPJ/Score/score40")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.SPJ
	submit.ProblemId = "Score"
	submit.UserCode = code
	submit.SubmitId = "16"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:",result)
	if result.Result != base.Score + 40 {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestSPJ100(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/SPJ/Score/score100")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.SPJ
	submit.ProblemId = "Score"
	submit.UserCode = code
	submit.SubmitId = "17"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:",result)
	if result.Result != base.Score + 100 {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestSPJAC(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/SPJ/Normal/ac")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.SPJ
	submit.ProblemId = "Normal"
	submit.UserCode = code
	submit.SubmitId = "18"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:",result)
	if result.Result != base.Accepted {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestSPJWA(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/SPJ/Normal/wa")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.SPJ
	submit.ProblemId = "Normal"
	submit.UserCode = code
	submit.SubmitId = "19"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 2000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:",result)
	if result.Result != base.WrongAnswer {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestFuncAC(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/Func/ac")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.Functional
	submit.ProblemId = "Func"
	submit.UserCode = code
	submit.SubmitId = "20"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:",result)
	if result.Result != base.Accepted {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestFuncWA(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/Func/wa")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.Functional
	submit.ProblemId = "Func"
	submit.UserCode = code
	submit.SubmitId = "21"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println(result)
	if result.Result != base.WrongAnswer {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestRefreeAC(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/Refree/ac")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.Referee
	submit.ProblemId = "Refree"
	submit.UserCode = code
	submit.SubmitId = "22"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println("result:",result)
	if result.Result != base.Score + 100 {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestRefreeScore100(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/Refree/score100")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.Referee
	submit.ProblemId = "Refree"
	submit.UserCode = code
	submit.SubmitId = "23"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println(result)
	if result.Result < base.Score {
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func TestRefreeScore0(t *testing.T) {
	currentDir := getCurrentDir()
	code := base.ReadCodeFromFile(currentDir + "/Refree/score0")

	submit := base.Submit{}
	submit.Language = base.CPP
	submit.ProblemType = base.Referee
	submit.ProblemId = "Refree"
	submit.UserCode = code
	submit.SubmitId = "24"

	sanbox := sandbox.StdSandbox{}
	sanbox.TimeLimit = 1000
	sanbox.MemoryLimit = 128*1024*1024
	result,_ := sanbox.Run(submit)
	fmt.Println(result)
	if result.Result != base.Score{
		t.Errorf("A plus B is not correct! judger error!")
		fmt.Println("result is:",result.Result)
	}
	//time.Sleep(time.Second * 3)
}

func getCurrentDir() (cd string){
	cd, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	cd = strings.Replace(cd, "/sandbox", "/code", -1)
	return
}

func setFile(cd string, pid string, fileName string){
	dataDir := base.RootDir + base.DataDir + "/" + pid
	//base.CreateDir(dataDir)
	inputContext := base.ReadAll(cd + "/A1/" + fileName)
	file,_ := os.OpenFile(dataDir + "/" + fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	_, _ = file.Write(inputContext)
}