# WINDOWS 安装指南

## NOTE

如果你使用 [chocolatey.org](https://chocolatey.org/) 包管理器, 你可按照以下方式轻松的安装 Go 语言:

```
choco install golang
```

## 1- 安装 Visual Studio Code 编辑器

1. 安装 VSCode, 但是不打开它
2. 打开: [https://code.visualstudio.com](https://code.visualstudio.com)
3. 选择 Windows，然后下载
4. 运行安装器

## 2- 安装 Git

1. 获取并运行安装程序。 打开: [https://git-scm.com/downloads](https://git-scm.com/downloads)
2. 选择VS Code作为默认编辑器
3. 勾选所有选择项
4. 选择: _"Use Git from the Windows Command Prompt"_
5. 编码: 选中: _"Checkout as is..." 选项_

## 3- 安装 Go

1. 打开 [https://golang.org/dl](https://golang.org/dl)
2. 选择 Windows 然后下载
3. 启动安装器

## 4- 配置 VS Code

1. 打开 VS Code, 在左侧的扩展页, 搜索 "go" 并安装
2. 完全关闭VS Code, 然后重新打开它

3. 打开查看菜单; 选择 **命令面板**
    1. 或者通过组合键 `cmd+shift+p`
    2. 输入: `go install`
    3. 选择 _"Go: Install/Update Tools"_
    4. 选中所有检验栏

## 5- 使用 Git-Bash

* 这节课中我将使用 bash 命令。Bash 只是在 osx 和 Linux 中使用的一个命令行界面。它是最流行的命令行界面之一。因此，如果您也想使用它，可以使用您安装的 Git-Bash，而不是使用 Windows 默认命令行。使用 Git-Bash，您可以像在 osx 或 Linux 上一样在命令行中键入命令。

* 如果不想使用 Git-Bash，也可以。这取决于你。但请注意，我将主要使用 bash 命令。因为它还允许更高级的命令

* 您也可以选择使用比Git-Bash更强大的替代方法：[Linux Subsystem for Windows](https://docs.microsoft.com/en-us/windows/wsl/install-win10)

* **因此，要使用 Git-Bash，请遵循以下步骤:**
    1. 只需从开始栏中搜索 git bash
    2. 或者, 如果有的话, 点击你桌面上的图标

    3. 同时, 设置 VS Code 默认使用 git-bash:
        1. 打开 VS Code
        2. 进入 命令面板
            1. 输入: `terminal`
            2. 选择: _"Terminal: Select Default Shell"_
            3. 最后, 选择: _"Git Bash"_

    4. **注意:** 通常情况下, 你可以在 `c:\` 目录下找到捏文件, 然而, 当你使用 git bash, 你可以使用 `/c/` 目录。 事实上是同一个目录, 只是一个缩写而已.


## 收工啦! 玩得开心! 🤩

<div style="page-break-after: always;"></div>

> 更多教程: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2018 Inanc Gumus
> 
> 学习 Go 语言编程课程
> 
> [点击此查看许可证](https://creativecommons.org/licenses/by-nc-sa/4.0/)
