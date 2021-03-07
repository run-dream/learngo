// 版权所有 © 2018 Inanc Gumus
// Go 语言编程课程
// 许可证: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// 更多教程  : https://learngoprogramming.com
// 亲自培训  : https://www.linkedin.com/in/inancgumus/
// 在 twitter 上关注我: https://twitter.com/inancgumus

// main 包是一个特殊的包
// 其允许 Go 创建可执行文件 
package main

/*
这是多行注释

import 关键字使得其他包的内容对此.go文件可用

import "fmt" 使你可以在此文件中使用 fmt 包的方法
*/
import "fmt"

// "func main" 是特殊的
//
// Go 需要程序从何启动
//
// main 函数创建了 Go 的启动点
//
// 编译完代码以后,
// Go 运行时将首先运行此方法
func main() {
	// import "fmt" 之后
	// "fmt" 包的 Println 方法变得可用

	// 试试看在控制台输出以下命令
	//   go doc -src fmt Println

	// Println 是从 "fmt" 包导出的

	// 首字母大写的等于导出的包
	fmt.Println("Hello Gopher!")

	// Go 自身不能调用 Printl 函数
	// 这就是你为何需要在此调用的原因
	// Go 只会自动的调用 `func main`
	// -----

	// 在字符字面量中 Go 支持 Unicode 编码
	// 源代码中也支持: KÖSTEBEK!
	//
	// 因为: 字面量 ~= 月代码

	// 练习： 去除以下注释 --> //
	// fmt.Println("Merhaba Köstebek!")

	// 不必要的注释:
	// "Merhaba Köstebek" 指 "Hello Gopher"
	// 在土耳其语中
}
