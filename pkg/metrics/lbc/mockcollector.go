package lbc

import (
	"time"
)

type MockCollector struct {
	Invocations map[string][]interface{}
}

type MockHistogramMetric struct {
	namespace string
	name      string
	duration  time.Duration
}

type MockCounterMetric struct {
	labelController       string
	labelErrorCategory    string
	labelErrorSubCategory string
}

func (m *MockCollector) ObservePodReadinessGateReady(namespace string, tgbName string, d time.Duration) {
	m.recordHistogram(MetricPodReadinessGateReady, namespace, tgbName, d)
}

func (m *MockCollector) ObserveControllerReconcileError(controller string, category string, subCategory string) {
	m.Invocations[MetricControllerReconcileErrors] = append(m.Invocations[MetricControllerReconcileErrors], MockCounterMetric{
		labelController:       controller,
		labelErrorCategory:    category,
		labelErrorSubCategory: subCategory,
	})
}

func (m *MockCollector) recordHistogram(metricName string, namespace string, name string, d time.Duration) {
	m.Invocations[metricName] = append(m.Invocations[MetricPodReadinessGateReady], MockHistogramMetric{
		namespace: namespace,
		name:      name,
		duration:  d,
	})
}

func NewMockCollector() MetricCollector {

	mockInvocations := make(map[string][]interface{})
	mockInvocations[MetricPodReadinessGateReady] = make([]interface{}, 0)
	mockInvocations[MetricControllerReconcileErrors] = make([]interface{}, 0)

	return &MockCollector{
		Invocations: mockInvocations,
	}
}
