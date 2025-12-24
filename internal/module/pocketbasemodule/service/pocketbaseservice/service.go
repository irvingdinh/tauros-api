package pocketbaseservice

import (
	"context"
	"log"
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"go.uber.org/fx"
)

type PocketbaseService interface{}

func NewPocketbaseService(lc fx.Lifecycle) PocketbaseService {
	i := &pocketbaseServiceImpl{
		app: pocketbase.New(),
	}

	i.app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		i.server = e.Server
		return e.Next()
	})

	lc.Append(fx.Hook{
		OnStart: i.Start,
		OnStop:  i.Stop,
	})

	return i
}

type pocketbaseServiceImpl struct {
	app    *pocketbase.PocketBase
	server *http.Server
}

func (i *pocketbaseServiceImpl) Start(ctx context.Context) error {
	go func() {
		err := i.app.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (i *pocketbaseServiceImpl) Stop(ctx context.Context) error {
	if i.server != nil {
		return i.server.Shutdown(ctx)
	}
	return nil
}
