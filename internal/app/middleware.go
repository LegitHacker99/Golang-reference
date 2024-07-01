package app

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.Handler) http.Handler

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *ResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func CreateMiddlewareStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}

		return next
	}
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &ResponseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		log.Printf("[%s] | %s %s | %d | %v",
			time.Now().Format(time.RFC3339), r.Method, r.URL.Path, rw.statusCode, time.Since(start))
		// log.Printf(" | %s %s | %d | %v", r.Method, r.URL.Path, rw.statusCode, time.Since(start))
	})
}
