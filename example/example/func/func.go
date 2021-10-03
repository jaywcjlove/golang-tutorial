package main
import "fmt"

type functinTyoe func(int, int)// // 声明了一个函数类型

func (f functinTyoe)Serve() {
  fmt.Println("serve2")
}

func serve(int,int) {
  fmt.Println("serve1")
}

func main() {
  a := functinTyoe(serve)
  a(1,2)
  a.Serve()
}