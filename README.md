### go学习实践

> 知识补充： [sync.once](/sync.once/main.go)

查看对应目录下的源代码（下同）

> 知识补充： [sync.waitGroup](/sync.waitGroup/main.go)

> 知识补充： [func.receiver](/func.receiver/main.go) 

成员函数，相当于类的方法，作用在receiver上

> 知识充电：Linux 常用SIG信号及其键值

[原文链接](https://blog.csdn.net/qq_38570571/article/details/79870441 "SIG信号") 
```
01 SIGHUP 挂起（hangup）
02 SIGINT 中断，当用户从键盘按^c键或^break键时
03 SIGQUIT 退出，当用户从键盘按quit键时
04 SIGILL 非法指令
05 SIGTRAP 跟踪陷阱（trace trap），启动进程，跟踪代码的执行
06 SIGIOT IOT指令
07 SIGEMT EMT指令
08 SIGFPE 浮点运算溢出
09 SIGKILL 杀死、终止进程
10 SIGBUS 总线错误
11 SIGSEGV 段违例（segmentation violation），进程试图去访问其虚地址空间以外的位置
12 SIGSYS 系统调用中参数错，如系统调用号非法
13 SIGPIPE 向某个非读管道中写入数据
14 SIGALRM 闹钟。当某进程希望在某时间后接收信号时发此信号
15 SIGTERM 软件终止（software termination）
16 SIGUSR1 用户自定义信号1
17 SIGUSR2 用户自定义信号2
18 SIGCLD 某个子进程死
19 SIGPWR 电源故障
```

> 知识补充：go mod

上次学习go时，还没有go mod，go 1.11之后，加入了包依赖，跟php的composer类似，当然要赶紧get✅

```bash
#首先修改go proxy 改为七牛云镜像
go env -w GOPROXY=https://goproxy.cn,direct

#1.生成go.mod文件

go mod init //当前目录

#2.整理项目中缺失的包

go mod tidy

#3.加载依赖包，自动归档到vendor目录

go mod vendor -v

```
> 知识补充： [struct.interface](/struct.interface/main.go) 

利用接口实现的一种编程模式，用于程序的解耦

> 知识补充：[struct.interface.check](/struct.interface.check/main.go)

接口完整性检查 `var _ 接口名称 = (*结构体名称)(nil)`

> 知识补充： [struct.nest](/struct.nest/main.go) 

结构体内嵌，也可实现解耦

> 知识补充：[time](/time/main.go)

常用的时间获取和格式转换

> 知识补充： [performance.test](/performance.test)

1. 字符串拼接推荐方法与性能测试对比 [string_joint_test.go](/performance.test/string_joint_test.go)

2. int转字符推荐方法与性能测试对比 [string_transform_test.go](/performance.test/string_transform_test.go)

> 知识补充：[func.options](/func.options)

使用高阶函数来实现配置选项的灵活定义

> 知识补充：[func.channel](/func.channel)

通过chan来实现类似管道式的函数