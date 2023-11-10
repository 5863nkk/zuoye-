package main

import "fmt"

type jisuan func(int, int) int

func jiafa(a, b int) int {
	return a + b
}
func jianfa(a, b int) int {
	return a - b
}
func chenfa(a, b int) int {
	return a * b
}
func chufa(a, b int) int {
	if b == 0 {
		fmt.Println("除数不能为0")
		return 0
	}
	return a / b
}
func calc(a, b int, f jisuan) (result int) {
	result = f(a, b)
	return
}
func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	result := calc(a, b, jiafa)
	fmt.Println("加法结果:", result)
	result = calc(a, b, jianfa)
	fmt.Println("减法结果:", result)
	result = calc(a, b, chenfa)
	fmt.Println("乘法结果:", result)
	result = calc(a, b, chufa)
	fmt.Println("除法结果:", result)
}
