package lbc

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricSubsystem = "awslbc"
)

// These metrics are exported to be used in unit test validation.
const (
	// MetricPodReadinessGateReady tracks the time to flip a readiness gate to true
	MetricPodReadinessGateReady = "readiness_gate_ready_seconds"
	// MetricControllerReconcileErrors tracks the number of controller errors
	MetricControllerReconcileErrors = "controller_reconcile_errors_total"
)

const (
	labelNamespace = "namespace"
	labelName      = "name"

	labelController       = "controller"
	labelErrorCategory    = "category"
	labelErrorSubCategory = "sub_category"
)

type instruments struct {
	podReadinessFlipSeconds   *prometheus.HistogramVec
	controllerReconcileErrors *prometheus.CounterVec
}

// newInstruments allocates and register new metrics to registerer
func newInstruments(registerer prometheus.Registerer) *instruments {
	podReadinessFlipSeconds := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: metricSubsystem,
		Name:      MetricPodReadinessGateReady,
		Help:      "Latency from pod getting added to the load balancer until the readiness gate is flipped to healthy.",
		Buckets:   []float64{10, 30, 60, 120, 180, 240, 300, 360, 420, 480, 540, 600},
	}, []string{labelNamespace, labelName})

	controllerReconcileErrors := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: metricSubsystem,
		Name:      MetricControllerReconcileErrors,
		Help:      "Counts the number of reconcile error, categorized by error type.",
	}, []string{labelController, labelErrorCategory, labelErrorSubCategory})

	registerer.MustRegister(podReadinessFlipSeconds, controllerReconcileErrors)
	return &instruments{
		podReadinessFlipSeconds:   podReadinessFlipSeconds,
		controllerReconcileErrors: controllerReconcileErrors,
	}
}
