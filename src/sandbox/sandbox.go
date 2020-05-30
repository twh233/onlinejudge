package sandbox

import (
	"base"
	"fmt"
	"judge"
	"os"
	"syscall"
)

type Sandbox interface {
	Run(submit base.Submit) base.Result
}

type StdSandbox struct {
	TimeLimit int64
	MemoryLimit int64

	TimeUsed int64
	MemoryUsed int64
}

func (s *StdSandbox)Run(submit base.Submit) (result base.Result, err error) {
	judgeDir := base.RootDir + base.JudgeDir + "/" + submit.SubmitId
	dataDir := base.RootDir + base.DataDir + "/" + submit.ProblemId

	_, _ = os.OpenFile(dataDir+"/000000.in", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0)
	_, _ = os.OpenFile(dataDir+"/000000.out", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0)

	_ = os.RemoveAll(judgeDir)
	exist, err := base.PathExists(judgeDir); if err != nil {
		fmt.Println("get dir error!", err)
		return
	}
	if exist {
		fmt.Println("judgeDir is exist, submitId is not correct!")
		_ = os.RemoveAll(judgeDir)
	}
	oldMask := syscall.Umask(0)
	err = os.Mkdir(judgeDir, os.ModePerm); if err != nil {
		fmt.Println("mkdir fail!")
		return
	}
	syscall.Umask(oldMask)
	judge2 := judge.CreateJudger(submit); if judge2 == nil {
		result.Result = base.SystemError
		return
	}
	result = judge2.Run(s.TimeLimit, s.MemoryLimit, submit.ProblemType)
	_ = os.RemoveAll(judgeDir)
	return
}

