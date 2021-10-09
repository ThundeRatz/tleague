package docker

import (
	"context"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
)

func BuildImage(path, dockerfile, tag string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	tar, err := archive.TarWithOptions(path, &archive.TarOptions{})
	if err != nil {
		return err
	}

	opts := types.ImageBuildOptions{
		Dockerfile: dockerfile,
		Tags:       []string{tag},
		Remove:     true,
	}

	res, err := cli.ImageBuild(ctx, tar, opts)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return jsonmessage.DisplayJSONMessagesStream(res.Body, os.Stdout, os.Stdout.Fd(), true, nil)
}
