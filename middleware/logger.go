package middleware

import (
	"net/http"
	"time"

	"github.com/charmbracelet/log"
)

func LoggerMiddleware() func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			si := &statusInterceptor{ResponseWriter: w}
			next.ServeHTTP(si, r)

			end := time.Since(start)

			if si.code == 0 {
				si.code = 200
			}

			log.Infof("%3d %12s - %-12s [%v]", si.code, r.RemoteAddr, r.URL.Path, end)

		})
	}
}

type statusInterceptor struct {
	http.ResponseWriter
	code int
}

func (s *statusInterceptor) WriteHeader(statusCode int) {
	s.code = statusCode
	s.ResponseWriter.WriteHeader(statusCode)
}
