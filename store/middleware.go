package main

import (
	"log"
	"os"
)

type LoggingMiddleware struct {
	inner Storer
	logger *log.Logger
}

func NewLogginMidleware(inner Storer) *LoggingMiddleware {
	return &LoggingMiddleware{
		inner: inner,
		logger: log.New(os.Stdout, "[log]", log.Ltime|log.Lmicroseconds),
	}
}

func (l *LoggingMiddleware) Get(key string) (string, error) {
	l.logger.Printf("GET %q", key)
	val, err := l.inner.Get(key)
	if err != nil {
		l.logger.Printf("GET %q -> miss (%v)", key, err)
	} else {
		l.logger.Printf("GET %q -> hit", key)
	}
	return val, err
}

func (l *LoggingMiddleware) Set(key string, value string) error {
	l.logger.Printf("SET %q", key)
	return l.inner.Set(key, value)
}


func (l *LoggingMiddleware) Delete(key string) {
	l.logger.Printf("DELETE %q", key)
	l.inner.Delete(key)
}

func (l *LoggingMiddleware) Keys() []string {
	got := l.inner.Keys()
	l.logger.Printf("LEN %q", got)
	return got
}