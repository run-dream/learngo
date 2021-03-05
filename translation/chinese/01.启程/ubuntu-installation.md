# Linux 安装指南

如果你想使用 snap, 你可按照以下方式轻松的安装 Go 语言:

    sudo snap install go --classic
    
否则, 请跟从以下步骤:

## 1. 更新本地包

  ```bash
  sudo apt-get update
  ```

## 2. 安装 git

  ```bash
  sudo apt-get install git
  ```

## 3. 安装 Go

有两种方式可选:

1- 通过 Web: 选择 Linux 然后开始下载

  ```bash
  firefox https://golang.org/dl
  ```

2- 通过 snap: 使用这种方式可以跳过第5步

  ```bash
  sudo snap install go --classic
  ```

## 4. 将 Go 复制到合适的目录

1. 找出下载的文件名

2. 使用文件名来解压文件

    ```bash
    gofile="DELETE_THIS_AND_TYPE_THE_NAME_OF_THE_DOWNLOADED_FILE_HERE (without its extension)"

    tar -C /usr/local -xzf ~/Downloads/$gofile
    ```

## 5. 将 Go 语言可执行目录加入到 PATH 中

1. 将 `go/bin` 目录添加到 `$PATH` 以便运行基础的 Go 命令

    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    ```

2. 将 "$HOME/go/bin" 添加到 $PATH

    ```bash
    echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.profile
    ```

## 安装 Go tools:

* 有一些非常方便的工具可以用来简化开发 (like goimports)

* 在使用 `go get` 之前需要安装代码版本管理程序，比如 Git

* 一下命令将创建 `~/go` 目录，并在新建的目录下下载 go tools

    * 这个目录也是你应该将代码放置的目录(如果你不打算使用 Go Modules 的话)
    ```bash
    go get -v -u golang.org/x/tools/...
    ```

## 安装VSCode (可选)

注意: 如果你喜欢的话可以选择其他编辑器. 然而, 这门课使用 Visual Studio Code (VSCode).

1. 打开 "Ubuntu Software" 应用

2. 搜索 VSCode, 然后点击 "Install"


## 可选步骤:

1. 在`$GOPATH`以外的目录，新建文件 hello.go

    ```bash
    cat <<EOF > hello.go
    package main

    import "fmt"

    func main() {
        fmt.Println("hello gopher!")
    }
    EOF
    ```

2. 运行程序

    ```bash
    go run hello.go
    应该输出: hello gopher!
    ```

<div style="page-break-after: always;"></div>

> 更多教程: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2018 Inanc Gumus
> 
> 学习 Go 语言编程课程
> 
> [点击此查看许可证](https://creativecommons.org/licenses/by-nc-sa/4.0/)
