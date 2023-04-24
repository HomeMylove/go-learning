package main

import "fmt"

// 一次声明多个变量
func main() {
	// 方式一
	var n1, n2, n3 int
	fmt.Println("n1=", n1, " n2=", n2, " n3=", n3) // n1 = 0 n2 = 0 n3 = 0

	// 方式二
	var a, b, c = 100, "Yae", 3.14
	fmt.Println("a=", a, " b=", b, " c=", c)

	// 方式三
	one, two, three := "one", 2, "三"
	fmt.Println("one=", one, " two=", two, " three=", three)
}
