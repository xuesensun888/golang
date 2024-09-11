package models

type Solution struct {
	ID           string `json:"id"`
	FaultTitle   string `json:"fault_title"`   // 与故障标题匹配
	SolutionText string `json:"solution_text"` // 解决方案的描述
}
