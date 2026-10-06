package main

import (
	"fmt"
	"sort"
	"time"
)

type TTLStore struct {
	data map[string]ttlEntry
	ttl time.Duration
}

type ttlEntry struct {
	value string
	expiresAt time.Time
}

func NewTTLStore(ttl time.Duration) *TTLStore{
	return &TTLStore{
		data: make(map[string]ttlEntry),
		ttl: ttl,
	}
}

func (t *TTLStore) Set(key string, value string) error {
	if key == "" {
		return ErrEmptyKey
	}

	t.data[key] = ttlEntry{
		value: value,
		expiresAt: time.Now().Add(t.ttl),
	}

	return nil
}

func (t *TTLStore) Get(key string) (string, error){
	if key == "" {
		return "", ErrEmptyKey
	}

	entry, ok := t.data[key]

	if !ok || time.Now().After(entry.expiresAt){
		delete(t.data, key)
		return "", fmt.Errorf("key %s doesnot exists", key)
	}

	return entry.value, nil
}

func (t *TTLStore) Delete(key string) {
	delete(t.data, key)
}

func (t *TTLStore) Keys() []string {
	keys := make([]string, 0, len(t.data))
	for key := range t.data {
		if !time.Now().After(t.data[key].expiresAt) {
			keys = append(keys, key)
		} else {
			delete(t.data, key)
		}
	}

	sort.Strings(keys)
	return keys
}


