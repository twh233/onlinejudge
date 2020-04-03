package judge

import (
	"bytes"
	"fmt"
	"os/exec"
	"base"
)

type Compiler interface {
	Compile() CompileResult
}

type CompileResult struct{
	compileErrorInf string
	compileResult int64
}

type C struct{
	codeDir string
	dataDir string
}

func (c *C)Compile() CompileResult{
	return CompileAdapter(base.C, c.codeDir, "Main")
}

type CPP struct{
	codeDir string
	dataDir string
}

func (cpp *CPP)Compile() CompileResult{
	return CompileAdapter(base.CPP, cpp.codeDir, "Main")
}

type JAVA struct{
	codeDir string
	dataDir string
}

func (java *JAVA)Compile() CompileResult{
	return CompileAdapter(base.Java, java.codeDir, "Main")
}

func CompileAdapter(language string, codeDir string, fileName string) (cr CompileResult){
	var cmd *exec.Cmd
	switch language {
		case base.C:
			cmd = exec.Command("g++", fileName + language, "-o", fileName)
		case base.CPP:
			cmd = exec.Command("g++","-std=c++11", fileName + language, "-o", fileName)
		case base.Java:
			cmd = exec.Command("javac", fileName + language)
		default:
			fmt.Println("无效语言")
			cr.compileResult = base.SystemError
			return
	}
	cmd.Dir = codeDir
	errInf := bytes.NewBuffer(nil)
	cmd.Stderr = errInf

	err := cmd.Run()
	if err != nil{
		//compile error information  --- string(errInf.Bytes())
		cr.compileResult = base.CompilationError
		cr.compileErrorInf = string(errInf.Bytes())
		return
	}
	cr.compileResult = base.Normal
	return
}