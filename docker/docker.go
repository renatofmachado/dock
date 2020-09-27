package docker

import (
	"fmt"
	"os/exec"
	"renatofmachado/dock/shell"
	"strings"
)

type Docker struct {
	path string
}

func New() *Docker {
	path, err := exec.LookPath("docker")

	if err != nil {
		panic("Docker is not installed.")
	}

	return &Docker{
		path: path,
	}
}

func (this *Docker) BootContainer(template string, parameters map[string]string) error {
	template = `run -d --name ${:container_name} ` + template

	for key, value := range parameters {
		template = strings.ReplaceAll(template, "${:"+key+"}", value)
	}

	_, err := shell.Exec(this.path, template)

	return err
}

func (this *Docker) EnsureImageIsDownloaded(organization string, image string, tag string) error {
	if !this.isImageDownloaded(organization, image, tag) {
		return this.downloadImage(organization, image, tag)
	}

	return nil
}

func (this *Docker) isImageDownloaded(organization string, image string, tag string) bool {
	_, err := shell.Exec(this.path, fmt.Sprintf("image inspect %s/%s:%s", organization, image, tag))

	if err != nil {
		return false
	}

	return true
}

func (this *Docker) downloadImage(organization string, image string, tag string) error {
	fmt.Println("Downloading Docker image..")

	shell.Exec(this.path, fmt.Sprintf("pull %s/%s:%s", organization, image, tag))

	return nil
}

func (this *Docker) ResolveContainerName(name string, tag string) string {
	return fmt.Sprintf("DOCK--%s--%s", name, tag)
}
