package main

import (
	"encoding/base64"
	"errors"
	"fmt"
)

type Store struct {
	data    map[string]string
	maxSize int
}

var ErrEmptyKey = errors.New("No key provided")
var ErrStoreFull = errors.New("Store is full")

func NewStore(maxSize int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
}

func (s *Store) SetKeyWithEncryption(key, val string) (string, error){
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := s.Set(key, encoded); err != nil{
		return "", err
	}

	return s.Get(key)
}

func (s *Store) Get(key string) (string, error) {
	val, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("key %s is not present", key)
	}
	return val, nil
}

func (s *Store) Set(key string, value string) error {
	if key == "" {
		return ErrEmptyKey
	}

	_, exists := s.data[key]

	if s.maxSize > 0 && len(s.data) >= s.maxSize && !exists {
		return fmt.Errorf("Set (%q) : %w", key, ErrStoreFull)
	}

	s.data[key] = value

	return nil
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

