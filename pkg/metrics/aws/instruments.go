package aws

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricSubSystem = "aws"

	metricAPICallsTotal          = "api_calls_total"
	metricAPICallDurationSeconds = "api_call_duration_seconds"
	metricAPICallRetries         = "api_call_retries"

	metricAPICall4xxTotal         = "api_call_4xx_total"
	metricAPICall5xxTotal         = "api_call_5xx_total"
	metricAPIAuthErrorsTotal      = "api_call_auth_errors_total"
	metricAPILimitExceededTotal   = "api_call_limit_exceeded_total"
	metricAPIThrottledTotal       = "api_call_throttled_total"
	metricAPIValidationErrorTotal = "api_call_validation_error_total"

	metricAPIRequestsTotal          = "api_requests_total"
	metricAPIRequestDurationSeconds = "api_request_duration_seconds"
)

const (
	labelService    = "service"
	labelOperation  = "operation"
	labelStatusCode = "status_code"
	labelErrorCode  = "error_code"
)

type instruments struct {
	apiCallsTotal            *prometheus.CounterVec
	apiCallDurationSeconds   *prometheus.HistogramVec
	apiCallRetries           *prometheus.HistogramVec
	apiRequestsTotal         *prometheus.CounterVec
	apiRequestDurationSecond *prometheus.HistogramVec

	apiCall4xxTotal             *prometheus.CounterVec
	apiCall5xxTotal             *prometheus.CounterVec
	apiCallAuthErrorsTotal      *prometheus.CounterVec
	apiCallLimitExceededTotal   *prometheus.CounterVec
	apiCallThrottledTotal       *prometheus.CounterVec
	apiCallValidationErrorTotal *prometheus.CounterVec
}

// newInstruments allocates and register new metrics to registerer
func newInstruments(registerer prometheus.Registerer) *instruments {
	apiCallsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: metricSubSystem,
		Name:      metricAPICallsTotal,
		Help:      "Total number of SDK API calls from the customer's code to AWS services",
	}, []string{labelService, labelOperation, labelStatusCode, labelErrorCode})
	apiCallDurationSeconds := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: metricSubSystem,
		Name:      metricAPICallDurationSeconds,
		Help:      "Perceived latency from when your code makes an SDK call, includes retries",
	}, []string{labelService, labelOperation})
	apiCallRetries := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: metricSubSystem,
		Name:      metricAPICallRetries,
		Help:      "Number of times the SDK retried requests to AWS services for SDK API calls",
		Buckets:   []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}, []string{labelService, labelOperation})

	apiRequestsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: metricSubSystem,
		Name:      metricAPIRequestsTotal,
		Help:      "Total number of HTTP requests that the SDK made",
	}, []string{labelService, labelOperation, labelStatusCode, labelErrorCode})
	apiRequestDurationSecond := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: metricSubSystem,
		Name:      metricAPIRequestDurationSeconds,
		Help:      "Latency of an individual HTTP request to the service endpoint",
	}, []string{labelService, labelOperation})

	apiCall4xxTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: metricSubSystem,
		Name:      metricAPICall4xxTotal,
		Help:      "Number of AWS API calls that resulted in 4xx error",
	}, []string{labelService, labelOperation, labelStatusCode, labelErrorCode})

	apiCall5xxTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: metricSubSystem,
		Name:      metricAPICall5xxTotal,
		Help:      "Number of AWS API calls that resulted in 5xx error",
	}, []string{labelService, labelOperation, labelStatusCode, labelErrorCode})

	apiCallAuthErrorsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: metricSubSystem,
		Name:      metricAPIAuthErrorsTotal,
		Help:      "Number of failed AWS API calls that due to auth or authrorization failures",
	}, []string{labelService, labelOperation, labelStatusCode, labelErrorCode})

	apiCallLimitExceededTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: metricSubSystem,
		Name:      metricAPILimitExceededTotal,
		Help:      "Number of failed AWS API calls that due to exceeding servce limit",
	}, []string{labelService, labelOperation, labelStatusCode, labelErrorCode})

	apiCallThrottledTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: metricSubSystem,
		Name:      metricAPIThrottledTotal,
		Help:      "Number of failed AWS API calls that due to throtting error",
	}, []string{labelService, labelOperation, labelStatusCode, labelErrorCode})

	apiCallValidationErrorTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: metricSubSystem,
		Name:      metricAPIValidationErrorTotal,
		Help:      "Number of failed AWS API calls that due to validation error",
	}, []string{labelService, labelOperation, labelStatusCode, labelErrorCode})

	registerer.MustRegister(apiCallsTotal, apiCallDurationSeconds, apiCallRetries, apiRequestsTotal, apiRequestDurationSecond, apiCall4xxTotal, apiCall5xxTotal, apiCallAuthErrorsTotal, apiCallLimitExceededTotal, apiCallThrottledTotal, apiCallValidationErrorTotal)

	return &instruments{
		apiCallsTotal:               apiCallsTotal,
		apiCallDurationSeconds:      apiCallDurationSeconds,
		apiCallRetries:              apiCallRetries,
		apiRequestsTotal:            apiRequestsTotal,
		apiRequestDurationSecond:    apiRequestDurationSecond,
		apiCall4xxTotal:             apiCall4xxTotal,
		apiCall5xxTotal:             apiCall5xxTotal,
		apiCallAuthErrorsTotal:      apiCallAuthErrorsTotal,
		apiCallLimitExceededTotal:   apiCallLimitExceededTotal,
		apiCallThrottledTotal:       apiCallThrottledTotal,
		apiCallValidationErrorTotal: apiCallValidationErrorTotal,
	}
}
