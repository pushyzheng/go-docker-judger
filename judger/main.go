package judger

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"go-docker-judger/conf"
	"go-docker-judger/models"
	"go-docker-judger/utils"
	"log"
	"strconv"
)

var cli *client.Client
var ctx = context.Background()

func InitCore() {
	var err error
	cli, err = client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}
}

// 启动容器
func startContainer(task models.JudgementTask) container.ContainerCreateCreatedBody {
	v2Path := "e:/usr/scripts"
	v2Volume := fmt.Sprintf("%s:/root", v2Path)

	binds := []string{conf.Volume.GetCodePath(), conf.Volume.GetCasePath(), v2Volume}

	casePath := fmt.Sprintf("../../cases/case_%d.txt", task.ProblemId)

	config := &container.Config{
		Image: conf.Container.GetImageName(),
		Cmd:   []string{"/bin/bash", getScriptName(task.Language), task.UserId, casePath}}

	// 容器资源配置，内存限制、CPU限制等
	memoryLimit := int64(task.MemoryLimit * 1024 * 1024)
	resourceConfig := container.Resources{Memory: memoryLimit}

	// 从本地镜像中创建容器，并传入配置选项
	resp, err := cli.ContainerCreate(ctx, config,
		&container.HostConfig{Binds: binds, Resources: resourceConfig},
		nil, nil, "")

	if err != nil {
		panic(err)
	}
	// 启动容器
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	return resp
}

// 运行容器，获取输出
func StartJudge(task models.JudgementTask) (*bytes.Buffer, *bytes.Buffer) {
	resp := startContainer(task)

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	// get output log
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true, ShowStderr: true})
	if err != nil {
		panic(err)
	}

	// get stdout and stderr
	stdout, stderr := new(bytes.Buffer), new(bytes.Buffer)
	_, err = stdcopy.StdCopy(stdout, stderr, out)
	if err != nil {
		panic(err)
	}
	return stdout, stderr
}

// 获取第一行的输出结果，即为判题的结果标志
func GetJudgeResult(stdout *bytes.Buffer) string {
	return utils.GetFirstLineByBytes(stdout.Bytes())
}

// 获取程序运行时间（不包括编译的时间）
func GetRuntimeTime(stderr *bytes.Buffer) float64 {
	line := utils.GetFirstLineByBytes(stderr.Bytes())
	timeStr := line[8 : len(line)-1]
	time, err := strconv.ParseFloat(timeStr, 64)
	if err != nil {
		return -1.0
	}
	return time
}

func Run(task models.JudgementTask, result *models.JudgementResult) {
	log.Println("Start a judgement task : ", task)

	stdout, stderr := StartJudge(task)
	status := GetJudgeResult(stdout)

	log.Println("[stdout]:\n ", stdout.String())
	log.Println("[stderr]:\n ", stderr.String())

	// 判断编译错误或者运行时异常
	if status == models.CE || status == models.RE {
		log.Println("Runtime exception or compile error")
		if utils.GetFirstLineByBytes(stderr.Bytes()) == "Killed" { // 超出内存限制判断
			status = models.MLE
			log.Println("Memory limit exceed")
		}
		result.Status = status
		result.ErrorInfo = stderr.String()
		return
	}

	// 校验是否超时
	time := GetRuntimeTime(stderr)
	log.Println("Runtime time is: ", time)
	if time > float64(task.TimeLimit) {
		log.Println("Time limit exceed error")
		result.Status = models.TLE
		return
	}

	// 校验用户程序输出的答案是否和标准答案一致
	err := VerifyAnswer(task, result)
	if err != nil {
		result.Status = models.SE
		result.ErrorInfo = err.Error()
		return
	}

	result.RuntimeTime = time
}

func getScriptName(language string) string {
	if language == "java" {
		return "run_java.sh"
	} else if language == "python" {
		return "run_python.sh"
	} else if language == "js" {
		return "run_js.sh"
	} else if language == "c" || language == "cpp" {
		return "run_c.sh"
	}
	return ""
}
