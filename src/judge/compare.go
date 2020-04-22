package judge

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"base"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

type Comparer interface{
	Compare(limit ResourcesLimit) CompareResult
}

type CompareResult struct{
	fileName map[string]int64

	compareResult int64
	timeUsed int64
	memoryUsed int64
}

type ResultFileCompare struct{
	standardOutputDir string
	userOutputDir     string
}

func (dn *DoNothing)Compare(limit ResourcesLimit) (cr CompareResult){
	cr.compareResult = base.Accepted
	return
}

func (resultFileCompare *ResultFileCompare)Compare(limit ResourcesLimit) (cr CompareResult) {
	files, err := ioutil.ReadDir(resultFileCompare.standardOutputDir)
	if err != nil{
		cr.compareResult = base.SystemError
		return
	}

	for _, f := range files{
		if strings.Contains(f.Name(), ".out") {
			if strings.Compare(f.Name(), "000000.out") == 0{
				continue
			}
			fileName := strings.Replace(f.Name(), ".out", ".user", -1)
			compareResult := CompareOneFile(resultFileCompare.standardOutputDir + "/" + f.Name(), resultFileCompare.userOutputDir+ "/" + fileName)
			if compareResult != base.Accepted {
				cr.compareResult = compareResult
				return
			}
		}
	}
	cr.compareResult = base.Accepted
	return
}

type SPJ struct{
	spjLanuage string
	spjDir string
	userOutputDir string
}

func (spj *SPJ)Compare(limit ResourcesLimit) (cr CompareResult){
	//compile spj file
	ce := CompileAdapter(spj.spjLanuage, spj.spjDir, "spj")
	if ce.compileResult == base.SystemError {
		cr.compareResult = base.SystemError
		return
	}else if ce.compileResult == base.CompilationError{
		cr.compareResult = base.SystemError
		return
	}
	//get user output files
	files, err := ioutil.ReadDir(spj.userOutputDir)
	if err != nil{
		cr.compareResult = base.SystemError
		return
	}
	// run spj file
	var scoreSum int64 = 0
	var fileCount int64 = 0
	cr.fileName = make(map[string]int64)
	for _, f := range files{
		if strings.Contains(f.Name(), ".user"){
			if strings.Compare(f.Name(), "000000.user") == 0{
				continue
			}
			inputFileName := strings.Replace(f.Name(), ".user", ".in", -1)
			outputFileName := strings.Replace(f.Name(), ".user", ".out", -1)
			result := RunSPJ(spj.spjLanuage, spj.spjDir, spj.spjDir + "/" + inputFileName, spj.spjDir + "/" + outputFileName, spj.userOutputDir + "/" + f.Name())
			fileName := strings.Replace(f.Name(), ".user", ".in", -1)
			cr.fileName[fileName] = result
			if result != base.Accepted && result < base.Score{
				cr.compareResult = result
				return
			}
			fileCount = fileCount + 1
			scoreSum = scoreSum + result
		}
	}
	cr.compareResult = scoreSum / fileCount
	return
}

type Refree struct{
	codeDir string
	refreeDir string
	language string
}

func (refree *Refree)Compare(limit ResourcesLimit) (cr CompareResult){
	//compile refree file
	ce := CompileAdapter(refree.language, refree.refreeDir, "refree")
	if ce.compileResult == base.SystemError {
		cr.compareResult = base.SystemError
		return
	}else if ce.compileResult == base.CompilationError{
		cr.compareResult = base.SystemError
		return
	}

	return RunRefree(refree.language, refree.refreeDir, refree.codeDir, limit)
}

func RunSPJ(language string, spjDir string, inputDir string, outputDir string, userDir string) int64 {
	var cmd *exec.Cmd
	switch language {
	case base.C, base.CPP:
		cmd = exec.Command(spjDir + "/spj", inputDir, outputDir, userDir)
	case base.Java:
		cmd = exec.Command("java", "spj", inputDir, outputDir, userDir)
	default:
		return base.SystemError
	}
	cmd.Dir = spjDir

	_ = cmd.Run()

	result := cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	if result == 0{
		result = 1
	}else if result == 4{
		result = 2
	}
	return int64(result)
}

