package demo

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"os"
)


func GetAllImage(cli *client.Client, ctx context.Context) {
	imgs, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	for _, image := range imgs {
		fmt.Println(image.RepoTags)
	}
}

func GetAllContainer(cli *client.Client, ctx context.Context) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}

func CreateHelloWorld(cli *client.Client, ctx context.Context) {
	// create container from local image
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "ubuntu",
		Cmd: []string{"ls"},
	}, nil, nil, "")

	if err != nil {
		panic(err)
	}
	// start container
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	// get output log
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}