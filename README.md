### go学习实践

##### 1.  知识补充： [sync.once](/sync.once/main.go)

##### 2. 知识补充： [sync.waitGroup](/sync.waitGroup/main.go)

##### 3. 知识补充： [func.receiver](/func.receiver/main.go) 
###### 成员函数，相当于类的方法，作用在receiver上

##### 4. 知识充电：Linux 常用SIG信号及其键值 [原文链接](https://blog.csdn.net/qq_38570571/article/details/79870441 "SIG信号") 
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

##### 5. 知识补充：go mod
###### 上次学习go时，还没有go mod，go 1.11之后，加入了包依赖，跟php的composer类似，当然要赶紧get✅

```bash
#首先修改go proxy 改为七牛云镜像
go env -w GOPROXY=https://goproxy.cn,direct

#打开go mod 支持
go env -w  GO111MODULE=on

#1.生成go.mod文件

go mod init //当前目录

#2.整理项目中缺失的包

go mod tidy

#3.加载依赖包，自动归档到vendor目录

go mod vendor -v

#4. 使用goland开发时，开启go fmt和goimports 工具，自动格式化和引入包
setting -> tools -> file watchers

```
##### 6. 知识补充： [struct.interface](/struct.interface/main.go) 
###### 利用接口实现的一种编程模式，用于程序的解耦

##### 7.知识补充：[struct.interface.check](/struct.interface.check/main.go)
###### 接口完整性检查 `var _ 接口名称 = (*结构体名称)(nil)`

##### 8. 知识补充： [struct.nest](/struct.nest/main.go) 
###### 结构体内嵌，也可实现解耦

##### 9. 知识补充：[time](/time/main.go)
###### 常用的时间获取和格式转换

##### 10. 知识补充： [performance.test](/performance.test)
###### 1. 字符串拼接推荐方法与性能测试对比 [string_joint_test.go](/performance.test/string_joint_test.go)
###### 2. int转字符推荐方法与性能测试对比 [string_transform_test.go](/performance.test/string_transform_test.go)

##### 11. 知识补充：[func.options](/func.options)
###### 使用高阶函数来实现配置选项的灵活定义

##### 12. 知识补充：[func.channel](/func.channel)
###### 通过chan来实现类似管道式的函数

##### 13. 知识补充：[genny.example](/genny.example)
###### 1. go generate的用例
###### 2. 通过genny第三方包实现go类似c++template的效果，来解决泛型需要类型检查的问题,`go get github.com/cheekybits/genny`

##### 14. 知识补充：[interface.reflect](/interface.reflect/main.go)
###### reflect的常用用法

##### 15. 知识补充：[interface.type](/interface.type/main.go)
###### interface的类型转换例子

##### 16. 知识补充：[pointer](/pointer/main.go)
###### 值类型和引用类型的区别，&和*的区别

##### 17. 知识补充：[build](/notic/build)
###### 通过编译时的报错来加深go的基础语法知识

##### 18. 知识补充：[channel](/chan.rand)
###### chan的用法，实现一个随机数的生成和打印

##### 19. 知识补充：[zookeeper](/zookeeper.lock)
###### 基于zookeeper实现的分布式锁的示例

##### 20. 知识补充：[grpc](/grpc-example)
###### 1. 根据proto分成grpc go的代码
###### 2. 实现server/client的ssl+token鉴权通讯
###### 3. 实例参考 [go语言中文网](http://www.topgoer.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/gRPC/)

##### 21. 知识补充：[sync.map](/sync.map)
###### 使用sync.map 实现并发安全