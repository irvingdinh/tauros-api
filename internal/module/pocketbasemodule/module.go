package pocketbasemodule

import (
	"go.uber.org/fx"

	"github.com/irvingdinh/tauros-api/internal/module/pocketbasemodule/service/pocketbaseservice"
)

func NewPocketbaseModule() fx.Option {
	return fx.Module("pocketbase",
		fx.Provide(pocketbaseservice.NewPocketbaseService),
		fx.Invoke(func(pocketbaseservice.PocketbaseService) {}),
	)
}
