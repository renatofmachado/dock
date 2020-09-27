package docker

import "strings"

func (this *Docker) ResolveTag(serviceName string) string {
	split := strings.Split(serviceName, "@")

	if len(split) > 1 {
		return split[1]
	}

	return "latest"
}
