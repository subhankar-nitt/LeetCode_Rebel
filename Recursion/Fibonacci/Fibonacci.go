package main

import "fmt"

func main(){
	fmt.Println(fibonacci(5))
}

func fibonacci(n int)int{
	if n<2{
		return n
	}
	return fibonacci(n-1)+fibonacci(n-2)
}
