package listener

import (
	"log"
	"os"

	"github.com/cpuguy83/docker-grand-ambassador/docker"
)

type Blog struct{}

func (b *Blog) Call(msg HubMessage) {

	if msg.Repository.RepoName != ServerConfig.Blog.Repo {
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

	if err := client.PullImage(ServerConfig.Blog.Repo); err != nil {
		return err
	}

	if err := client.RemoveContainer(ServerConfig.Blog.Name, true, false); err != nil {
		return err
	}

	container := map[string]interface{}{
		"Image":      ServerConfig.Blog.Repo,
		"HostConfig": map[string]interface{}{},
		"Name":       ServerConfig.Blog.Name,
	}

	return client.RunContainer(container)
}
