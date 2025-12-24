package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/irvingdinh/tauros-api/internal/module/pocketbasemodule"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(
			pocketbasemodule.NewPocketbaseModule(),
		).Run()
	},
}
