package main

import (
	"log"
	"net/http"
	"time"
)

type LoggerMiddleware struct {
	handler http.Handler
}

func (l *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s, %s, %s, %v", r.RemoteAddr, r.Method,
		r.URL.Path, time.Since(start))
}

func NewLoggerMiddleware(handlerToWrap http.Handler) *LoggerMiddleware {
	return &LoggerMiddleware{handlerToWrap}
}
