# IDE
- Launch the VS Code editor
- Open the extension manager or alternatively, press Ctrl + Shift + x
- In the search box, type "go" and hit enter
- Find the Go extension by the GO team at Google and install the extension
- After the installation is complete, open the command palette by pressing Ctrl + Shift + p
- Run the Go: Install/Update Tools command
- Select all the provided tools and click OK

# Start
在 go 中程序入口必须是 main 函数，并且在 main 包内

Init like npm init:
> go mod init example.com/hello 

Run:
>go run .\helloworld.go
> go run .

If you want to save the program as an executable, type and run:
>go build .\helloworld.go

# build
 go run 命令已经包含了编译和运行。它使用一个临时目录来构建程序，执行完然后清理掉临时目录。你可以执行以下命令来查看临时文件的位置：
>go run --work main.go


# 本地文档
> godoc -http=:6060
然后浏览器中访问 http://localhost:6060