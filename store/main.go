package main

import (
	"fmt"
)

func main() {
	s := NewStore(2)

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

	encrypted, err := s.SetKeyWithEncryption("e", "hey there")
	if err != nil{
		fmt.Println(err)
		return
 	}

	fmt.Println(encrypted)
}
