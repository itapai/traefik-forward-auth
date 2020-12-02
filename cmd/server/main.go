package main

import (
	"os"

	cmd "github.com/itapai/traefik-forward-auth/pkg/commands"
)

func main() {
	command := cmd.NewForwardAuth()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
