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
