// request_logger.go
package middleware

import (
	"net/http"
	"runtime/debug"
	"time"
  
	log "github.com/sirupsen/logrus"
  )
  
  // responseWriter is a minimal wrapper for http.ResponseWriter that allows the
  // written HTTP status code to be captured for logging.
  type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
  }
  
  func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
  }
  
  func (rw *responseWriter) Status() int {
	return rw.status
  }
  
  func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
	  return
	}
  
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
  
	return
  }
  
  // LoggingMiddleware logs the incoming HTTP request & its duration.
	func LoggingMiddleware(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Error(
					"err", err,
					"trace", debug.Stack(),
				)
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			log.WithFields(log.Fields{
				"status": wrapped.status,
				"method": r.Method,
				"path": r.URL.EscapedPath(),
				"duration": time.Since(start),
			}).Info("API Logged in successfully")
			
		}
		return http.HandlerFunc(fn)
	}