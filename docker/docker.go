package docker

import (
	"os/exec"
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
