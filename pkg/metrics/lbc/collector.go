package lbc

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricCollector interface {
	// ObservePodReadinessGateReady this metric is useful to determine how fast pods are becoming ready in the load balancer.
	// Due to some architectural constraints, we can only emit this metric for pods that are using readiness gates.
	ObservePodReadinessGateReady(namespace string, tgbName string, duration time.Duration)
	ObserveControllerReconcileError(controller string, category string, subCategory string)
}

type Collector struct {
	instruments *instruments
}

type noOpCollector struct{}

func (n *noOpCollector) ObservePodReadinessGateReady(_ string, _ string, _ time.Duration) {
}

func NewCollector(registerer prometheus.Registerer) *Collector {
	instruments := newInstruments(registerer)
	return &Collector{
		instruments: instruments,
	}
}

func (c *Collector) ObservePodReadinessGateReady(namespace string, tgbName string, duration time.Duration) {
	c.instruments.podReadinessFlipSeconds.With(prometheus.Labels{
		labelNamespace: namespace,
		labelName:      tgbName,
	}).Observe(duration.Seconds())
}

func (c *Collector) ObserveControllerReconcileError(controller string, category string, subCategory string) {
	c.instruments.controllerReconcileErrors.With(prometheus.Labels{
		labelController:       controller,
		labelErrorCategory:    category,
		labelErrorSubCategory: subCategory,
	}).Inc()
}
