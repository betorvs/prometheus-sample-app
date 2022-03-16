package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path", "status"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path", "status"})

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		rw := newResponseWriter(w)
		next.ServeHTTP(rw, r)

		statusCode := rw.statusCode

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path, fmt.Sprint(statusCode)))

		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
		totalRequests.WithLabelValues(path, fmt.Sprint(statusCode)).Inc()

		timer.ObserveDuration()
	})
}

func init() {
	prometheus.Register(totalRequests)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)
}

func version(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "My application version is "+os.Getenv("VERSION"))
}

func ping(w http.ResponseWriter, req *http.Request) {
	// n is a random number
	n := rand.Intn(10)
	// t is the number used to multiple in time.Sleep causes
	// keep it as is to avoid
	t := n
	if n == 0 {
		t = 1
	}
	if os.Getenv("DEBUG") == "true" {
		log.Printf("Random number %v\n", n)
		log.Printf("Used number %v\n", t)
	}
	switch os.Getenv("FEATURE_FLAG") {
	case "instable":
		time.Sleep(time.Duration(t*1000) * time.Millisecond)
		if n >= 0 && n <= 3 {
			w.WriteHeader(500)
		}
		fmt.Fprintln(w, "Pong")

	case "broken":
		time.Sleep(time.Duration(t*10) * time.Millisecond)
		if n >= 0 && n <= 8 {
			w.WriteHeader(500)
		}
		fmt.Fprintln(w, "Pong")

	case "quick":
		// it uses n to allow some very quick answers
		time.Sleep(time.Duration(n*10) * time.Millisecond)
		// only one change to return an error here
		if n == 8 {
			w.WriteHeader(500)
		}
		fmt.Fprintln(w, "Pong")

	default:
		// with average response time higher and without errors
		time.Sleep(time.Duration(t*100) * time.Millisecond)
		fmt.Fprintln(w, "Pong")

	}
}

func main() {
	router := mux.NewRouter()
	router.Use(prometheusMiddleware)

	// Prometheus endpoint
	router.Path("/metrics").Handler(promhttp.Handler())

	// Prometheus endpoint
	router.HandleFunc("/version", version)

	// Printing variables
	if os.Getenv("FEATURE_FLAG") != "" {
		log.Printf("FEATURE_FLAG=%s\n", os.Getenv("FEATURE_FLAG"))
	}

	// Ping endpoint
	router.HandleFunc("/ping", ping)

	// Serving static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Serving requests on port 8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
