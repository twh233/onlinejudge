package judge

import (
	"base"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"ptrace"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type RunResult struct {
	runResult int64
	timeUsed int64
	memoryUsed int64

	fileName map[string] int64
}

type ResourcesLimit struct {
	timeLimit int64
	memoryLimit int64
	problemType int64
}

type Runner interface {
	Run(limit ResourcesLimit) RunResult
}

type DoNothing struct {

}
func (dn *DoNothing)Run(limit ResourcesLimit) (rr RunResult) {
	rr.runResult = base.Accepted
	return
}

func (c *C)Run(limit ResourcesLimit) RunResult {
	return RunAdapter(base.C, c.codeDir, c.dataDir, "Main", ".in", ".user", limit)
}

func (cpp *CPP)Run(limit ResourcesLimit) RunResult {
	return RunAdapter(base.CPP, cpp.codeDir, cpp.dataDir, "Main", ".in", ".user", limit)
}

func (java *JAVA)Run(limit ResourcesLimit) RunResult {
	return RunAdapter(base.Java, java.codeDir, java.dataDir, "Main", ".in", ".user", limit)
}

func RunAdapter(language string, codeDir string, inputDataDir string, exeName string, oldSuffix string, newSuffix string, limit ResourcesLimit) (rr RunResult) {
	files, err := ioutil.ReadDir(inputDataDir)

	if err != nil {
		fmt.Println("Run ReadDir fail:", err)
		rr.runResult = base.SystemError
		return
	}
	var timeTotal int64 = 0
	var memoryMax int64 = 0
	inputFileCount := 0
	rr.fileName = make(map[string]int64)

	for _, file := range files {
		if strings.Contains(file.Name(), oldSuffix) {
			temp := make(chan RunResult)

			go func() {
				tp := RunByOneFile(language, codeDir, inputDataDir, codeDir, exeName, oldSuffix, newSuffix, file, limit)
				temp <- tp
			}()
			tempRr := <-temp
			//if strings.Compare(file.Name(), "000000.in") != 0 {
				inputFileCount++
				rr.fileName[file.Name()] = tempRr.runResult
				if tempRr.runResult != base.Accepted && limit.problemType != base.SPJ {
					rr.runResult = tempRr.runResult
					rr.timeUsed = tempRr.timeUsed
					rr.memoryUsed = tempRr.memoryUsed
					return
				}
				timeTotal = timeTotal + tempRr.timeUsed
				memoryMax = base.GetMaxInt64(memoryMax, tempRr.memoryUsed)
			//}
		}
	}

	if inputFileCount == 0 {
		return RunWithoutInputFile(language, codeDir, inputDataDir, codeDir, exeName, newSuffix, limit)
	}
	rr.memoryUsed = memoryMax
	rr.timeUsed = timeTotal

	if timeTotal > limit.timeLimit {
		rr.runResult = base.TimeLimitExceeded
		return
	}

	rr.runResult = base.Accepted
	return
}

func RunByOneFile(language string, codeDir string, inputDataDir string, outputDataDir string, exeName string, oldSuffix string, newSuffix string, f os.FileInfo, limit ResourcesLimit) (rr RunResult) {
	outputFileName := strings.Replace(f.Name(), oldSuffix, ".out", -1)

	inputFile, err := os.Open(inputDataDir + "/" + f.Name())

	defer inputFile.Close()
	if err != nil {
		fmt.Println("Run RunByOneFile inputFile fail:", err)
		rr.runResult = base.SystemError
		return
	}
	//create output file
	fileName := strings.Replace(f.Name(), oldSuffix, newSuffix, -1)

	outputFile, err := os.OpenFile(outputDataDir + "/" + fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	defer outputFile.Close()
	if err != nil {
		fmt.Println("Run RunByOneFile outputFile fail:", err)
		rr.runResult = base.SystemError
		return
	}
	var cmd *exec.Cmd
	switch language {
	case base.C, base.CPP:
		cmd = exec.Command(codeDir + "/" + exeName)
	case base.Java:
		cmd = exec.Command("java", exeName)
	case base.Python:
		cmd = exec.Command("python", exeName)
	default:
		fmt.Println("无效语言")
		rr.runResult = base.SystemError
		return
	}
	cmd.Stdout = outputFile
	cmd.Stdin = inputFile
	cmd.Dir = codeDir
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}
	err = cmd.Start(); if err != nil {
		fmt.Println("star err",err)
	}

	var timeUsage int64
	var memUsage int64

	cmd.Wait()

	var regs syscall.PtraceRegs
	pid := cmd.Process.Pid
	exit := true

	standardOutputFileSize := base.GetFileSize(inputDataDir + "/" + outputFileName)
	//https://www.jb51.cc/go/523369.html
	for {
		// 记得 PTRACE_SYSCALL 会在进入和退出syscall时使 tracee 暂停，所以这里用一个变量控制，RAX的内容只打印一遍
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			isAllowSysCall := ptrace.IsAllowSysCall(regs.Orig_rax) //SystemId
			if isAllowSysCall != base.Normal && language != base.Java {
				rr.runResult = isAllowSysCall
				return
			}
			if err != nil {
				break
			}
		}
		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			break
		}
		_, err = syscall.Wait4(pid, nil, 0, nil)

		ok, rss, runningTime, cpuTime := base.GetResourceUsage(cmd.Process.Pid)

		if !ok {
			break
		}
		timeUsage = base.GetMaxInt64(timeUsage, cpuTime); if timeUsage > limit.timeLimit || runningTime > 3 * limit.timeLimit {
			cmd.Process.Kill()
			cmd2 := exec.Command("kill", "-9", strconv.Itoa(pid))
			_ = cmd2.Run()
			rr.runResult = base.TimeLimitExceeded
			rr.timeUsed = timeUsage
			rr.memoryUsed = memUsage
			return
		}

		memUsage = base.GetMaxInt64(memUsage, rss); if memUsage * 3 > limit.memoryLimit * 2 {
			cmd.Process.Kill()
			cmd2 := exec.Command("kill", "-9", strconv.Itoa(pid))
			_ = cmd2.Run()
			rr.runResult = base.MemoryLimitExceeded
			rr.timeUsed = timeUsage
			rr.memoryUsed = memUsage
			return
		}
		outputFileSize := base.GetFileSize(outputDataDir + "/" + fileName)
		if outputFileSize > 3 * standardOutputFileSize {
			cmd.Process.Kill()
			cmd2 := exec.Command("kill", "-9", strconv.Itoa(pid))
			_ = cmd2.Run()
			rr.runResult = base.OutputLimitExceeded
			rr.timeUsed = timeUsage
			rr.memoryUsed = memUsage
			return
		}

		exit = !exit
	}

	cmd2 := exec.Command("kill", "-9", strconv.Itoa(pid))
	_ = cmd2.Run()

	rr.timeUsed = timeUsage
	rr.memoryUsed = memUsage

	if strings.Compare(f.Name(), "000000.in") != 0 {
		rr.runResult = CompareOneFile(outputDataDir + "/" + fileName, inputDataDir + "/" + outputFileName)
	} else {
		rr.runResult = base.Accepted
	}

	fmt.Println(rr)
	return
}

