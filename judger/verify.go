package judger

import (
	"fmt"
	"io/ioutil"
	"log"
	"pushy.site/go-docker-judger/conf"
	"pushy.site/go-docker-judger/models"
	"pushy.site/go-docker-judger/utils"
)

// 读取文件的bytes
func ReadBytesByPath(path string) []byte {
	outputBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return outputBytes
}

// 校验答案的主方法
func VerifyAnswer(task models.JudgementTask) (models.VerificationResult, error) {
	var result models.VerificationResult

	outputPath := fmt.Sprintf("%s/%s/result.txt", conf.Volume.CodeHostPath, task.UserId)
	outputBytes := ReadBytesByPath(outputPath)
	if outputBytes == nil {
		return result, fmt.Errorf("OUTPUT PATH NOT FOUND")
	}

	answerPath := fmt.Sprintf("%s/answer_%d.txt", conf.Volume.AnswerHostPath, task.ProblemId)
	answerBytes := ReadBytesByPath(answerPath)
	if answerBytes == nil {
		return result, fmt.Errorf("ANSWER PATH NOT FOUND")
	}

	wrongLine := doVerify(outputBytes, answerBytes, &result)
	log.Println("wrong line: ", wrongLine)

	if result.Status == models.WA {
		result.LastInput = getLastInput(task.ProblemId, wrongLine)
	}

	return result, nil
}

// 获取测试样例中对应的行数即为最后的输入
func getLastInput(problemId int, wrongLine int) string {
	casePath := fmt.Sprintf("%s/case_%d.txt", conf.Volume.CaseHostPath, problemId)
	data := ReadBytesByPath(casePath)
	return utils.GetLineByBytes(data, wrongLine)
}

// 逐行校验用户程序输出与标准答案
func doVerify(outputBytes []byte, answerBytes []byte, result *models.VerificationResult) int {
	count := utils.GetLineCountByBytes(answerBytes)
	log.Println("Total line: ", count)

	for i := 1; i <= count; i++ {
		outputLine := utils.GetLineByBytes(outputBytes, i)
		answerLine := utils.GetLineByBytes(answerBytes, i)
		log.Println("output: ", outputLine)
		log.Println("answer: ", answerLine)

		if outputLine != answerLine {
			result.Status = models.WA
			result.ExpectedOutput = answerLine
			result.LastOutput = outputLine

			return i
		}
	}
	result.Status = models.AC

	return -1
}