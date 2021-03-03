# 编程者路线图

你好!

如果你是一个有经验的开发者, 你可能只想看以下课程。 根据以下路线图, 你将需要看50课而不是180课。

这门课程从最基本的开始，一步一步地进行到最后。因此，每一个层次上的主题的复杂性都在增加。为了让每个人都方便，我特意设计了它。

如果你觉得有些题目对你来说很容易，那么就看复习课，做测验和练习，甚至连那部分的课都跳过，你以后随时可以再来观看课程。

学得开心!

## 课程

* **写下你第一个 Go 语言程序**
  * 请观看以下所有课程

* **掌握 Go 语言的类型系统**
  * Go Doc 是什么?
  * 如何写一个库包?
  * 每种 Go 类型都有零值
  * 空白标识符是什么?
  * 让我们来定义一对变量把!
  * 什么是类型推演?
  * 如何简洁的声明多个变量?
  * 为什么不能在包的层级使用简短声明?
  * 什么是重新声明?
  * 什么时候可以使用简单声明?
  * 从命令行中获取输入以及学习切片
  * 了解 os.Args 的基础知识
  * 使用 os.Args 问候人们
  * 如何使用 Printf 来打印格式化的输出、
  * 将摄氏度转换为华氏度
  * 将尺转换为米
  * 什么是原始文字字符串?
  * 怎么获取字符串的长度?
  * 什么是预先声明的类型以及段末总结.
  * 如何理解无类型常量
    * 学习常量的规则
    * 复习: 常量
    * 无类型常量到底是如何工作的?
    * 什么是默认类型?
    * 案例: time.Duration
    * 什么是 iota?
    * 命名规范

* **控制流和错误处理**
  * 观看关于 "Pass Me: 创建一个密保程序"的课程
  * 观看 "理解 Go 的错误处理"
  * 在 case 条件中使用多个值
  * 失败语句是怎么工作的?
  * 解决方案: 一天的各个部分
  * 复习: Switch 语句
  * 如何继续一个循环? (额外内容: 调试)
  * 创建乘法表
  * 如何遍历切片?
  * 使用简单的 For Range 方式!

* **初学者项目**
  * 请观看所有的课程.

* **剩余部分**
  * 你可以从此开始看剩下的所有课程。他们是中级到高级水平的课程。

## 到此为止! 学的开心! 🤩

---

# 额外内容: 为什么你要学习 Go 语言?

**总而言之:** Go 与 Python/Javascript 一样简单, 并和 C/C++ 一样快。比起 C/C++，用 Go 工作更享受。 你可以走低级路线，也可以保持高级路线。

## Go 语言被用于做什么?

Go主要用于网络公司：Google、Facebook、Twitter、Uber、Apple、Dropbox、Soundcloud、Medium、Mozilla Firefox、Github、Docker、Kubernetes和Heroku。

**Go语言的最佳使用场景:** 跨平台命令行工具, 分布式网络应用, 云技术如微服务和无服务, Web APIs, 数据库引擎, 大数据处理管道, 嵌入式开发等。

**Go语言的非最佳使用场景(但还行):** 桌面应用, 操作系统, 内核驱动, 游戏开发等。

## 谁设计了 Go ?

Go 是由产业界最有影响里的人之一设计的:

* Unix: Ken Thompson
* UTF-8, Plan 9: Rob Pike
* Hotspot JVM (Java Virtual Machine): Robert Griesemer

## 你能挣多少钱?

* [Go 语言薪水](https://www.payscale.com/research/US/Skill=Go_(Golang)_Programming_Language/Salary)

## [Go 发布8年以后](https://blog.golang.org/8years):

> 如今，**每一家云公司都在Go**中实现了其云基础设施的关键组件，包括Google cloud、AWS、microsoftazure、Heroku等。Go 是像阿里巴巴、Cloudflare和Dropbox这样的云公司的关键部分。Go 是开放基础设施的关键部分，包括Kubernetes、cloudfoundry、Openshift、NATS、Docker、Istio、Etcd、Consul、Juju等等。越来越多的公司选择Go来构建云基础设施解决方案。

## 你能用 Go 实现什么?

* [Go 实现的网络驱动](https://www.net.in.tum.de/fileadmin/bibtex/publications/theses/2018-ixy-go.pdf) (_与C驱动程序相比只有10%的性能差距_)
* [Google gVisor](https://cloud.google.com/blog/products/gcp/open-sourcing-gvisor-a-sandboxed-container-runtime) (_用Go编写的用户空间内核_)
* [多平台任天堂模拟器](https://humpheh.github.io/goboy/)
* [Docker: 容器系统](https://github.com/moby/moby)
* [Kubernetes: 容器调度和管理](https://github.com/kubernetes/kubernetes)
* VM映像重复数据消除工具
* 聊天服务器
* RUM 信标采集器
* 时间序列数据库引擎、客户机、命令行工具等。
* Map-reduce 库
* 集群化的前端优化反向代理，具有动态内容重写、图像大小调整、缓存、Lua事件处理程序执行（所有多租户）
* 地理分布的反向代理CDN节点
* 具有事件处理程序和对等报告的运行状况管理守护程序
* 纯Go DNS服务器
* 与MySQL交互的API后端
* Linux进程捕获/还原实用程序
* 隐藏我们的关键服务器的方向代理
* 用于生成发票的HTML->PDF转换器
* 类似tinyurl.com以及goo.gl的短链接系统
* 短信服务
* 信用卡支付网关
* JSON Web令牌包
* 实时图像处理服务
* 3d渲染场/内容生产流水线（相当大的项目）
* 生产lxc容器部署
* 自动化测试管理

参考: [这篇 Reddit 文章](https://www.reddit.com/r/golang/comments/5nac2b/what_have_you_used_go_for_in_your_professional/).

## 查看更多信息：

* [关于 Go ：概述](https://blog.learngoprogramming.com/about-go-language-an-overview-f0bee143597c)
* [你为什么要学 Go ？](https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65)
* [云基础设施的新兴语言](https://redmonk.com/dberkholz/2014/03/18/go-the-emerging-language-of-cloud-infrastructure/)
* [使用 Go 的公司](https://github.com/golang/go/wiki/GoUsers)
* [Go 的八年](https://blog.golang.org/8years)
* [推特：使用 Go 一天处理50亿个会话](https://blog.twitter.com/engineering/en_us/a/2015/handling-five-billion-sessions-a-day-in-real-time.html)
* [一名 C++ 开发者对 Go 的看法](https://www.murrayc.com/permalink/2017/06/26/a-c-developer-looks-at-go-the-programming-language-part-1-simple-features/)

<div style="page-break-after: always;"></div>

> 更多教程: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
> 
> Copyright © 2018 Inanc Gumus
> 
> 学习 Go 语言编程课程
> 
> [点击此查看许可证](https://creativecommons.org/licenses/by-nc-sa/4.0/)

