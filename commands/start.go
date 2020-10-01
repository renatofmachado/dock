package commands

import (
	"errors"
	"fmt"
	"renatofmachado/dock/docker"
	"renatofmachado/dock/services"

	"github.com/urfave/cli/v2"
)

func Start(c *cli.Context) error {
	serviceName := c.Args().First()

	if serviceName == "" {
		return errors.New("No service name was received")
	}

	docker := docker.New()

	service, err := services.Resolve(serviceName)

	tag := docker.ResolveTag(service.Name)

	containerID, err := docker.FindContainerId(service.Name, tag)

	if err != nil {
		Enable(c)

		containerID, err = docker.FindContainerId(service.Name, tag)

		if err != nil {
			return err
		}
	}

	fmt.Printf("Starting %s@%s (ID: %s)\n", service.Name, tag, containerID)
	return docker.StartContainer(containerID)
}
