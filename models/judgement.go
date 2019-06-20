package models

import (
	"encoding/json"
	"time"
)

type JudgementTask struct {
	Id          string `json:"id"`           // 任务ID
	ProblemId   int    `json:"problem_Id"`   // 问题ID
	UserId      string `json:"user_id"`      // 用户ID
	Language    string `json:"language"`     // 代码语言
	TimeLimit   int    `json:"time_limit"`   // 时间限制，单位为s
	MemoryLimit int    `json:"memory_limit"` // 内存限制，单位为MB

	//Timestamp time.Time `json:"timestamp"`
}

type JudgementResult struct {
	Id        string `json:"id"`        // 对应的任务ID
	Succeed   bool   `json:"succeed"`   // 是否成功判题
	Result    string `json:"result"`    // 判题结果
	Duration  int    `json:"duration"`  // 运行的时长
	Memory    int    `json:"memory"`    // 占用的内存
	ErrorInfo string `json:"errorInfo"` // 错误信息

	Timestamp time.Time `json:"timestamp"`
}

func (result *JudgementResult) ToJsonString() []byte {
	data, _ := json.Marshal(result)
	return data
}
