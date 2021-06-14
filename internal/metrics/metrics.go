package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	createCounter *prometheus.CounterVec
	updateCounter *prometheus.CounterVec
	removeCounter *prometheus.CounterVec
)

func RegisterMetrics() {
	createCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "create_count",
			Help: "Number of successful created projects.",
		},
		[]string{"operation"},
	)
	updateCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "update_count",
			Help: "Number of successful updated projects.",
		},
		[]string{"operation"},
	)
	removeCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "remove_count",
			Help: "Number of successful removed projects.",
		},
		[]string{"operation"},
	)

	// must register counter on init
	prometheus.MustRegister(createCounter, updateCounter, removeCounter)
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
