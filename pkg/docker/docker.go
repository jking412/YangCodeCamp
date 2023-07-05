package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
)

type Client struct {
	client *client.Client
	ctx    context.Context
}

var cli *Client

func init() {
	dockerCli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println(err)
	}

	cli = &Client{
		client: dockerCli,
		ctx:    context.Background(),
	}
	err = initJSEnvironment()
	if err != nil {
		panic(err)
	}
	fmt.Println("docker init success")
}
