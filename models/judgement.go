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
	Id            string  `json:"id"`             // 对应的任务ID
	Succeed       bool    `json:"succeed"`        // 是否成功判题
	Status        string  `json:"status"`         // 判题结果
	RuntimeTime   float64 `json:"runtime_time"`   // 运行的时长
	RuntimeMemory int     `json:"runtime_memory"` // 占用的内存

	WrongLine      int    `json:"wrong_line"`      // 错误的行数
	LastInput      string `json:"last_input"`      // 最后输入
	LastOutput     string `json:"last_output"`     // 最后输出
	ExpectedOutput string `json:"expected_output"` // 期望的正确输出
	ErrorInfo      string `json:"error_info"`      // 错误信息

	Timestamp time.Time `json:"timestamp"`
}

func (result *JudgementResult) ToJsonString() []byte {
	data, _ := json.Marshal(result)
	return data
}
