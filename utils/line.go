package utils

import (
	"strings"
)

// 获取行数
func GetLineCountByBytes(data []byte) int {
	if len(data) == 0 {
		return 0
	}
	result := 1
	for i, each := range data {
		if each == '\n' && i != len(data) - 1 {
			result++
		}
	}
	return result
}

// 获取前几行
func GetPreLineByBytes(data []byte, line int) string {
	if len(data) == 0 {
		return ""
	}
	if GetLineCountByBytes(data) < line {
		return ""
	}
	var result strings.Builder
	var resultEnd int

	index := 1
	for i, each := range data {
		if each == '\n' {
			if line == index {
				resultEnd = i
				break
			}
			index++
		}
	}
	result.Write(data[:resultEnd])
	return result.String()
}

// 获取第几行
func GetLineByBytes(data []byte, line int) string {
	if len(data) == 0 {
		return ""
	}
	if GetLineCountByBytes(data) < line {
		return ""
	}

	var result strings.Builder
	var start, end int

	index := 1
	for i, each := range data {
		if each == '\n' {
			if line == index {
				end = i
				break
			}
			start = i + 1
			index++
		}
	}
	result.Write(data[start:end])
	return result.String()
}

// 获取第一行
func GetFirstLineByBytes(data []byte) string {
	return GetPreLineByBytes(data, 1)
}
