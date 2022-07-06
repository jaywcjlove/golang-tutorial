package main
import "fmt"
func main() {
  // 声明一个长度为5的整数数组
  // 一旦数组被声明了，那么它的数据类型跟长度都不能再被改变。
  var array1 [5]int

  fmt.Printf("array1: %d\n\n", array1)

  // 声明一个长度为5的整数数组
  // 初始化每个元素
  array2 := [5]int{12, 123, 1234, 12345, 123456}
  array2[1] = 5000
  fmt.Printf("array2: %d\n\n", array2[1])

  // n 是一个长度为 10 的数组
  var n [10]int 
  var i,j int

  /* 为数组 n 初始化元素 */         
  for i = 0; i < 10; i++ {
    n[i] = i + 100 /* 设置元素为 i + 100 */
  }

  /* 输出每个数组元素的值 */
  for j = 0; j < 10; j++ {
    fmt.Printf("Element[%d] = %d\n", j, n[j] )
  }

  /* 数组 - 5 行 2 列*/
  var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
  var e, f int

  /* 输出数组元素 */
  for  e = 0; e < 5; e++ {
    for f = 0; f < 2; f++ {
        fmt.Printf("a[%d][%d] = %d\n", e,f, a[e][f] )
    }
  }
}