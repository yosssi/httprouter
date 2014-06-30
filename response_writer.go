package httprouter

import "net/http"

// ResponseWriter is an interface for writing a response.
type ResponseWriter interface {
	http.ResponseWriter
	// Status returns the status code of the response.
	Status() int
}

// responseWriter represents an HTTP response.
type responseWriter struct {
	http.ResponseWriter
	status int
}

// Header invokes the response writer's Header method.
func (res *responseWriter) Header() http.Header {
	return res.ResponseWriter.Header()
}

// Write invokes the response writer's Write method.
func (res *responseWriter) Write(b []byte) (int, error) {
	if res.status == 0 {
		res.status = http.StatusOK
	}
	return res.ResponseWriter.Write(b)
}

// WriteHeader invokes the response writer's WriteHeader method.
func (res *responseWriter) WriteHeader(i int) {
	res.status = i
	res.ResponseWriter.WriteHeader(i)
}

// Status returns the response's status.
func (res *responseWriter) Status() int {
	return res.status
}

// newResponseWriter creates and returns a response writer.
func newResponseWriter(w http.ResponseWriter) ResponseWriter {
	return &responseWriter{
		ResponseWriter: w,
	}
}
