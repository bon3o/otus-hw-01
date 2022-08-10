package internalhttp

import (
	"fmt"
	"net/http"

	"github.com/bon3o/otus-hw-01/hw12_13_14_15_calendar/internal/app"
)

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode  int
	BytesLength int
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	n, err := w.ResponseWriter.Write(data)
	w.BytesLength += n

	return n, err
}

func loggingMiddleware(next http.Handler, log app.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		myWriter := &ResponseWriter{w, 0, 0}
		next.ServeHTTP(myWriter, r)
		log.Info(fmt.Sprintf("request from %s, method: %s. StatusCode: %d", r.RemoteAddr, r.Method, myWriter.StatusCode))
	})
}
