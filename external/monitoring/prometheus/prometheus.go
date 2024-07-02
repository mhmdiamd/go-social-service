package monitoring

import "github.com/prometheus/client_golang/prometheus"

type PromMetrics struct {
	RequestCounter  *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
}

func NewPromMetrics() *PromMetrics {
	return &PromMetrics{}
}

func (p *PromMetrics) SetRequestCounter() {
	p.RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path", "status"},
	)
}

func (p *PromMetrics) SetRequestDuration() {
	p.RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
}
