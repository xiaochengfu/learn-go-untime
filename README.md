### 通过学习nsq源码来学习go

#### nsqd main源码解析

#### 第一个包 go-svc 可以与Linux配合使用的Windows Service包装器

> 知识充电：Linux 常用SIG信号及其键值

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

原文链接：https://blog.csdn.net/qq_38570571/article/details/79870441
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

> 知识补充： sync.once

可查看代码示例学习，路径 `sync.once/main.go`

> 知识补充： sync.waitGroup

可查看代码示例学习，路径 `sync.waitGroup/main.go`

