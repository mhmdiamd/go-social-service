package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/mhmdiamd/go-social-service/domain/auth"
	categoryComunity "github.com/mhmdiamd/go-social-service/domain/category-community"
	"github.com/mhmdiamd/go-social-service/domain/community"
	communityMember "github.com/mhmdiamd/go-social-service/domain/community_member"
	eventDemographics "github.com/mhmdiamd/go-social-service/domain/event-demographics"
	"github.com/mhmdiamd/go-social-service/external/database"
	"github.com/mhmdiamd/go-social-service/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path", "status"},
	)

	tesCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_test",
			Help: "Total number of HTTP requests 123.",
		},
		[]string{"method", "path", "status"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(tesCounter)
	prometheus.MustRegister(requestDuration)
	// prometheus.MustRegister(prometheusMatrics.RequestDuration)
}

func main() {
	filename := "./cmd/api/config.yaml"

	if os.Getenv("APP_ENV") == "staging" {
		filename = "./cmd/api/config-staging.yaml"
	}

	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.Db)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Print("db connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})

	router.Use(func(c *fiber.Ctx) error {
		start := time.Now()

		duration := time.Since(start).Seconds()
		path := c.Path()
		method := c.Method()
		status := c.Response().StatusCode()

		// Record metrics
		requestCounter.WithLabelValues(method, path, strconv.Itoa(status)).Inc()
		requestDuration.WithLabelValues(method, path).Observe(duration)
		return c.Next()
	})

	router.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	auth.Init(router, db)
	categoryComunity.Init(router, db)
	eventDemographics.Init(router, db)
	community.Init(router, db)
	communityMember.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
