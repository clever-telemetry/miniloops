package runner

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	ExecCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "miniloops",
		Subsystem: "runner",
		Name:      "execution_count",
		Help:      "Execution count for each Loop",
	}, []string{"namespace", "loop"})

	ExecErrorCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "miniloops",
		Subsystem: "runner",
		Name:      "execution_error_count",
		Help:      "Execution count for each Loop",
	}, []string{"namespace", "loop"})

	ExecDuration = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "miniloops",
		Subsystem: "runner",
		Name:      "execution_duration",
		Help:      "Execution duration for each Loop",
		ConstLabels: map[string]string{
			"unit": "ms",
		},
	}, []string{"namespace", "loop"})

	FetchedCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "miniloops",
		Subsystem: "runner",
		Name:      "fetched_datatpoints",
		Help:      "Datapoints fetched by Loop",
	}, []string{"namespace", "loop"})

	OperationsCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "miniloops",
		Subsystem: "runner",
		Name:      "operations_count",
		Help:      "Operations count for each Loop",
	}, []string{"namespace", "loop"})
)

func init() {
	metrics.Registry.MustRegister(
		ExecCount,
		ExecErrorCount,
		ExecDuration,
		FetchedCount,
		OperationsCount,
	)
}