func RunWithoutInputFile(lanuage string, codeDir string, inputDataDir, outputDataDir string, exeName string, suffix string, limit ResourcesLimit) (rr RunResult) {
	outputFile, err := os.OpenFile(outputDataDir + "/1" + suffix, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0)
	defer outputFile.Close()
	if err != nil {
		rr.runResult = base.SystemError
		return
	}

	var cmd *exec.Cmd
	switch lanuage{
	case base.C, base.CPP:
		cmd = exec.Command(codeDir + "/" + exeName)
	case base.Java:
		cmd = exec.Command("java", exeName)
	case base.Python:
		cmd = exec.Command("python", exeName)
	default:
		rr.runResult = base.SystemError
		return
	}

	cmd.Stdout = outputFile
	cmd.Dir = codeDir
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}
	err = cmd.Start(); if err != nil {
		fmt.Println("star err",err)
	}

	var timeUsage int64
	var memUsage int64

	_ = cmd.Wait()

	var regs syscall.PtraceRegs
	pid := cmd.Process.Pid
	exit := true

	standardOutputFileSize := base.GetFileSize(inputDataDir + "/1.out")

	go func(){
		for{
			ok, rss, runningTime, cpuTime := base.GetResourceUsage(cmd.Process.Pid)
			if !ok {
				break
			}
			timeUsage = base.GetMaxInt64(timeUsage, cpuTime); if timeUsage > limit.timeLimit || runningTime > 3 * limit.timeLimit {
				cmd.Process.Kill()
				cmd2 := exec.Command("kill", "-9", strconv.Itoa(pid))
				_ = cmd2.Run()
				rr.runResult = base.TimeLimitExceeded
				rr.timeUsed = timeUsage
				rr.memoryUsed = memUsage
				return
			}
			memUsage = base.GetMaxInt64(memUsage, rss); if memUsage*3 > limit.memoryLimit*2 {
				cmd.Process.Kill()
				cmd2 := exec.Command("kill", "-9", strconv.Itoa(pid))
				_ = cmd2.Run()
				rr.runResult = base.MemoryLimitExceeded
				rr.timeUsed = timeUsage
				rr.memoryUsed = memUsage
				return
			}
			outputFileSize := base.GetFileSize(outputDataDir + "/1.user" )
			if outputFileSize > 3 * standardOutputFileSize {
				cmd.Process.Kill()
				cmd2 := exec.Command("kill", "-9", strconv.Itoa(pid))
				_ = cmd2.Run()
				rr.runResult = base.OutputLimitExceeded
				rr.timeUsed = timeUsage
				rr.memoryUsed = memUsage
				return
			}
			time.Sleep(time.Millisecond * 10)
		}
	}()

	for {
		// 记得 PTRACE_SYSCALL 会在进入和退出syscall时使 tracee 暂停，所以这里用一个变量控制，RAX的内容只打印一遍
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			isAllowSysCall := ptrace.IsAllowSysCall(regs.Orig_rax)
			if isAllowSysCall != base.Normal {
				rr.runResult = isAllowSysCall
				return
			}
			if err != nil {
				break
			}
		}
		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			break
		}
		_, err = syscall.Wait4(pid, nil, 0, nil)

		exit = !exit
	}

	cmd2 := exec.Command("kill", "-9", strconv.Itoa(pid))
	_ = cmd2.Run()


	rr.timeUsed = timeUsage
	rr.memoryUsed = memUsage
	rr.runResult = CompareOneFile(outputDataDir + "/1.user", inputDataDir + "/1.out")

	return
}

