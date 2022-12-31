package main

import (
	"fmt"
)

const(
	login = iota 			// 0
	logout					// 1
	user = iota + 2			// 2+2
	account = iota * 3		// 3*3
)

const(
	apple,banana  = iota + 1,iota + 2		// 0+1 0+2
	peach,pear								// 1+1 1+2 
	orange,mango							// 2+1 2+2
)

func main(){
	fmt.Println(login,logout,user,account)
	fmt.Println(apple,banana,peach,pear,orange,mango)
}

