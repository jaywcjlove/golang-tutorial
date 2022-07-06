package main
import (
  "fmt"
  "math"
)

/* 定义一个 interface */
type shape interface {
  area() float64
}

/* 定义一个 circle */
type circle struct {
  x,y,radius float64
}

/* 定义一个 rectangle */
type rectangle struct {
  width, height float64
}

/* 定义一个circle方法 (实现 shape.area())*/
func(circle circle) area() float64 {
  return math.Pi * circle.radius * circle.radius
}

/* 定义一个rectangle方法 (实现 shape.area())*/
func(rect rectangle) area() float64 {
  return rect.width * rect.height
}

/* 定义一个shape的方法*/
func getArea(shape shape) float64 {
  return shape.area()
}

func main() {
  circle := circle{x:0,y:0,radius:5}
  rectangle := rectangle {width:10, height:5}

  fmt.Printf("circle area: %f\n",getArea(circle))
  fmt.Printf("rectangle area: %f\n",getArea(rectangle))
}
  