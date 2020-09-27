package docker

import (
	"errors"
	"fmt"
	"renatofmachado/dock/shell"
	"strings"
)

type container struct {
	id     string
	name   string
	status string
	ports  string
}

func (this *Docker) BootContainer(template string, parameters map[string]string) error {
	template = `run -d --name ${:container_name} ` + template

	for key, value := range parameters {
		template = strings.ReplaceAll(template, "${:"+key+"}", value)
	}

	_, err := shell.Exec(this.path, template)

	return err
}

func (this *Docker) StartContainer(containerID string) error {
	_, err := shell.Exec(this.path, fmt.Sprintf("start %s", containerID))

	return err
}

func (this *Docker) StopContainer(containerID string) error {
	_, err := shell.Exec(this.path, fmt.Sprintf("stop %s", containerID))

	return err
}

func (this *Docker) ListContainers() ([]*container, error) {
	output, err := shell.Exec(this.path, "ps -a --filter \"name=DOCK-\" --format \"table {{.ID}}|{{.Names}}|{{.Status}}|{{.Ports}}\"")

	if err != nil {
		return []*container{}, err
	}

	containers := strings.Split(output, "\n")

	if len(containers) == 1 {
		return []*container{}, errors.New("There are no available containers")
	}

	// Ignore the first line of headers
	containers = containers[1:]

	list := []*container{}

	for _, line := range containers {
		if line == "" {
			continue
		}

		containerInfo := strings.Split(line, "|")

		list = append(list, &container{
			id:     containerInfo[0],
			name:   containerInfo[1],
			status: containerInfo[2],
			ports:  containerInfo[3],
		})
	}

	return list, nil
}

func (this *Docker) FindContainerId(name string, tag string) (string, error) {
	containers, err := this.ListContainers()

	if err != nil {
		return "", err
	}

	containerName := this.ResolveContainerName(name, tag)

	for _, container := range containers {
		if container.name == containerName {
			return container.id, nil
		}
	}

	return "", errors.New("Container not found")
}

func (this *Docker) ResolveContainerName(name string, tag string) string {
	return fmt.Sprintf("DOCK--%s--%s", name, tag)
}
