package main

import (
	"log"
	"os"
	"renatofmachado/dock/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "enable",
				Aliases: []string{"e"},
				Usage:   "Enable a service",
				Action:  commands.Enable,
			},
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "Start a service",
				Action:  commands.Start,
			},
			{
				Name:   "stop",
				Usage:  "Stop a service",
				Action: commands.Stop,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
