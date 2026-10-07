package main

import (
	"encoding/base64"
	"fmt"
)

type Storer interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string)
	Keys() []string
}

func SetKeyWithEncryption(store Storer, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}

	storedValue, err := store.Get(key)
	if err != nil {
		return "", err
	}

	return storedValue, nil
}

func main() {
	// keyValueStore := NewKeyValueStore(2)

	// _ = s.Set("a", "42")
	// _ = s.Set("b", "43")

	// if err := s.Set("b", "12"); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// val, err := s.Get("d")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(val)

	// encrypted, err := SetKeyWithEncryption(keyValueStore, "e", "hey there")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(encrypted)

	// ttlStore := NewTTLStore(time.Second * 2)

	// encryptedTTLValue, err := SetKeyWithEncryption(ttlStore, "e", "hey there")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(encryptedTTLValue)

	plain := NewKeyValueStore(20)
	logger := NewLogginMidleware(plain)

	encrypted, err := SetKeyWithEncryption(logger, "a", "hey there")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encrypted)
}
