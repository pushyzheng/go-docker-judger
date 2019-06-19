package judger

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"pushy.site/go-docker-judger/conf"
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
func startContainer(id string, caseId string) container.ContainerCreateCreatedBody {
	binds := []string{conf.Volume.GetCodePath(), conf.Volume.GetCasePath()}

	casePath := fmt.Sprintf("case%s.txt", caseId)

	config := &container.Config{
		Image: conf.Container.GetImageName(),
		Cmd:   []string{"sh", "run.sh", casePath},}

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
func StartJudge(id string, caseId string) (*bytes.Buffer, *bytes.Buffer) {
	resp := startContainer(id, caseId)

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

// 获取第一行的输出结果，即为判题的结果标志，有WA/RE/AC...
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

func Run() {
	stdout, stderr := StartJudge("1", "2")

	result := GetJudgeResult(stdout)

	fmt.Println("判题结果：" + result)
	fmt.Println("[out] \n", stdout)
	fmt.Println("[error] \n", stderr)
}
