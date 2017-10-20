package main
import "fmt"
func main() {
	// 2. 在输出 world
	defer fmt.Println("world")
	// 1. 先输出 hello
	fmt.Println("hello")
	// 输出结果
	// =======
	// hello
	// world


	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
	// 输出结果
	// =======
	// counting
	// done
	// 9
	// 8
	// 7
	// 6
	// 5
	// 4
	// 3
	// 2
	// 1
	// 0
}