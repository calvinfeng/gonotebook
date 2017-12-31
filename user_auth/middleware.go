package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type HttpMiddleware func(http.Handler) http.Handler

func NewServerLoggingMiddleware() HttpMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logrus.Infof("%s %s %s %s", r.Proto, r.Method, r.URL, r.Host)
			next.ServeHTTP(w, r)
		})
	}
}
