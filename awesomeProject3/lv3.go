package main

import "fmt"

func anm(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		} else {
			return true
		}
	}
	return true
}

func main() {
	ans := anm(19)
	if ans {
		fmt.Print("是素数")
	} else {
		fmt.Print("不是素数")
	}
}
