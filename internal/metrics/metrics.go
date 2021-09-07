package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type AccountMetrics interface {
	IncCreatedCounter()
	IncreaseCreatedCounter(int)
	IncUpdatedCounter()
	IncRemovedCounter()
}

func RegisterMetrics() AccountMetrics {
	return &account{
		created: promauto.NewCounter(prometheus.CounterOpts{
			Name:      "account_created_count",
			Help:      "Total number of created accounts",
		}),
		updated: promauto.NewCounter(prometheus.CounterOpts{
			Name:      "account_updated_count",
			Help:      "Total number of updated accounts",
		}),
		removed: promauto.NewCounter(prometheus.CounterOpts{
			Name:      "account_removed_count",
			Help:      "Total number of removed accounts",
		}),
	}
}

type account struct {
	created prometheus.Counter
	updated prometheus.Counter
	removed prometheus.Counter
}

func (a *account) IncCreatedCounter() {
	a.created.Inc()
}

func (a *account) IncreaseCreatedCounter(quantity int) {
	a.created.Add(float64(quantity))
}

func (a *account) IncUpdatedCounter() {
	a.updated.Inc()
}

func (a *account) IncRemovedCounter() {
	a.removed.Inc()
}
