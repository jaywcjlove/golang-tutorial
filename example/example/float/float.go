package main
import "fmt"

func main() {
  var x float64
  x = 20.0
  fmt.Println(x)
  fmt.Printf("x is of type %T\n", x)

  a := float64(20.0)
  b := 42 
  fmt.Println(a)
  fmt.Println(b)
  fmt.Printf("a is of type %T\n", a)
  fmt.Printf("b is of type %T\n", b)
}