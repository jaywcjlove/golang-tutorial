Go语言快速入门
---

安装之前需要了解及新建几个必要的文件目录：

- GOROOT 目录，该目录为解压压缩包所存放的目录。（建议 linux 环境解压到 /usr/local 目录，windows 环境解压到 C:\ProgramFiles 目录）
- 新建 GOPATH 目录，即为我们的“工作目录”，该目录可以有多个，建议只设置一个。
- 在 GOPATH 目录下新建 src 目录，该目录用于存放第三方库源码，以及存放我们的项目的源码。
- 在 GOPATH 目录下新建 bin 目录，该目录用于存放项目中所生成的可执行文件。
- 在 GOPATH 目录下新建 pkg 目录，该目录用于存放编译生成的库文件。


<details>
<summary>点击展开目录菜单</summary>

<!-- TOC -->

- [安装](#安装)
  - [CentOS7中安装](#centos7中安装)
- [编程基础](#编程基础)
- [基本类型](#基本类型)
- [常量变量](#常量变量)

<!-- /TOC -->

</details>

## 安装

### CentOS7中安装

<details>
<summary>yum安装</summary>

CentOS7 可以只用使用yum安装

```bash
yum install golang  
```

</details>

<details>
<summary>源码安装</summary>

源码下载

```bash
wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
tar zxvf go1.8.linux-amd64.tar.gz -C /usr/local
```

新建GOPATH目录

```bash
mkdir -p $HOME/gopath
```

编辑 `vim /etc/profile` 添加环境变量。

```bash
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
export GOPATH=$HOME/wwwroot/gofile
```

使其立即生效

```bash
source /etc/profile
```

其它命令

```bash
cat $GOROOT/VERSION  # 查看版本
$GOROOT/src/all.bash # 测试用例正确
```

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
<summary>注释方法</summary>

```go
// 单行注释

/*
  多行注释
*/
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

- 长度：1字节
- 取值范围：true/false
- 只能使用true/false值，不能使用数字代替布尔值

</details>

<details>
<summary>整形 int/uint</summary>

- int/uint
- 根据平台可能为32/64位

</details>

<details>
<summary>8位整型 int8/uint8</summary>

- int8/uint8
- 长度：1字节
- 取值范围：-128~127/0~255

</details>

<details>
<summary>字节型 byte</summary>

- byte(uint8别名)

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
<summary>浮点型 int64/uint64</summary>

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

```go
// 声明一个长度为5的整数数组
// 一旦数组被声明了，那么它的数据类型跟长度都不能再被改变。
var array [5]int

// 声明一个长度为5的整数数组
// 初始化每个元素
array := [5]int{12, 123, 1234, 12345, 123456}
```

</details>

<details>
<summary>结构类型 struct</summary>

- struct

</details>

<details>
<summary>字符串类型 string</summary>

- string

</details>

<details>
<summary>接口类型 inteface</summary>

- inteface

</details>

<details>
<summary>函数类型 func</summary>

- func

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
<summary>常量</summary>

```go
const (
  PI     = 3.14
  const1 = "1"
)
```

</details>

<details>
<summary>变量</summary>

- 全局变量名 以大写开头
- 全局变量不可以省略 var ，可以使用并行的方式
- 所有变量都可以使用类型推断
- 局部变量不可以使用`var()`简写的形式

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

</details>

<details>
<summary>变量</summary>

```go
// 只能类型显式转换
var a float32 = 1.1
b := int(a)

// 不兼容的类型不能转换类型
```

</details>

