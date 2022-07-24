package logRequestHandler

import (
    "io"
    "log"
    "net/http"
    "time"
)

// Logger records incoming requests and uses the log package to print http method,
// route url and duration to screen/logs.
func Logger(inner http.Handler, name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        // Save start time to calculate duration
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf(
            "%s\t\t%s\t\t%s\t\t%s",
            r.Method,
            r.RequestURI,
            name,
            time.Since(start),
        )
    })
}
