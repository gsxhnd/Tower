package di

import (
	"github.com/gsxhnd/go-api-template/service_a/routes"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	r routes.Router
}

func NewApplication(r routes.Router) *Application {
	return &Application{
		r: r,
	}
}

func (a *Application) Run() error {
	var g errgroup.Group

	g.Go(func() error {
		return a.r.Run()
	})

	if err := g.Wait(); err != nil {
		return err
	} else {
		return nil
	}
}
