package utils

import (
	"bytes"
	"strings"
)

func GetLineByBytes(buffer *bytes.Buffer, line int) string {
	data := buffer.Bytes()

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

func GetFirstLineByBytes(buffer *bytes.Buffer) string {
	return GetLineByBytes(buffer, 1)
}