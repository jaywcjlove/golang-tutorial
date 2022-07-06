package main
import (
  "bytes"
  "fmt"
)
func main() {
  // 这里不能写成 b := []byte{"Golang"}，这里是利用类型转换。
  b := []byte("Golang")
  subslice1 := []byte("go")
  subslice2 := []byte("Go")
  // func Contains(b, subslice [] byte) bool
  // 检查字节切片b ，是否包含子字节切片 subslice
  fmt.Println(bytes.Contains(b, subslice1))
  fmt.Println(bytes.Contains(b, subslice2))


  s2 := []byte("同学们，上午好")
  m := func(r rune) rune {
    if r == '上' {
      r = '下'
    }
    return r
  }
  fmt.Println(string(s2))
  // func Map(mapping func(r rune) rune, s []byte) []byte
  // Map函数: 首先将 s 转化为 UTF-8编码的字符序列，
  // 然后使用 mapping 将每个Unicode字符映射为对应的字符，
  // 最后将结果保存在一个新的字节切片中。
  fmt.Println(string(bytes.Map(m, s2)))


  s3 := []byte("google")
  old := []byte("o")
  //这里 new 是一个字节切片，不是关键字了
  new := []byte("oo")
  n := 1
  // func Replace(s, old, new []byte, n int) []byte
  //返回字节切片 S 的一个副本， 并且将前n个不重叠的子切片 old 替换为 new，如果n < 0 那么不限制替换的数量
  fmt.Println(string(bytes.Replace(s3, old, new, n)))
  fmt.Println(string(bytes.Replace(s3, old, new, -1)))


  // 将字节切片 转化为对应的 UTF-8编码的字节序列，并且返回对应的 Unicode 切片。
  s4 := []byte("中华人民共和国")
  r1 := bytes.Runes(s4)
  // func Runes(b []byte) []rune
  fmt.Println(string(s4), len(s4))  // 字节切片的长度
  fmt.Println(string(r1), len(r1))  // rune 切片的长度


  // 字节切片 的每个元素，依旧是字节切片。
  s5 := [][]byte{
    []byte("你好"),
    []byte("世界"),  //这里的逗号，必不可少
  }
  sep := []byte(",")
  // func Join(s [][]byte, sep []byte) []byte
  // 用字节切片 sep 吧 s中的每个字节切片连接成一个，并且返回.
  fmt.Println(string(bytes.Join(s5, sep)))
}