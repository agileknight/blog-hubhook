package listener

import (
	"log"
	"os"

	"github.com/cpuguy83/docker-grand-ambassador/docker"
)

type Blog struct{}

func (b *Blog) Call(msg HubMessage) {

	if msg.Repository.RepoName != ServerConfig.Blog.hub_repo {
		return
	}

	if err := b.buildAndRun(); err != nil {
		log.Println("Could not build blog:", err)
	}
}

func (b *Blog) buildAndRun() error {
	client, err := docker.NewClient(os.Getenv("DOCKER_SOCKET"))
	if err != nil {
		return err
	}

	if err := client.PullImage(ServerConfig.Blog.hub_repo); err != nil {
		return err
	}

	if err := client.RemoveContainer(ServerConfig.Blog.container_name, true, false); err != nil {
		return err
	}

	container := map[string]interface{}{
		"Image":      ServerConfig.Blog.hub_repo,
		"HostConfig": map[string]interface{}{},
		"Name":       ServerConfig.Blog.container_name,
	}

	return client.RunContainer(container)
}
