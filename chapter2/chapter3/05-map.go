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
	countryCapitalMap["Japan"] = " "

	fmt.Println(countryCapitalMap)

	a := countryCapitalMap["China"]

	fmt.Println("china's capital is", a)
	//取map的时候可以使用指示器
	val, ok := countryCapitalMap["Japan"]
	if ok {
		fmt.Println("Japan's capital is", val)
	} else {
		fmt.Println("Japan's capital not exists")
	}
}
