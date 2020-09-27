package docker

import (
	"fmt"
	"renatofmachado/dock/shell"
)

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
