package main

import (
	"fmt"
)

func main(){
	a,b := 10,20
	swap(a,b)
	fmt.Println(a,b)
	swap2(&a,&b)
	fmt.Println(a,b)
}

// 值传递
func  swap(a,b int)  {
	temp := a 
	a = b
	b = temp
}

// 引用传递
func swap2(a,b *int)  {
	temp := *a  
	*a = *b
	*b = temp
}

