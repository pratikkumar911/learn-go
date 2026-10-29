package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 11, 7, 3}

	sort.Ints(arr);

	left := 0
	right := len(arr)-1

	for left < right{
		if arr[left] + arr[right] == 9{
			fmt.Println(left, right)
		} else if arr[left] + arr[right] > 9{
			right--;
		} else {
			left++
		}
	}
}