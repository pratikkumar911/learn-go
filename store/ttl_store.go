package main

import (
	"fmt"
	"time"
)

type TTLStore struct {
	data map[string]ttlEntry
}

type ttlEntry struct {
	value string
	expiresAt time.Time
}