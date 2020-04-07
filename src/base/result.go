package base

type Result struct {
	Result             int64                 `json:"Result"`
	Score              int64                 `json:"Score"`
	CompileErrorInf    string                `json:"CompileErrorInf"`

	TimeUsed           int64                 `json:"TimeUsed"`
	MemoryUsed         int64                 `json:"MemoryUsed"`

	FileName           map[string] int64     `json:"FileName"`
}
