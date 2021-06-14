package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	createCounter *prometheus.CounterVec
	updateCounter *prometheus.CounterVec
	removeCounter *prometheus.CounterVec
)

func RegisterMetrics() {
	// create a new counter vector
	createCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "create_count",
			Help: "Number of successful created projects.",
		},
		[]string{"quantity"}, // labels
	)
	updateCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "update_count",
			Help: "Number of successful updated projects.",
		},
		[]string{"quantity"}, // labels
	)
	removeCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "remove_count",
			Help: "Number of successful removed projects.",
		},
		[]string{"quantity"}, // labels
	)

	// must register counter on init
	prometheus.MustRegister(createCounter, updateCounter, removeCounter)
}

func CreateCounterInc(status string) {
	if createCounter != nil {
		createCounter.WithLabelValues(status).Inc()
	}
}

func UpdateCounterInc(status string) {
	if updateCounter != nil {
		updateCounter.WithLabelValues(status).Inc()
	}
}

func RemoveCounterInc(operation string) {
	if removeCounter != nil {
		removeCounter.WithLabelValues(operation).Inc()
	}
}
