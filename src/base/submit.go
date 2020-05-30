package base

type Submit struct{
	SubmitId      string  `json:"SubmitId"`
	ProblemId     string  `json:"ProblemId"`
	ProblemType   int64   `json:"ProblemType"`
	UserCode      string  `json:"UserCode"`
	Language      string  `json:"Language"`
	TimeLimit     int64   `json:"TimeLimit"`
	MemoryLimit   int64   `json:"MemoryLimit"`
}

func ReturnResult(result int64) string {
	var result_map map[int64]string
	result_map = make(map[int64]string)

	result_map[-1] = "Normal"
	result_map[0] = "Waiting"
	result_map[1] = "Accepted"
	result_map[2] = "WrongAnswer"
	result_map[3] = "CompilationError"
	result_map[4] = "RuntimeError"
	result_map[5] = "TimeLimitExceeded"
	result_map[6] = "MemoryLimitExceeded"
	result_map[7] = "OutputLimitExceeded"
	result_map[8] = "PresentationError"
	result_map[9] = "Danger"
	result_map[10] = "Running"
	result_map[11] = "SystemError"
	result_map[12] = "Judging"
	result_map[128] = "Score"

	return result_map[result]
}