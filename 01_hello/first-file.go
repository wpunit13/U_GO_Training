package main

import "fmt"

// func wrapper() func() int {
// 	x := 0
// 	fmt.Println("memory address of x is ", &x)
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// func main() {
// 	increment := wrapper()
// 	increment2 := wrapper()
// 	//fmt.Println(increment())
// 	//fmt.Println(increment2())
// 	//fmt.Println(increment2())
// }

func main() {
	for i := 0; i <= 51; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	fmt.Println('d' - 'a' + 10)
	fmt.Println()
}
