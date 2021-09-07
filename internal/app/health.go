package app

import "github.com/ozonva/ova-account-api/internal/health"

type Healther interface {
	Health() error
}

func (a *App) Check() []health.Resource {
	var resources []health.Resource

	db := health.Resource{Name: "Store", Status: health.Ok}
	if err := a.Store.Health(); err != nil {
		db.Status = health.Error
		db.Message = err.Error()
	}

	kafka := health.Resource{Name: "Kafka", Status: health.Ok}
	if err := a.Producer.Health(); err != nil {
		db.Status = health.Error
		db.Message = err.Error()
	}

	return append(resources, db, kafka)
}
