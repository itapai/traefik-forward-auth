package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/itapai/traefik-forward-auth/pkg/server"
)

func NewForwardAuth() *cobra.Command {
	c := &server.Config{}
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "Run the traefik forward auth proxy",
		Run: func(cmd *cobra.Command, _ []string) {
			server.Run(c)
		},
	}

	cmd.Flags().StringVarP(&c.Addr, "addr", "a", ":4181", "addr")
	cmd.Flags().StringVar(&c.JWKS, "jwks", os.Getenv("JWKS_URI"), "jwks url (required)")

	// TODO
	// if jwks empty, throw error

	return cmd

}
