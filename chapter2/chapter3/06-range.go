package main

import (
	"fmt"
)

func main() {
	// 构造一个map
	countryCapitalMap := make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Roma"
	countryCapitalMap["China"] = "Beijing"
	countryCapitalMap["India"] = "New Delhi"

	for k, v := range countryCapitalMap {
		fmt.Println(k, "'s capital is", v)
	}

	// range也可以遍历数组
	five := [5]int{1, 2, 3, 4, 5}
	for i, v := range five {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	// _表示的站位
	for _, v := range five {
		fmt.Printf(" %d\n", v)
	}
}
