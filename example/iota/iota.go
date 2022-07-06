package main
import "fmt"
func main() {
  const (
    // 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
    // 所以 a=0, b=1, c=2 可以简写为如下形式：
    a = iota   //0
    b          //1
    c          //2
    d = "ha"   //独立值，iota += 1
    e          //"ha"   iota += 1
    f = 100    //iota +=1
    g          //100  iota +=1
    h = iota   //7,恢复计数
    i          //8
  )
  fmt.Println(a,b,c,d,e,f,g,h,i)
}