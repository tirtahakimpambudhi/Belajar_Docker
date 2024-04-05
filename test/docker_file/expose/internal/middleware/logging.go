package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Logging struct {
	l  *slog.Logger
}

func (m *Logging) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request)  {
		m.l.Info(fmt.Sprintf("%s : %s",r.Method,r.URL.Path))
		next.ServeHTTP(w,r)
	})
}