func RunRefree(language string, refreeDir string, codeDir string, limit ResourcesLimit) (cr CompareResult){
	var refreeCmd *exec.Cmd
	var codeCmd *exec.Cmd
	switch language {
	case base.C, base.CPP:
		refreeCmd = exec.Command(refreeDir + "/refree")
		codeCmd = exec.Command(codeDir + "/Main")
	case base.Java:
		refreeCmd = exec.Command("java", "refree")
		codeCmd = exec.Command("java", "Main")
	default:
		cr.compareResult = base.SystemError
		return
	}
	refreeCmd.Dir = refreeDir
	codeCmd.Dir = codeDir
	//codeCmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	refreeOutputPipe, _ := refreeCmd.StdoutPipe()
	codeCmd.Stdin = refreeOutputPipe

	codeOutputPipe, _ := codeCmd.StdoutPipe()
	refreeCmd.Stdin = codeOutputPipe

	_ = refreeCmd.Start()
	_ = codeCmd.Start()

	go func() {
		// workaround 不知道为什么标准的wait不行
		var rusage syscall.Rusage
		_, _, err := wait(codeCmd.Process.Pid, &rusage)
		if err != nil {
			fmt.Println("wait error", err)
		}
	}()

	_ = refreeCmd.Wait()
	/*go func() {
		// workaround 不知道为什么标准的wait不行
		var rusage syscall.Rusage
		_, _, err := wait(refreeCmd.Process.Pid, &rusage)
		if err != nil {
			fmt.Println("wait error", err)
		}
	}()*/


	var timeUsage int64 = 0
	var memUsage int64 = 0

	for {
		ok, rss, runningTime, cpuTime := base.GetResourceUsage(codeCmd.Process.Pid)
		if !ok {
			break
		}

		timeUsage = base.GetMaxInt64(timeUsage, cpuTime); if timeUsage > limit.timeLimit || runningTime > 3 * limit.timeLimit {
			codeCmd.Process.Kill()
			refreeCmd.Process.Kill()
			cmd2 := exec.Command("kill", "-9", strconv.Itoa(codeCmd.Process.Pid))
			_ = cmd2.Run()
			cmd2 = exec.Command("kill", "-9", strconv.Itoa(refreeCmd.Process.Pid))
			_ = cmd2.Run()
			cr.compareResult = base.TimeLimitExceeded
			cr.timeUsed = timeUsage
			cr.memoryUsed = memUsage
			return
		}
		memUsage = base.GetMaxInt64(memUsage, rss); if memUsage*3 > limit.memoryLimit*2 {
			codeCmd.Process.Kill()
			refreeCmd.Process.Kill()
			cmd2 := exec.Command("kill", "-9", strconv.Itoa(codeCmd.Process.Pid))
			_ = cmd2.Run()
			cmd2 = exec.Command("kill", "-9", strconv.Itoa(refreeCmd.Process.Pid))
			_ = cmd2.Run()
			cr.compareResult = base.MemoryLimitExceeded
			cr.timeUsed = timeUsage
			cr.memoryUsed = memUsage
			return
		}
	}

	result := refreeCmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
	cr.compareResult = int64(result)
	cr.timeUsed = timeUsage
	cr.memoryUsed = memUsage

	return
}

func IsSameWithoutInvisableChar(dir1 string, dir2 string) int64{
	byte1 := base.ReadAll(dir1)
	byte2 := base.ReadAll(dir2)
	str1 := string(byte1)
	str2 := string(byte2)

	//remove str1's invaisable char
	str1 = strings.Replace(str1, "\t", "", -1)
	str1 = strings.Replace(str1, "\n", "", -1)
	str1 = strings.Replace(str1, " ", "", -1)
	str1 = strings.Replace(str1, "\r", "", -1)

	str2 = strings.Replace(str2, "\t", "", -1)
	str2 = strings.Replace(str2, "\n", "", -1)
	str2 = strings.Replace(str2, " ", "", -1)
	str2 = strings.Replace(str2, "\r", "", -1)

	if strings.Compare(str1, str2) == 0{
		return base.Accepted
	}else {
		return base.WrongAnswer
	}
}

func CompareOneFile(dir1 string, dir2 string) int64{
	if IsSameWithoutInvisableChar(dir1, dir2) != base.Accepted{
		return base.WrongAnswer
	}

	file1, err := os.Open(dir1)
	defer file1.Close()
	if err != nil {
		return base.SystemError
	}

	file2, err := os.Open(dir2)
	defer file2.Close()
	if err != nil {
		return base.SystemError
	}

	reader1 := bufio.NewReader(file1)
	reader2 := bufio.NewReader(file2)

	for {
		str1, err1 := reader1.ReadString('\n')
		str2, err2 := reader2.ReadString('\n')

		str1 = strings.Replace(str1, "\r", "", -1)
		str2 = strings.Replace(str2, "\r", "", -1)
		str1 = strings.Replace(str1, "\n", "", -1)
		str2 = strings.Replace(str2, "\n", "", -1)
		str1 = strings.TrimSpace(str1)
		str2 = strings.TrimSpace(str2)
		//fmt.Println(str1)
		//fmt.Println(str2)

		if strings.Compare(str1, str2) != 0 {
			return base.PresentationError
		}

		if err1 == io.EOF || err2 == io.EOF {
			break
			/*if err1 == io.EOF && err2 == io.EOF {
				break
			}
			if err1 == io.EOF{
				str2, err2 := reader2.ReadString('\n')
				str2 = strings.Replace(str2, "\r", "", -1)
				str2 = strings.Replace(str2, "\n", "", -1)
				if str2 == "" && err2 == io.EOF{
					break
				}
				return base.PresentationError
			} else if err2 == io.EOF {
				str1, err1 := reader1.ReadString('\n')
				str1 = strings.Replace(str1, "\r", "", -1)
				str1 = strings.Replace(str1, "\n", "", -1)
				if str1 == "" && err1 == io.EOF{
					break
				}
				return base.PresentationError
			}*/
		}
	}
	return base.Accepted
}

func wait(pid int, rusage *syscall.Rusage) (int, *syscall.WaitStatus, error) {
	var status syscall.WaitStatus
	var siginfo [128]byte
	// 阻止到wai4执行成功
	psig := &siginfo[0]
	_, _, e := syscall.Syscall6(syscall.SYS_WAITID, 1, uintptr(pid), uintptr(unsafe.Pointer(psig)), syscall.WEXITED|syscall.WNOWAIT, 0, 0)
	// psig可能被回收
	// keepalive
	// wait return
	runtime.KeepAlive(psig)
	if e != 0 {
		if e != syscall.ENOSYS {
			return 0, nil, os.NewSyscallError("waitid", e)
		}
	}
	wpid, err := syscall.Wait4(pid, &status, 0, rusage) // for rusage collect
	if err != nil {
		return 0, nil, err
	}
	return wpid, &status, err
}
