package middleware

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Logging struct {
	l  *logrus.Logger
}

func NewLogging(l *logrus.Logger) *Logging {
	return &Logging{l: l}
}

func (m *Logging) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request)  {
		m.l.Info(fmt.Sprintf("%s : %s",r.Method,r.URL.Path))
		next.ServeHTTP(w,r)
	})
}

