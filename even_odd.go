package main

import (
	"fmt"
)

func main() {
	nums := []int{1,2,3,4,5,6,7,8,9,10}

	for i := range nums {
		if nums[i] % 2 == 0 {
			fmt.Println(nums[i], "is even")
		} else {
			fmt.Println(nums[i], "is odd")
		}	
	}	
}