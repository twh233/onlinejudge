package judge

import (
	"base"
	"fmt"
)

type Judger struct {
	initer Initer
	compiler Compiler
	runner Runner
	comparer Comparer
}

func (judger *Judger)Run(timeLimit int64, memoryLimit int64, problemType int64) (result base.Result) {
	//init
	initResult := judger.initer.Init()
	if initResult != base.Normal {
		result.Result = initResult
		return
	}
	fmt.Println("init Success")
	//compile
	compileResult := judger.compiler.Compile()
	if compileResult.compileResult == base.SystemError {
		result.Result = compileResult.compileResult
		return
	}else if compileResult.compileResult == base.CompilationError {
		result.Result = compileResult.compileResult
		result.CompileErrorInf = compileResult.compileErrorInf
		return
	}
	fmt.Println("compile Success")
	limit := ResourcesLimit{}
	limit.memoryLimit = memoryLimit
	limit.timeLimit = timeLimit
	limit.problemType = problemType
	//run
	fmt.Println("judge Test run :" , limit)
	runResult := judger.runner.Run(limit)
	result.FileName = make(map[string]int64)

	result.TimeUsed = runResult.timeUsed
	result.MemoryUsed = runResult.memoryUsed
	result.FileName = runResult.fileName
	if runResult.runResult != base.Accepted && runResult.runResult < base.Score {
		result.Result = runResult.runResult
		return
	}
	fmt.Println("run Success")

	//compare
	compareResult := judger.comparer.Compare(limit)
	result.Result = compareResult.compareResult
	if problemType == base.SPJ {
		result.FileName = compareResult.fileName
	} else if problemType == base.Referee {
		result.TimeUsed = compareResult.timeUsed
		result.MemoryUsed = compareResult.memoryUsed
	}
	fmt.Println("compare Success")
	return
}

func CreateJudger(submit base.Submit) *Judger{
	factory := NewFactory()
	judger := &Judger{}
	judger.initer = factory.CreateIniter(submit); if judger.initer == nil{
		return nil
	}
	judger.compiler = factory.CreateCompiler(submit); if judger.compiler == nil{
		return nil
	}
	judger.runner = factory.CreateRunner(submit); if judger.runner == nil{
		return nil
	}
	judger.comparer = factory.CreateComparer(submit); if judger.comparer == nil{
		return nil
	}
	return judger
}