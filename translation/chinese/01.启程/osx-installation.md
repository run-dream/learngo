# OS X 安装指南

## 注意

如果你已安装了[homebrew](https://brew.sh), 你可按照以下方式轻松的安装 Go 语言:

```
# 如果你尚未安装 git, 按照以下命令安装:
brew install git

# 然后安装 go
brew install go

# 在 ~/.bash_profile 中添加 GOBIN 路径
export PATH=${HOME}/go/bin:$PATH
```

## 1- 安装 VS Code 编辑器

1. 安装它，但不要打开它
2. 打开: [https://code.visualstudio.com](https://code.visualstudio.com)
3. 选择 OS X (Mac) 然后开始下载
4. 解压下载好的文件并将其移动到你的`~/Applications`目录

## 2- 安装 Git

1. 获取并运行安装程序。 打开: [https://git-scm.com/downloads](https://git-scm.com/downloads)
2. 选择VS Code作为默认编辑器

## 3- 安装 Go

1. 打开 [https://golang.org/dl](https://golang.org/dl)
2. 选择 OS X (Mac)
3. 启动安装器

## 4- 配置 VS Code

1. 打开 VS Code, 在左侧的扩展页, 搜索 "go" 并安装
2. 完全关闭VS Code, 然后重新打开它

3. 打开查看菜单; 选择 **命令面板**
    1. 或者通过组合键 `cmd+shift+p`
    2. 输入: `go install`
    3. 选择 _"Go: Install/Update Tools"_
    4. 选中所有检验栏

4. 执行完以后, 再次打开 **命令面板**
    1. 输入: `shell`
    2. 选择: _"Install 'code' command in PATH"_
        1. **注意:** 如果你是在Windows系统上操作的不需要执行此操作

## 收工啦! 玩得开心! 🤩

<div style="page-break-after: always;"></div>

> 更多教程: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2018 Inanc Gumus
> 
> 学习 Go 语言编程课程
> 
> [点击此查看许可证](https://creativecommons.org/licenses/by-nc-sa/4.0/)
