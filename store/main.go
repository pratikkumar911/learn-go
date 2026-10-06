package main

import (
	"fmt"
	"time"
	"encoding/base64"
)

func SetKeyWithEncryption(store Store, key, val string) (string, error){
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := store.Set(key, encoded); err != nil{
		return "", err
	}

	return store.Get(key)
}

func main() {
	keyValueStore := NewStore(2)

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

	encrypted, err := SetKeyWithEncryption(*keyValueStore, "e", "hey there")
	if err != nil{
		fmt.Println(err)
		return
 	}

	fmt.Println(encrypted)

	ttlStore := NewTTLStore(time.Second * 2)

	encryptedTTLValue, err := SetKeyWithEncryption(*ttlStore, "e", "hey there")
	if err != nil{
		fmt.Println(err)
		return
 	}

	fmt.Println(encryptedTTLValue)
}
