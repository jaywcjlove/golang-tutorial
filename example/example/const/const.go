package main
import "unsafe"

// 常量可以用len(), cap(), unsafe.Sizeof()常量计算表达式的值。
// 常量表达式中，函数必须是内置函数，否则编译不过：
const (
  a = "abc"
  b = len(a)
  c = unsafe.Sizeof(a)
)

func main(){
  const (
    PI     = 3.14
    const1 = "1"
  )
  const LENGTH int = 10
  const e, f, g = 1, false, "str" //多重赋值
  println(a, b, c,PI, LENGTH)
}