package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	createCounter *prometheus.CounterVec
	updateCounter *prometheus.CounterVec
	removeCounter *prometheus.CounterVec
)

func RegisterMetrics() {

	createCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "create_count",
		Help: "Number of successful created notes.",
	},
		[]string{"operation"})

	updateCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "update_count",
		Help: "Number of successful updated notes.",
	},
		[]string{"operation"})

	removeCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "remove_count",
		Help: "Number of successful removed notes.",
	},
		[]string{"operation"})
}

func CreateCounterInc(operation string) {
	if createCounter != nil {
		createCounter.With(prometheus.Labels{"operation": operation}).Inc()
	}
}

func UpdateCounterInc(operation string) {
	if updateCounter != nil {
		updateCounter.With(prometheus.Labels{"operation": operation}).Inc()
	}
}

func RemoveCounterInc(operation string) {
	if removeCounter != nil {
		removeCounter.With(prometheus.Labels{"operation": operation}).Inc()
	}
}
