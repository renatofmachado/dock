package commands

import (
	"bufio"
	"fmt"
	"os"
	"renatofmachado/dock/docker"
	"renatofmachado/dock/services"
	"strings"

	"github.com/urfave/cli/v2"
)

func Enable(c *cli.Context) error {
	docker := docker.New()

	serviceName := c.Args().First()

	service, err := services.Resolve(serviceName)

	if err != nil {
		return err
	}

	tag := docker.ResolveTag(serviceName)

	parameters := map[string]string{
		"container_name": docker.ResolveContainerName(service.Name, tag),
		"port":           service.Port,
		"organization":   service.Organization,
		"image_name":     service.Image,
		"tag":            tag,
	}

	loadServiceParameters(service, parameters)

	docker.EnsureImageIsDownloaded(service.Organization, service.Image, tag)

	fmt.Printf("Enabling %s\n", service.Name)

	err = docker.BootContainer(service.Template, parameters)

	if err != nil {
		return err
	}

	fmt.Printf("%s arrived at the dock and is now ready.\n", service.Name)
	return nil
}

func loadServiceParameters(service *services.Service, parameters map[string]string) error {
	reader := bufio.NewReader(os.Stdin)
	for _, parameter := range service.Parameters {
		defaultValue := ""

		if parameter.DefaultValue != "" {
			defaultValue = fmt.Sprintf(" (default: %s)", parameter.DefaultValue)
		}

		fmt.Printf(parameter.Question+"%s: ", defaultValue)
		text, err := reader.ReadString('\n')

		if err != nil {
			return err
		}

		text = strings.Replace(text, "\n", "", -1)
		text = strings.TrimSpace(text)

		if text == "" {
			text = parameter.DefaultValue
		}

		parameters[parameter.Key] = text
	}

	return nil
}
