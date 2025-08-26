# 包
当你想去命名一个包的时候，可以通过 package 关键字，提供一个值，而不是完整的层次结构（例如：「shopping」或者 「db」）。当你想去导入一个包的时候，你需要指定完整路径。

循环引用:
把公用代码抽取到models,utils等无外部依赖模块，然后导入:
```go
package db

import (
	"moduleDemo/shopping/models" // 同一个module内导入包总是从module名开始
)

func LoadItem2(id int) *models.Item {
	return &models.Item{
		Price: 9.001,
	}
}
```

## 可见性
Go 用了一个简单的规则去定义什么类型和函数可以包外可见。如果类型或者函数名称以一个大写字母开始，它就具有了包外可见性。如果以一个小写字母开始，它就不可以。

## go get
> go get 包名@版本号
也可以从github: 安装到 $GOPATH\pkg\mod
> go get github.com/gin-gonic/gin@v1.9.1

获取最新发布版:
>go get github.com/gin-gonic/gin@latest

安装 master 分支最新提交:
>go get github.com/gin-gonic/gin@master

查看$GOPATH:
>cmd: echo $GOPATH
>powershell: $env:GOPATH

使用:
```go
import (
  "github.com/mattn/go-sqlite3"
)
```

安装所有包：
>go get
它将更新所有包
>go get -u
 查看当前项目依赖版本
>go list -m all

# 接口
接口是定义了合约但并没有实现的类型。举个例子：
```go
type Logger interface {
  Log(message string)
}
```
接口有助于将代码与特定的实现进行分离

针对接口而不是具体实现的编程会使我们很轻松的修改（或者测试）任何代码都不会产生影响。
```go
type Server struct {
  logger Logger
}
```
或者是一个函数参数（或者返回值）：
```go
func process(logger Logger) {
  logger.Log("hello!")
}
```
在 Go 中，下面的情况是隐式发生的。如果你的结构体有一个函数名为 Log 且它有一个 string 类型的参数并没有返回值，那么这个结构体被视为 Logger 。这减少了使用接口的冗长：
```go
type ConsoleLogger struct {}
func (l ConsoleLogger) Log(message string) { // 值接收者
// func (l *ConsoleLogger) Log(...)  // 指针接收者
  fmt.Println(message)
}
```
这其实可以理解为：一个隐藏了第一个参数的函数，这个参数就是 (l ConsoleLogger)
接收者让语法更像面向对象中的“对象.方法()”

Go 倾向于使用小且专注的接口。Go 的标准库基本上由接口组成。像 io 包有一些常用的接口诸如 io.Reader ， io.Writer ， io.Closer 等。如果你编写的函数只需要一个能调用 Close() 的参数，那么你应该接受一个 io.Closer 而不是像 io 这样的父类型。


#  module
Create module: In the greetings directory
>go mod init example.com/greetings

redirect Go tools from its module path (where the module isn't) to the local directory: In /hello
> go mod edit -replace example.com/greetings=../greetings
go mod edit 是 Go 语言中用于 直接编辑 go.mod 文件 的命令行工具: 
>go mod edit [flags] [file]

synchronize the example.com/hello module's dependencies, adding those required by the code, but not yet tracked in the module.: in /hello:
>go mod tidy
会删除未使用的依赖，确保 go.mod 和 go.sum 干净

Run: in /hello
>go run .

# Build
## The **go build** command 
compiles the packages, along with their dependencies, but it doesn't install the results.

Discover the Go install path, where the go command will install the current package
> go list -f '{{.Target}}'

Add the Go install directory to your system's shell path
> export PATH=$PATH:/path/to/your/install/directory
or:
> set PATH=%PATH%;C:\path\to\your\install\directory

change the install target by setting the GOBIN variable to the shell path of go:
> go env -w GOBIN=/path/to/your/bin

Once you've updated the shell path, run the go install command to compile and install the package.
>go install

Run your application by simply typing its name:
> hello

## The **go install** command 
compiles and installs the packages.

# multi-module workspaces
Initialize the module: in /workspaceDemo/hello
>go mod init example.com/hello
Plz confirm hello/go.sum, hello/go.mod(验证下载一致性) exists so that /hello can be a valid module!

Add a dependency on the golang.org/x/example/hello/reverse package by using go get.
> go get golang.org/x/example/hello/reverse


Initialize the workspace, in /workspaceDemo
> go work init ./hello
go work init <dir> 的含义是：以指定目录中的模块作为初始成员来初始化 workspace。

In the workspace directory, run:
> go run ./hello
The Go command includes all the modules in the workspace as main modules.

## Download and modify the golang.org/x/example/hello module
In /workspaceDemo:
>git clone https://go.googlesource.com/example

Add the module to the workspace：
>go work use ./example/hello

Add the new function to reverse a number to the golang.org/x/example/hello/reverse package.
Create a new file named int.go in the workspace/example/hello/reverse directory containing the following contents:
```go
package reverse

import "strconv"

// Int returns the decimal reversal of the integer i.
func Int(i int) int {
    i, _ = strconv.Atoi(String(strconv.Itoa(i)))
    return i
}
```

Modify the contents of workspace/hello/hello.go to contain the following contents:
```go
package main

import (
    "fmt"

    "golang.org/x/example/hello/reverse"
)

func main() {
    fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
```

In /workspaceDemo:
>go run ./hello
Go resolves the golang.org/x/example/hello/reverse import using the go.work file.

go.work can be used instead of adding **replace** directives to work across multiple modules.

Since the two modules are in the same workspace it’s easy to make a change in one module and use it in another.

make a release of the golang.org/x/example/hello module, for example at v0.1.0  by tagging a commit on the module’s version control repository： 
>https://go.dev/doc/modules/release-workflow

Other cmds:
- **go work use [-r] [dir]** adds a use directive to the go.work file for dir, if it exists, and removes the use directory if the argument directory doesn’t exist. The -r flag examines subdirectories of dir recursively.
- **go work edit** edits the go.work file similarly to go mod edit
- **go work sync** syncs dependencies from the workspace’s build list into each of the workspace modules.