# go-gin-study
在前面我们已经了解到 Go 依赖包管理的历史情况，接下来我们将正式的进入使用，首先你需要有一个你喜欢的目录，例如：$ mkdir ~/go-application && cd ~/go-application，然后执行如下命令：
```shell script
$ mkdir go-gin-example && cd go-gin-example

$ go env -w GO111MODULE=on

$ go env -w GOPROXY=https://goproxy.cn,direct

$ go mod init github.com/EDDYCJY/go-gin-example
go: creating new go.mod: module github.com/EDDYCJY/go-gin-example
```
$ ls
go.mod
mkdir xxx && cd xxx：创建并切换到项目目录里去。
go env -w GO111MODULE=on：打开 Go modules 开关（目前在 Go1.13 中默认值为 auto）。
go env -w GOPROXY=...：设置 GOPROXY 代理，这里主要涉及到两个值，第一个是 https://goproxy.cn，它是由七牛云背书的一个强大稳定的 Go 模块代理，可以有效地解决你的外网问题；第二个是 direct，它是一个特殊的 fallback 选项，它的作用是用于指示 Go 在拉取模块时遇到错误会回源到模块版本的源地址去抓取（比如 GitHub 等）。
go mod init [MODULE_PATH]：初始化 Go modules，它将会生成 go.mod 文件，需要注意的是 MODULE_PATH 填写的是模块引入路径，你可以根据自己的情况修改路径。
在执行了上述步骤后，初始化工作已完成，我们打开 go.mod 文件看看，如下：

module github.com/EDDYCJY/go-gin-example

go 1.13
默认的 go.mod 文件里主要是两块内容，一个是当前的模块路径和预期的 Go 语言版本。
