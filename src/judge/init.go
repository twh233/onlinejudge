package judge

import (
	"base"
	"bufio"
	"io"
	"os"
	"strings"
)

type Initer interface {
	Init() int64
}

type InitNode struct {
	outputDir string
	code string
}

func (initNode *InitNode)Init() int64 {
	out, err := os.OpenFile(initNode.outputDir, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0)
	if err != nil {
		return base.SystemError
	}
	defer  out.Close()
	_, err = out.WriteString(initNode.code)
	if err != nil {
		return base.SystemError
	}
	return base.Normal
}

type InsertCode struct{
	inputDir string
	outputDir string
	code string
}

func (insertCode *InsertCode)Init() int64{
	//read
	in, err := os.Open(insertCode.inputDir)
	if err != nil{
		return base.SystemError
	}
	defer  in.Close()
	reader := bufio.NewReader(in)
	//write
	out, err := os.OpenFile(insertCode.outputDir, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0)
	if err != nil {
		return base.SystemError
	}
	defer  out.Close()
	//insert
	for{
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if line != "EOF" {
				_, err = out.WriteString(line)
				if err != nil {
					return base.SystemError
				}
			}
			break
		}
		if strings.Contains(line, "__CODE__"){
			line = insertCode.code
		}
		_, err = out.WriteString(line)
		if err != nil{
			return base.SystemError
		}
	}
	return base.Normal
}