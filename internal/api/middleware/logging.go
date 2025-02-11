package middleware

import (
	"log"
	"net/http"
	"time"
)

type respWriterWithStatus struct {
	http.ResponseWriter
	statusCode int
}

func (rw *respWriterWithStatus) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := &respWriterWithStatus{ResponseWriter: w, statusCode: http.StatusOK}
		start := time.Now()
		next.ServeHTTP(wrappedWriter, r)
		log.Printf("[%s] %s | status: %d | time: %vms", r.Method, r.URL.Path, wrappedWriter.statusCode, time.Since(start).Milliseconds())
	})
}
