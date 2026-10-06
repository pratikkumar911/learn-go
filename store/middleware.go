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