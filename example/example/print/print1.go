package main
import "fmt"
import "os"

type point struct {
  x, y int
}

// Go by Example: String Formatting
// https://gobyexample.com/string-formatting
func main() {

  // 旗标、宽度、精度、索引
  fmt.Printf("\n旗标、宽度、精度、索引\n----\n")
  // fmt.Printf("|%0+- #[1]*.[2]*[3]d|%0+- #[1]*.[2]*[4]d|\n", 8, 4, 32, 64)
  fmt.Printf("|%v|\n", 8)

  // 格式化排版  
  fmt.Printf("\n格式化排版\n----\n")
  fmt.Printf("|%6d|%9d|\n", 12, 345)
  fmt.Printf("|%6.2f|%9.2f|\n", 1.2, 3.45)
  fmt.Printf("|%-6.2f|%-9.2f|\n", 1.2, 3.450999)

  // 浮点型精度
  fmt.Printf("\n浮点型精度\n----\n")
  fmt.Printf("|%f|%8.4f|%8.f|%.4f|%.f|\n", 3.2, 3.2, 3.2, 3.2, 3.2)
  fmt.Printf("|%.3f|%.3g|\n", 12.345678, 12.345678)
  fmt.Printf("|%.2f|\n", 12.345678+12.345678i)

  // 字符串精度
  fmt.Printf("\n字符串精度\n----\n")
  s := "你好世界！"
  fmt.Printf("|%29s|%8.2s|%8.s|%.2s|%.s|\n", s, s, s, s, s)
  fmt.Printf("|%29x|%8.2x|%8.x|%.2x|%.x|\n", s, s, s, s, s)

  // 带引号字符串
  fmt.Printf("\n带引号字符串\n----\n")
  s1 := "Hello 世界!"       // CanBackquote
  s2 := "Hello\n世界!"      // !CanBackquote
  fmt.Printf("%q\n", s1)  // 双引号
  fmt.Printf("%#q\n", s1) // 反引号成功
  fmt.Printf("%#q\n", s2) // 反引号失败
  fmt.Printf("%+q\n", s2) // 仅包含 ASCII 字符

  // Unicode 码点
  fmt.Printf("\nUnicode 码点\n----\n")
  fmt.Printf("%U, %#U\n", '好', '好')
  fmt.Printf("%U, %#U\n", '\n', '\n')

  // 接口类型将输出其内部包含的值
  fmt.Printf("\n接口类型将输出其内部包含的值\n----\n")
  var i interface{} = struct {
    name string
    age  int
  }{"AAA", 20}
  fmt.Printf("%v\n", i)  // 只输出字段值
  fmt.Printf("%+v\n", i) // 同时输出字段名
  fmt.Printf("%#v\n", i) // Go 语法格式

  // 结构类型将输出其内部包含的值
  fmt.Printf("\n结构类型将输出其内部包含的值\n----\n")
  p := point{1, 2}
  fmt.Printf("%v\n", p)
  fmt.Printf("%+v\n", p)
  fmt.Printf("%#v\n", p)

  // 输出类型
  fmt.Printf("\n输出类型\n----\n")
  fmt.Printf("%T\n", i)
  fmt.Fprintf(os.Stderr, "an %s\n", "errorss")
}