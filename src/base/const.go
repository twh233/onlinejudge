package base

const (
	Normal              int64 = -1  //默认
	Waiting             int64 = 0  //WT	用户程序正在排队等待测试
	Accepted            int64 = 1  //AC	用户程序输出正确的结果
	WrongAnswer         int64 = 2  //WA	用户程序输出错误的结果
	CompilationError    int64 = 3  //CE	用户程序编译错误
	RuntimeError        int64 = 4  //RE	用户程序发生运行时错误
	TimeLimitExceeded   int64 = 5  //TLE	用户程序运行时间超过题目的限制
	MemoryLimitExceeded int64 = 6  //MLE	用户程序运行内存超过题目的限制
	OutputLimitExceeded int64 = 7  //OLE	用户程序输出的结果大大超出正确答案的长度
	PresentationError   int64 = 8  //PE   用户程序输出格式错误
	Danger              int64 = 9  //DE   用户程序带有攻击性
	Running             int64 = 10 //RING	运行中
	SystemError         int64 = 11 //SE	用户程序不能被评测系统正常运行
	Judging             int64 = 12 //JD   用户程序评判中
	Score               int64 = 128 //分数
)

const (
	ICPC          int64 = 0 //ICPC题目类型
	SPJ           int64 = 1 //特判
	Referee       int64 = 2 //交互式
	Functional    int64 = 3 //单个函数 如：Leetcode
)

const (
	C       string = ".c" //c
	CPP     string = ".cpp" //c++
	Python  string = ".python" //python
	Go      string = ".go" //golang
	Java    string = ".java" //Java
	CSharp  string = "" //C#
)

const (
	RootDir  string = "/root"
	JudgeDir string = "/judge_file/judge"
	DataDir  string = "/judge_file/data"
)

