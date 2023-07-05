package docker

import (
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"os/exec"
)

func initJSEnvironment() error {

	images, err := cli.client.ImageList(cli.ctx, types.ImageListOptions{})
	if err != nil {
		return err
	}

	var pullFlag = true
	for _, image := range images {
		if len(image.RepoTags) > 0 && image.RepoTags[0] == "node:latest" {
			pullFlag = false
			break
		}
	}

	if pullFlag {
		_, err := cli.client.ImagePull(cli.ctx, "node:latest", types.ImagePullOptions{})
		if err != nil {
			return err
		}
	}

	containers, err := cli.client.ContainerList(cli.ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return err
	}

	for _, container := range containers {
		if container.Names[0] == "/js" {
			err = cli.client.ContainerStart(cli.ctx, container.ID, types.ContainerStartOptions{})
			if err != nil {
				return err
			}
			return nil
		}
	}

	resp, err := cli.client.ContainerCreate(cli.ctx, &container.Config{
		Image: "node:latest",
		Tty:   true,
	}, nil, nil, nil, "js")
	if err != nil {
		return err
	}

	err = cli.client.ContainerStart(cli.ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	return nil
}

func CheckJS(content, answer string) error {
	cmd := exec.Command("docker", "exec", "js", "node", "-e", content)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	if string(output[0:len(output)-1]) != answer {
		return errors.New("answer not match")
	}

	return nil
}
