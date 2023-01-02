package main

import (
	"fmt"
)

func main() {
	a2 := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(a2)

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("i = %d,j = %d,val = %d\n", i, j, a2[i][j])
		}
	}
}
