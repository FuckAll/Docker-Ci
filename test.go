package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	fmt.Println(len(a))
	fmt.Println(a[1:len(a)])
}
