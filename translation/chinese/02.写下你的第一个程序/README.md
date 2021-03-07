# 指南: 写下你的第一个 Go 程序

你好!

在你学完这个章节的所有课程之后，你可以记录下这个指南用于参考

你也在学完这份视频课程，可以打印这份课程

愿你学有所成!

---

## 输入命令行指令:

* 进入某个文件夹: `cd directoryPath`

* **WINDOWS:**

    * 显示文件夹下所有文件: `dir`

* **OS X & LINUXes:**

    * 显示文件夹下所有文件: `ls`

## 构建 & 运行 GO 程序:

* **构建 Go 程序:**

    * 在程序目录下, 输入:
        * `go build`

* **运行 Go 程序:**

    * 在程序目录下, 输入:
        * `go run main.go`

## $GOPATH 是什么?

* _$GOPATH_ 是指向你下载和存储自己的 Go 文件的目录的环境变量.

    * **Windows**系统上, 位于: `%USERPROFILE%\go`

    * **On OS X & Linux**, 位于: `~/go`

    * **注意:** 不要手动设置你的 GOPATH。 默认位于你的用户目录。

* **GOPATH 有三个目录:**

    * **src:** 包含了你自己和其他人的源代码. 你可以在这个目录下的子目录构建并运行程序

    * **pkg:** 包含编译好的包文件. Go 使用此目录加速编译和链接

    * **bin:** 包含编译好的可执行的 Go 程序。当你在程序目录调用 go install 以后, Go 会将可执行文件加到此目录.

        * _如果他们不处于环境变量 `PATH`中时，你可能需要把他们加进去._

## 你的源代码应该放在哪里?

* `$GOPATH/src/github.com/yourUsername/programDirectory`

* **案例:**

    * 我的 Github 用户名: inancgumus

    * 所以, 我把我所有的程序放在这个目录下: `~/go/src/github.com/inancgumus/`

    * 加入我有一个叫做 hello 的程序, 我将其放在: `~/go/src/github.com/inancgumus/hello`

## 第一个程序

* **新建目录:**
    * **OS X & Linux (or git bash):**
        * 新建目录:
            * `mkdir -p ~/go/src/yourname/hello`
        * 进入目录:
            * `cd ~/go/src/yourname/hello`

    * **Windows:**
        * 新建目录:
            * `mkdir c:\Go\src\yourname\hello`
        * 进入目录:
            * `cd c:\Go\src\yourname\hello`

* 在 Visual Studio Code 下 新建 `code main.go`
* 添加以下代码并保存
* 返回命令行
    * 运行: `go run main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hi! I want to be a Gopher!")
}
```

以上

## 注意:

* *Go Modules* 的方式支持你在你想要的任目录下运行你的程序。这仍只是一个实验特性, 当它文档以后, 我会更新这个课程关于 Go Modules 的信息

<div style="page-break-after: always;"></div>

> 更多教程: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2018 Inanc Gumus
> 
> 学习 Go 语言编程课程
> 
> [点击此查看许可证](https://creativecommons.org/licenses/by-nc-sa/4.0/)