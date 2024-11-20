package reqstate

import "net/http"

type StatusCodeResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewStatusCodeResponseWriter(w http.ResponseWriter) http.ResponseWriter {
	return &StatusCodeResponseWriter{w, 200}
}

func (rw *StatusCodeResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statusCode = statusCode
}

func GetStatusCode(rw http.ResponseWriter) int {
	return rw.(*StatusCodeResponseWriter).statusCode
}
