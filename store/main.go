package main

import "fmt"

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, bool){
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Set(key string, value string) {
	s.data[key] = value
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func NewStore() *Store {
	return  &Store{
		data: make(map[string]string),
	}
}

func main() {
	s := NewStore()
	s.Set("a", "42")
	s.Set("b", "43")

	value,_ := s.Get("a")
	fmt.Println(value)

	s.Delete("a")
	_,exists := s.Get("a")

	fmt.Println(exists)
}