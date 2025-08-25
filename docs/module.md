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