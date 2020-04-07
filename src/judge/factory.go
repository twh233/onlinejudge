package judge

import (
	"base"
	"fmt"
)

type Factory struct {

}

func NewFactory() *Factory {
	return &Factory{}
}

func (this *Factory)CreateIniter(submit base.Submit) Initer {
	inputDir := base.RootDir + base.DataDir + "/" + submit.ProblemId + "/insert" + submit.Language
	outputDir := base.RootDir + base.JudgeDir + "/" + submit.SubmitId + "/Main" + submit.Language
	switch submit.ProblemType {
		case base.Functional:
			insertCode := &InsertCode{}
			insertCode.inputDir = inputDir
			insertCode.code = submit.UserCode
			insertCode.outputDir = outputDir
			return insertCode
		case base.ICPC, base.SPJ, base.Referee:
			initNode := &InitNode{}
			initNode.outputDir = outputDir
			initNode.code = submit.UserCode
			return initNode
		default:
			fmt.Println("无效题型")
			return nil
	}
}

func (this *Factory)CreateCompiler(submit base.Submit) Compiler {
	dir := base.RootDir + base.JudgeDir + "/" + submit.SubmitId
	switch  submit.Language {
		case base.C:
			c := &C{}
			c.codeDir = dir
			return c
		case base.CPP:
			cpp := &CPP{}
			cpp.codeDir = dir
			return cpp
		case base.Java:
			java := &JAVA{}
			java.codeDir = dir
			return java
		default:
			fmt.Println("无效语言")
			return nil
	}
}

func (this *Factory)CreateRunner(submit base.Submit) Runner {
	codeDir := base.RootDir + base.JudgeDir + "/" + submit.SubmitId
	dataDir := base.RootDir + base.DataDir + "/" + submit.ProblemId

	switch submit.ProblemType {
		case base.Referee:
			dn := &DoNothing{}
			return dn
	}

	switch submit.Language {
		case base.C:
			c := &C{}
			c.codeDir = codeDir
			c.dataDir = dataDir
			return c
		case base.CPP:
			cpp := &CPP{}
			cpp.codeDir = codeDir
			cpp.dataDir = dataDir
			return cpp
		case base.Java:
			java := &JAVA{}
			java.codeDir = codeDir
			java.dataDir = dataDir
			return java
		default:
			fmt.Println("无效语言")
			return nil
	}
}

func (this *Factory)CreateComparer(submit base.Submit) Comparer {
	userOutputDir := base.RootDir + base.JudgeDir + "/" + submit.SubmitId
	standardOutputDir := base.RootDir + base.DataDir + "/" + submit.ProblemId
	switch submit.ProblemType {
		case base.ICPC, base.Functional:
			dn := &DoNothing{}
			return dn
		case base.SPJ:
			spj := &SPJ{}
			spj.spjLanuage = submit.Language
			spj.spjDir = standardOutputDir
			spj.userOutputDir = userOutputDir
			return spj
	case base.Referee:
			refree := &Refree{}
			refree.lanuage = submit.Language
			refree.codeDir = userOutputDir
			refree.refreeDir = standardOutputDir
			return refree
	default:
		fmt.Println("无效类型")
		return nil
	}
}