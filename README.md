Go语言快速入门
===

通过简单的例子，来快速入门Go语言基础编程、语法等各种语言特性，主要面向新手级别的学习者。下面所有例子均来源于网络，看文需谨慎后果自负。

安装之前需要了解及新建几个必要的文件目录：

- GOROOT 目录，该目录为解压压缩包所存放的目录。
    >（建议 linux 环境解压到 /usr/local 目录，windows 环境解压到 C:\ProgramFiles 目录）
- 新建 GOPATH 目录，即为我们的“工作目录”，该目录可以有多个，建议只设置一个。
- GOPATH 目录下新建 `src` 目录，该目录用于存放第三方库源码，以及存放我们的项目的源码。
- GOPATH 目录下新建 `bin` 目录，该目录用于存放项目中所生成的可执行文件。
- GOPATH 目录下新建 `pkg` 目录，该目录用于存放编译生成的库文件。

目录
---

<details>
<summary>点击展开目录菜单</summary>

<!-- TOC -->

- [安装Go](#安装go)
- [运行Go](#运行go)
- [编程基础](#编程基础)
- [基本类型](#基本类型)
- [常量变量](#常量变量)
- [语言运算符](#语言运算符)
- [资源导航](#资源导航)

<!-- /TOC -->

</details>

## 安装Go

<details>
<summary>CentOS7中通过yum安装</summary>

```bash
# CentOS7 可以只用使用yum安装
yum install golang  
```

</details>

<details>
<summary>CentOS7中通过源码安装</summary>

```bash
# 源码下载
# 官网源码 https://golang.org/dl/ 需要翻墙
wget https://storage.googleapis.com/golang/go1.9.darwin-amd64.pkg
tar zxvf go1.8.linux-amd64.tar.gz -C /usr/local
# 新建GOPATH目录
mkdir -p $HOME/gopath
```

</details>

<details>
<summary>Mac中通过brew命令安装</summary>

使用[home brew](http://brew.sh/)安装方便快捷安装Go，如果你想要在你的 Mac 系统上安装 Go，则必须使用 `Intel 64` 位处理器，Go 不支持 `PowerPC` 处理器。

```bash
brew update && brew upgrade # 更新 Homebrew 的信息
brew install git            # 安装 git
brew install go             # 安装 go
```

</details>


<details>
<summary>Mac中通过源码安装</summary>

通过源代码编译安装的过程与环境变量的配置与在 Linux 系统非常相似，因此不再赘述。

**注意事项：**

在 Mac 系统下使用到的 C 工具链是 Xcode 的一部分，因此你需要通过安装 Xcode 来完成这些工具的安装。你并不需要安装完整的 Xcode，而只需要安装它的命令行工具部分。

</details>


<details>
<summary>环境变量配置</summary>

通过`go env`查看go的详细信息

```bash
→ go env

GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin"
GOPATH="/Users/kenny/go"
GORACE=""
GOROOT="/usr/local/Cellar/go/1.9/libexec"
GOTOOLDIR="/usr/local/Cellar/go/1.9/libexec/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/j7/3xly5sk567s65ny5dnr__3b80000gn/T/go-build377856897=/tmp/go-build -gno-record-gcc-switches -fno-common"
CXX="clang++"
CGO_ENABLED="1"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
```

如果需要修改默认的环境变量配置修改 `vim ~/.bash_profile` 或者 `vim ~/.zshrc`

```bash
#GOROOT
# CentOS 中如下设置 GOROOT，看你安装的路径
# export GOROOT=/usr/local/go
# Mac OS 中通过命令行工具brew安装如下配置 GOROOT
export GOROOT=/usr/local/Cellar/go/1.9/libexec
#GOPATH root bin
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
#GOPATH
export GOPATH=$HOME/go
#GOPATH bin
export PATH=$PATH:$GOPATH/bin
```

使其立即生效

```bash
source /etc/profile
```

</details>


<details>
<summary>标准命令详解</summary>

```bash
→ go --help
Go is a tool for managing Go source code.
Go是用于管理Go源代码的工具。

Usage用法:
	go command [arguments]

The commands are:

  build       命令用于编译我们指定的源码文件或代码包以及它们的依赖包。
  clean       删除掉执行其它命令时产生的一些文件和目录。
  doc         命令可以打印附于Go语言程序实体上的文档。
  env         用于打印Go语言的环境信息。
  bug         启动错误报告。
  fix         把指定代码包的所有Go语言源码文件中的旧版本代码修正为新版本的代码。
  fmt         在包源上运行gofmt。
  generate    通过处理源生成Go文件。
  get         下载或更新安装指定的代码包及其依赖包，并对它们进行编译和安装。
  install     用于编译并安装指定的代码包及它们的依赖包。
  list        列出指定的代码包的信息。
  run         命令可以编译并运行命令源码文件。
  test        对Go语言编写的程序进行测试。
  tool        运行指定的go工具
  version     打印Go的版本信息
  vet         用于检查Go语言源码中静态错误的简单工具。

Use "go help [command]" for more information about a command.

Additional help topics:

  c           calling between Go and C
  buildmode   description of build modes
  filetype    file types
  gopath      GOPATH environment variable
  environment environment variables
  importpath  import path syntax
  packages    description of package lists
  testflag    description of testing flags
  testfunc    description of testing functions

Use "go help [topic]" for more information about that topic.
```

其它命令

```bash
cat $GOROOT/VERSION  # 查看版本
$GOROOT/src/all.bash # 测试用例正确
```

</details>

## 运行Go

<details>
<summary>通过go命令运行</summary>

我们先写一段GO代码，很简单就是打印输出一个`hello world!`, 保存为[hello.go](./example/hello/hello.go)文件

```go
package main
import "fmt" 
func main() {
   fmt.Println("Hello, World!")
}
```

命令运行`go`程序，在[hello.go](./example/hello/hello.go)这个当前目录下运行下面命令，可以输出`hello world!`。

```bash
go run hello.go
```

</details>

<details>
<summary>通过go命令编译运行</summary>

GO程序的代码是可以直接编译成`exe文件` 或者 `二进制文件`直接运行，在[hello.go](./example/hello/hello.go)目录下运行下面命令，即可把go程序编译成二进制文件

```bash
go build hello.go
```

上面命令文件可以编译成一个`hello`可执行文件，然后直接在当前目录下 `./hello` 运行，可以输出`hello world!`。

</details>

## 编程基础

<details>
<summary>内置关键字</summary>

```go
break      default       func      interface    select
case       defer         go        map          struct
chan       else          goto      package      switch
const      fallthrough   if        range        type
continue   for           import    retrun       var
```

</details>

<details>
<summary>预定义标识符</summary>

```go
append  bool    byte    cap     close  complex complex64 complex128 uint16
copy    false   float32 float64 imag   int     int8      int16      uint32
int32   int64   iota    len     make   new     nil       panic      uint64
print   println real    recover string true	   uint	     uint8      uintptr
```

</details>

<details>
<summary>行分隔符</summary>

- 在 Go 程序中，一行代表一个语句结束，不需要分隔符。
- 打算将多个语句写在同一行，它们则必须使用 `;` 人为区分，并不鼓励这种做法。

</details>

<details>
<summary>注释方法</summary>

```go
// 单行注释

/*
  多行注释
*/
```

</details>

<details>
<summary>标识符</summary>

- 标识符用来命名变量、类型等程序实体。
- 第一个字符必须是字母或下划线而不能是数字

有效标识符

```
mahesh   kumar   abc   move_name   a_123
myname50   _temp   j   a23b9   retVal
```

无效标识符

```bash
1ab  #（以数字开头）
case #（Go 语言的关键字）
a+b  #（运算符是不允许的）
```

</details>

<details>
<summary>包导入</summary>

```go
import "fmt"
import "os"
import "io"
```

简写方式如下

```go
import (
  "fmt"
  "os"
  "io"
)
```

</details>

<details>
<summary>包别名</summary>

```go
import(
  ff "fmt"
)

// 或者
import ff "fmt"

// 别名包调用
ff.Println('Hello World!')
```

</details>


<details>
<summary>省略调用</summary>

```go
import(
  . "fmt"
)
func main() {
  // 省略调用
  Println('Hello World!')
}
```

</details>


<details>
<summary>可见性规则</summary>

Go语言中约定使用 **大小写** 来决定常量、变量、类型、接口、结构或函数是否可以被外部包所调用

- 函数名字首字母 **小写** 即为 `private` 私有的
- 函数名字首字母 **大写** 即为 `public` 公有

</details>

## 基本类型

<details>
<summary>布尔型</summary>

```go
var b bool = true
```

- 长度：1字节
- 取值范围：true/false
- 只能使用true/false值，不能使用数字代替布尔值

</details>

<details>
<summary>整形 int/uint</summary>

```go
package main
import "fmt"
func main() {
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
}
```

- int/uint
- 根据平台可能为32/64位

</details>

<details>
<summary>8位整型 int8/uint8</summary>

```go
u8 := []uint8{98, 99}

a := byte(255)  //11111111 这是byte的极限， 因为 a := byte(256)//越界报错， 0~255正好256个数，不能再高了
b := uint8(255) //11111111 这是uint8的极限，因为 c := uint8(256)//越界报错，0~255正好256个数，不能再高了
c := int8(127)  //01111111 这是int8的极限， 因为 b := int8(128)//越界报错， 0~127正好128个数，所以int8的极限只是256的一半
d := int8(a)    //11111111 打印出来则是-0000001，int8(128)、int8(255)、int8(byte(255))都报错越界，因为int极限是127，但是却可以写：int8(a)，第一位拿来当符号了
e := int8(c)    //01111111 打印出来还是01111111
fmt.Printf("%08b %d \n", a, a)
fmt.Printf("%08b %d \n", b, b)
fmt.Printf("%08b %d \n", c, c)
fmt.Printf("%08b %d \n", d, d)
fmt.Printf("%08b %d \n", e, e)
```

- int8/uint8
- 长度：1字节
- 取值范围：-128~127/0~255

</details>

<details>
<summary>字节型 byte</summary>

```go
// 这里不能写成 b := []byte{"Golang"}，这里是利用类型转换。
b := []byte("Golang")
c := []byte("go")
d := []byte("Go")
println(b,c,d)
```

- byte(uint8别名)

**基本处理函数**

- `Contains()` 返回是否包含子切片
- `Count()` 子切片非重叠实例的数量
- `Map()` 函数，将byte 转化为Unicode，然后进行替换
- `Repeat()` 将切片复制count此，返回这个新的切片
- `Replace()` 将切片中的一部分 替换为另外的一部分
- `Runes()` 将 S 转化为对应的 UTF-8 编码的字节序列，并且返回对应的Unicode 切片
- `Join()` 函数，将子字节切片连接到一起。

可以参考下面列子来理解上面7个方法，例子 [byte.go](./example/byte/byte.go)

```go
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
```

</details>

<details>
<summary>16位整型 int16/uint16</summary>

- int16/uint16
- 长度：2字节
- 取值范围：-32768~32767/0~65535

</details>

<details>
<summary>32位整型 int32(别名rune)/uint32</summary>

- int32(别名rune)/uint32
- 长度：4字节
- 取值范围：-2^32/2~2^32/2-1/0~2^32-1

</details>

<details>
<summary>64位整型 int64/uint64</summary>

- int64/uint64
- 长度：8字节
- 取值范围：-2^64/2~2^64/2-1/0~2^64-1

</details>

<details>
<summary>浮点型 float32/float64</summary>

```go
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
```

实例：[float.go](./example/float/float.go)

- float32/float64
- 长度：4/8字节
- 小数位：精确到 7/15 位小数

</details>

<details>
<summary>复数 complex64/complex128</summary>

- complex64/complex128
- 长度：8/16

</details>

<details>
<summary>指针整数型 uintptr</summary>

- uintptr
- 保存指正的 32 位或者 64 位整数型

</details>

<details>
<summary>数组类型 array</summary>

数组声明语法

```go
var variable_name [SIZE]variable_type
```

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。下面是一个简单的对数组操作的例子[array.go](./example/array/array.go)

```go
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
```

初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：

```go
var array1 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：

```go
float32 salary = array1[9]
```

以上实例读取了数组`array1`第`10`个元素的值。

多维数组，下面例子

```go
// 三行四列
a = [3][4]int{  
 {0, 1, 2, 3} ,   /*  第一行索引为 0 */
 {4, 5, 6, 7} ,   /*  第二行索引为 1 */
 {8, 9, 10, 11}   /*  第三行索引为 2 */
}
```

访问多维数组

```go
// 访问第2行第3列
int val = a[2][3]
```

</details>

<details>
<summary>结构类型 struct</summary>

```go
type identifier struct {
    field1 type1
    field2 type2
    ...
}
```

简单的结构体

```go
type T struct {a, b int}
```

结构体里的字段都有 名字，像 `field1`、`field2` 等，如果字段在代码中从来也不会被用到，那么可以命名它为 `_`。上面简单的结构体定义，下面调用方法：

```go
var s T
s.a = 5
s.b = 8
```

数组可以看作是一种结构体类型，不过它使用下标而不是具名的字段。

```go
var t *T
t = new(T)
```

上面简单的管用语句方法`t := new(T)`，变量 `t` 是一个指向 `T` 的指针，此时结构体字段的值是它们所属类型的零值。

声明 `var t T` 也会给 `t` 分配内存，并零值化内存，但是这个时候 `t` 是类型`T`。在这两种方式中，`t` 通常被称做类型 `T` 的一个实例（instance）或对象（object）。

一个非常简单的例子[structs_fields.go](./example/structs/structs_fields.go)运行例子查看结果：

```bash
→ go run test/structs_fields.go

The int is: 10
The float is: 15.500000
The string is: Chris
&{10 15.5 Chris}
```

**使用 new**

</details>

<details>
<summary>字符串类型 string</summary>

```go
var str string //声明一个字符串
str = "Go lang"  //赋值
ch :=str[0]    //获取第一个字符
len :=len(str) //字符串的长度,len是内置函数 ,len=5
```

len函数是Go中内置函数，不引入strings包即可使用。len(string)返回的是字符串的字节数。len函数所支持的入参类型如下：

- len(Array) 数组的元素个数
- len(*Array) 数组指针中的元素个数,如果入参为nil则返回0
- len(Slice) 数组切片中元素个数,如果入参为nil则返回0
- len(map) 字典中元素个数,如果入参为nil则返回0
- len(Channel) Channel buffer队列中元素个数

</details>

<details>
<summary>接口类型 inteface</summary>

```go
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
```

实例：[inteface.go](./example/inteface/inteface.go)

</details>

<details>
<summary>函数类型 func</summary>

```go
package main
import "fmt"
type functinTyoe func(int, int) // 声明了一个函数类型
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
```

实例：[func.go](./example/func/func.go)

</details>

<details>
<summary>引用类型 func</summary>


**切片**

> 是一种可以动态数组，可以按我们的希望增长和收缩。

- slice

**Map**

> 是一种无序的键值对的集合。是一种集合，所以我们可以像迭代数组和 slice 那样迭代它。

- map

```go
// 通过 make 来创建
dict := make(map[string]int)
// 通过字面值创建
dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

// 给 map 赋值就是指定合法类型的键，然后把值赋给键
colors := map[string]string{}
colors["Red"] = "#da1337"

// 不初始化 map , 就会创建一个 nil map。nil map 不能用来存放键值对，否则会报运行时错误
var colors map[string]string
colors["Red"] = "#da1337"
// Runtime Error:
// panic: runtime error: assignment to entry in nil map

//选择是只返回值，然后判断是否是零值来确定键是否存在。
value := colors["Blue"]
if value != "" {
  fmt.Println(value)
}
```

在函数间传递 map 不是传递 map 的拷贝。所以如果我们在函数中改变了 map，那么所有引用 map 的地方都会改变

```go
func main() {
  colors := map[string]string{
     "AliceBlue":   "#f0f8ff",
     "Coral":       "#ff7F50",
     "DarkGray":    "#a9a9a9",
     "ForestGreen": "#228b22",
  }
  for key, value := range colors {
      fmt.Printf("Key: %s  Value: %s\n", key, value)
  }
  removeColor(colors, "Coral")
  for key, value := range colors {
      fmt.Printf("Key: %s  Value: %s\n", key, value)
  }
}
func removeColor(colors map[string]string, key string) {
    delete(colors, key)
}
```


**通道**

- chan

</details>

<details>
<summary>类型别名</summary>

```go
type (
  byte int8
  rune init32
  文本 string
)
var b 文本
b = "别名类型，可以是中文！"
```

</details>

## 常量变量

<details>
<summary>常量 const</summary>

```go
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
```

上面例子[const.go](./example/const/const.go)

**iota** 特殊常量，可以认为是一个可以被编译器修改的常量。[iota.go](./example/iota/iota.go)

```go
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
```

</details>

<details>
<summary>变量 var</summary>

```go
var (
  name  = "gopher"
  name1 = "1"
)
// 变量声明
var a int
a = 11 /* 赋值 */

// 变量声明 并赋值
var b int = 12

// 应用在函数体内的方式
var a, b, c, d int = 1, 2, 3, 4
// a =1 
// b =2 
// c =3 
// d =4 


var a, _, c, d int = 1, 2, 3, 4
// 忽略 _ 返回值忽略
```

- 全局变量名 以大写开头
- 全局变量不可以省略 var ，可以使用并行的方式
- 所有变量都可以使用类型推断
- 局部变量不可以使用`var()`简写的形式

</details>

<details>
<summary>变量的类型转换</summary>

```go
// 只能类型显式转换
var a float32 = 1.1
// 省略var, 简短形式，使用 := 赋值操作符
b := int(a)
// 不兼容的类型不能转换类型
```

</details>

<details>
<summary>多变量声明</summary>

```go
var x, y int
// 这种因式分解关键字的写法一般用于声明全局变量
var (
  a int
  b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"
```

</details>

## 语言运算符

<details>
<summary>算术运算符</summary>

```go
package main
import "fmt"
func main() {
  var a int = 21
  var b int = 10
  var c int
  c = a + b
  fmt.Printf("第一行 - c 的值为 %d\n", c ) // 第一行 - c 的值为 31
  c = a - b
  fmt.Printf("第二行 - c 的值为 %d\n", c ) // 第二行 - c 的值为 11
  c = a * b
  fmt.Printf("第三行 - c 的值为 %d\n", c ) // 第三行 - c 的值为 210
  c = a / b
  fmt.Printf("第四行 - c 的值为 %d\n", c ) // 第四行 - c 的值为 2
  c = a % b
  fmt.Printf("第五行 - c 的值为 %d\n", c ) // 第五行 - c 的值为 1
  a++
  fmt.Printf("第六行 - c 的值为 %d\n", a ) // 第六行 - c 的值为 22
  a--
  fmt.Printf("第七行 - c 的值为 %d\n", a ) // 第七行 - c 的值为 21
}
```

下表列出了所有Go语言的算术运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| + | 相加  | A + B 输出结果 30 |
| - | 相减  | A - B 输出结果 -10 |
| * | 相乘  | A * B 输出结果 200 |
| / | 相除  | B / A 输出结果 2 |
| % | 求余  | B % A 输出结果 0 |
| ++ | 自增 | A++ 输出结果 11 |
| -- | 自减 | A-- 输出结果 9 |

</details>

<details>
<summary>关系运算符</summary>

```go
package main
import "fmt"
func main() {
   var a int = 21
   var b int = 10
   if( a == b ) {
      fmt.Printf("第一行 - a 等于 b\n" )
   } else {
      fmt.Printf("第一行 - a 不等于 b\n" )
   }
   if ( a < b ) {
      fmt.Printf("第二行 - a 小于 b\n" )
   } else {
      fmt.Printf("第二行 - a 不小于 b\n" )
   } 
   
   if ( a > b ) {
      fmt.Printf("第三行 - a 大于 b\n" )
   } else {
      fmt.Printf("第三行 - a 不大于 b\n" )
   }
   /* 让我们改变a和b的值 */
   a = 5
   b = 20
   if ( a <= b ) {
      fmt.Printf("第四行 - a 小于等于 b\n" )
   }
   if ( b >= a ) {
      fmt.Printf("第五行 - b 大于等于 a\n" )
   }
}
```

下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| ==  | 检查两个值是否相等，如果相等返回 True 否则返回 False。 |	(A == B) 为 False |
| !=  | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。 |	(A != B) 为 True |
| > | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。 |	(A > B) 为 False |
| < | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。 |	(A < B) 为 True |
| >=  | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 |	(A >= B) 为 False |
| <=  | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 | A <= B) 为 True |

</details>

<details>
<summary>逻辑运算符</summary>

```go
package main
import "fmt"
func main() {
  var a bool = true
  var b bool = false
  if ( a && b ) {
    fmt.Printf("第一行 - 条件为 true\n" )
  }
  if ( a || b ) {
    fmt.Printf("第二行 - 条件为 true\n" )
  }
  /* 修改 a 和 b 的值 */
  a = false
  b = true
  if ( a && b ) {
    fmt.Printf("第三行 - 条件为 true\n" )
  } else {
    fmt.Printf("第三行 - 条件为 false\n" )
  }
  if ( !(a && b) ) {
    fmt.Printf("第四行 - 条件为 true\n" )
  }
}
```

下表列出了所有Go语言的逻辑运算符。假定 A 值为 `True`，B 值为 `False` d。

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| && | 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 |	(A && B) 为 False |
| \|\| | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。|	(A || B) 为 True |
| ! | 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。 |	!(A && B) 为 True |

</details>

<details>
<summary>位运算符</summary>

```go
package main
import "fmt"
func main() {

  var a uint = 60	/* 60 = 0011 1100 */  
  var b uint = 13	/* 13 = 0000 1101 */
  var c uint = 0          

  c = a & b       /* 12 = 0000 1100 */ 
  fmt.Printf("第一行 - c 的值为 %d\n", c ) // 第一行 - c 的值为 12

  c = a | b       /* 61 = 0011 1101 */
  fmt.Printf("第二行 - c 的值为 %d\n", c )  // 第二行 - c 的值为 61

  c = a ^ b       /* 49 = 0011 0001 */
  fmt.Printf("第三行 - c 的值为 %d\n", c ) // 第三行 - c 的值为 49

  c = a << 2     /* 240 = 1111 0000 */
  fmt.Printf("第四行 - c 的值为 %d\n", c ) // 第四行 - c 的值为 240

  c = a >> 2     /* 15 = 0000 1111 */
  fmt.Printf("第五行 - c 的值为 %d\n", c )  // 第五行 - c 的值为 15
}
```

Go 语言支持的位运算符如下表所示。假定 A 为60，B 为13：

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| & | 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 |	(A & B) 结果为 12, 二进制为 0000 1100 |
| \| | 按位或运算符 \| 是双目运算符。 其功能是参与运算的两数各对应的二进位相或。 |	(A \| B) 结果为 61, 二进制为 0011 1101 |
| ^ | 按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。 |	(A ^ B) 结果为 49, 二进制为 0011 0001 |
| << | 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 | A << 2 结果为 240 ，二进制为 1111 0000 |
| >> | 右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。| A >> 2 结果为 15 ，二进制为 0000 1111 |

</details>

<details>
<summary>赋值运算符</summary>

```go
package main
import "fmt"
func main() {
  var a int = 21
  var c int

  c =  a
  fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )
  // 第 1 行 - =  运算符实例，c 值为 = 21
  c +=  a
  fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )
  // 第 2 行 - += 运算符实例，c 值为 = 42
  c -=  a
  fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )
  // 第 3 行 - -= 运算符实例，c 值为 = 21
  c *=  a
  fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )
  // 第 4 行 - *= 运算符实例，c 值为 = 441
  c /=  a
  fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )
  // 第 5 行 - /= 运算符实例，c 值为 = 21
  c  = 200; 
  c <<=  2
  fmt.Printf("第 6 行  - <<= 运算符实例，c 值为 = %d\n", c )
  // 第 6 行  - <<= 运算符实例，c 值为 = 800
  c >>=  2
  fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )
  // 第 7 行 - >>= 运算符实例，c 值为 = 200
  c &=  2
  fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )
  // 第 8 行 - &= 运算符实例，c 值为 = 0
  c ^=  2
  fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )
  // 第 9 行 - ^= 运算符实例，c 值为 = 2
  c |=  2
  fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )
  // 第 10 行 - |= 运算符实例，c 值为 = 2
}
```

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| = | 简单的赋值运算符，将一个表达式的值赋给一个左值s | C = A + B 将 A + B 表达式结果赋值给 C |
| += | 相加后再赋值s | C += A 等于 C = C + A |
| -= | 相减后再赋值s | C -= A 等于 C = C - A |
| *= | 相乘后再赋值s | C *= A 等于 C = C * A |
| /= | 相除后再赋值s | C /= A 等于 C = C / A |
| %= | 求余后再赋值s | C %= A 等于 C = C % A |
| <<= | 左移后赋值s | C <<= 2 等于 C = C << 2 |
| >>= | 右移后赋值s | C >>= 2 等于 C = C >> 2 |
| &= | 按位与后赋值s | C &= 2 等于 C = C & 2 |
| ^= | 按位异或后赋值s | C ^= 2 等于 C = C ^ 2 |
| \|= | 按位或后赋值s | C \|= 2 等于 C = C \| 2 |

</details>

<details>
<summary>其他运算符</summary>

```go
package main
import "fmt"
func main() {
  var a int = 4
  var b int32
  var c float32
  var ptr *int

  /* 运算符实例 */
  fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a ); // 第 1 行 - a 变量类型为 = int
  fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b ); // 第 2 行 - b 变量类型为 = int32
  fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c ); // 第 3 行 - c 变量类型为 = float32

  /*  & 和 * 运算符实例 */
  ptr = &a	/* 'ptr' 包含了 'a' 变量的地址 */
  fmt.Printf("a 的值为  %d\n", a);   // a 的值为  4
  fmt.Printf("*ptr 为 %d\n", *ptr); // *ptr 为 4
}
```

| 运算符 | 描述 | 实例 |
| ---- | ---- | ---- |
| & | 返回变量存储地址 | &a; 将给出变量的实际地址。 |
| * | 指针变量。 | *a; 是一个指针变量 |

</details>

<details>
<summary>运算符优先级</summary>

```go
package main
import "fmt"
func main() {
  var a int = 20
  var b int = 10
  var c int = 15
  var d int = 5
  var e int;
  // 通过使用括号来临时提升某个表达式的整体运算优先级。
  e = (a + b) * c / d;      // ( 30 * 15 ) / 5
  fmt.Printf("(a + b) * c / d 的值为 : %d\n",  e );
  e = ((a + b) * c) / d;    // (30 * 15 ) / 5
  fmt.Printf("((a + b) * c) / d 的值为  : %d\n" ,  e );
  e = (a + b) * (c / d);   // (30) * (15/5)
  fmt.Printf("(a + b) * (c / d) 的值为  : %d\n",  e );
  e = a + (b * c) / d;     //  20 + (150/5)
  fmt.Printf("a + (b * c) / d 的值为  : %d\n" ,  e );  
}
```

有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

| 优先级 | 运算符 |
| ---- | ---- |
| 7 | ^ ! |
| 6 | * / % << >> & &^ |
| 5 | + - \| ^ |
| 4 | == != < <= >= > |
| 3 | <- |
| 2 | && |
| 1 | \|\| |

</details>

## 资源导航

<details>
<summary>官方</summary>

- [Playground](http://play.golang.org)：Go 语言代码在线运行

</details>

<details>
<summary>国内镜像</summary>

- [Go 指南国内镜像](http://tour.golangtc.com/)
- [Go 语言国内下载镜像](http://www.golangtc.com/download)
- [Go 官方网站国内镜像](http://docs.studygolang.com/)

</details>

<details>
<summary>Web 框架</summary>

- [Macaron](https://go-macaron.com/)：模块化 Web 框架
- [Beego](http://beego.me/)：重量级 Web 框架
- [Revel](https://github.com/revel/revel)：较早成熟的重量级 Web 框架
- [Martini](https://github.com/go-martini/martini): 一个强大为了编写模块化 Web 应用而生的 Go 语言框架

</details>

<details>
<summary>ORM 以及数据库驱动</summary>

- [xorm](https://github.com/go-xorm/xorm)：支持 MySQL、PostgreSQL、SQLite3 以及 MsSQL
- [mgo](http://labix.org/mgo)：MongoDB 官方推荐驱动

</details>

<details>
<summary>辅助站点</summary>

- [Go Walker](https://gowalker.org)：Go 语言在线 API 文档
- [gobuild.io](http://gobuild.io/)：Go 语言在线二进制编译与下载
- [Rego](http://regoio.herokuapp.com/)：Go 语言正则在线测试
- [gopm.io](https://gopm.io)：科学下载第三方包

</details>
