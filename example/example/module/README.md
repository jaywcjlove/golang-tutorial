
## 初始化工程

```bash
mkdir mod2
cd mod2
go mod init mod2
```

## 安装包

```bash
go get -m github.com/gin-gonic/gin@v1.3.0

go get github.com/gorilla/mux    # 匹配最新的一个 tag
go get github.com/gorilla/mux@latest    # 和上面一样
go get github.com/gorilla/mux@v1.6.2    # 匹配 v1.6.2
go get github.com/gorilla/mux@e3702bed2 # 匹配 v1.6.2
go get github.com/gorilla/mux@c856192   # 匹配 c85619274f5d
go get github.com/gorilla/mux@master    # 匹配 master 分支
```

1. 通过 `go get` 切换版本
2. 执行 `go mod vendor` 将依赖复制到vendor下
3. 通过 `go mod tidy` 命令可以移除 `go.mod` 中不再使用的依赖

## 设置代理

```bash
# bash mac
export GOPROXY=https://goproxy.io
```

## mod 命令

```bash
go help mod
# Go mod provides access to operations on modules.

# Note that support for modules is built into all the go commands,
# not just 'go mod'. For example, day-to-day adding, removing, upgrading,
# and downgrading of dependencies should be done using 'go get'.
# See 'go help modules' for an overview of module functionality.

Usage:

	go mod <command> [arguments]

The commands are:

	download    download modules to local cache
	edit        edit go.mod from tools or scripts
	graph       print module requirement graph
	init        initialize new module in current directory
	tidy        add missing and remove unused modules
	vendor      make vendored copy of dependencies
	verify      verify dependencies have expected content
	why         explain why packages or modules are needed

Use "go help mod <command>" for more information about a command.
```

- `download`: download modules to local cache (下载依赖的module到本地cache)
- `edit`: edit go.mod from tools or scripts (编辑go.mod文件)
- `graph`: print module requirement graph (打印模块依赖图)
- `init`: initialize new module in current directory (再当前文件夹下初始化一个新的module, 创建go.mod文件)
- `tidy`: add missing and remove unused modules (增加丢失的module，去掉未用的module)
- `vendor`: make vendored copy of dependencies (将依赖复制到vendor下)
- `verify`: verify dependencies have expected content (校验依赖)
- `why`: explain why packages or modules are needed (解释为什么需要依赖)