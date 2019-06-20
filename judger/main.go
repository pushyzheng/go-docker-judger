package judger

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"io/ioutil"
	"pushy.site/go-docker-judger/conf"
	"pushy.site/go-docker-judger/models"
	"strings"
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
	binds := []string{conf.Volume.GetCodePath(), conf.Volume.GetCasePath()}

	//codePath := fmt.Sprintf("code/%s/Main.java", task.UserId)
	casePath := fmt.Sprintf("../../cases/case%s.txt", "2")

	config := &container.Config{
		Image: conf.Container.GetImageName(),
		Cmd:   []string{"sh", "run.sh", task.UserId, casePath},}

	// create container from local image
	resp, err := cli.ContainerCreate(ctx, config,
		&container.HostConfig{Binds: binds},
		nil, "")

	if err != nil {
		panic(err)
	}
	// start container
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
	data := stdout.Bytes()

	var result strings.Builder
	var resultEnd int
	for i, each := range data {
		if each == '\n' {
			resultEnd = i
			break
		}
	}
	result.Write(data[:resultEnd])
	return result.String()
}

func Run(task models.JudgementTask) (string, string) {
	fmt.Println("Start a judgement task :")
	fmt.Println(task)

	stdout, stderr := StartJudge(task)
	result := GetJudgeResult(stdout)
	if result != "" {  // 编译错误或者运行时异常
		fmt.Println(result)
		return result, stderr.String()
	}

	outputPath := fmt.Sprintf("%s/%s/result.txt", conf.Volume.CodeHostPath, task.UserId)
	outputBytes, err := ioutil.ReadFile(outputPath)
	if err != nil {
		fmt.Println("The output path not found")
	}
	output := string(outputBytes)

	answerPath := fmt.Sprintf("%s/answer_%s.txt", conf.Volume.AnswerHostPath, "1")
	answerBytes, err := ioutil.ReadFile(answerPath)
	if err != nil {
		fmt.Println("The answer path not found")
	}
	answer := string(answerBytes)

	if output == answer {
		result = models.AC
	} else {
		result = models.WA
	}

	return result, ""
}
