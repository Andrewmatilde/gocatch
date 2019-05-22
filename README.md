# gocatch

gocatch 是一个生产者消费者模型的并发框架，可以提供给爬虫等项目使用.gocatch依赖于goroutine与sync,time库.

用户需要实现：
	type WorkMachine interface {
		Work(u *StrStack) ResPipe
	}

	type AnalyzeMachine interface {
		Analyze(u *StrStack, er ResPipe)
	}
这两个接口，

StrStack实现了一个字符串栈

ResPipe是WorkMachine与AnalyzeMachine之间的传输中间件，
使用
	func (e ResPipe) GetResValueInterface() interface{}
可以取出WorkMachine放入ResPipe中的Res interface{}
由于go的反射机制非常恶心，取出后需要强制类型转换才能使用.

实现了WorkMachine与AnalyzeMachine接口后
声明一个manager：
	type Manager struct {
		WorkLineNum int
		AnaLineNum  int
		Worker      WorkMachine
		Analyst     AnalyzeMachine
		Stack       *StrStack
	}

这将会根据WorkLineNum,AnaLineNum,开启work与analyze线程
但是所有完成之前，由于希望用户确定线程等待的时间
将让用户创建WorkLine与AnaLine的列表

	func CreatWorkLineList(m Manager, BreakTime time.Duration) []WorkLine
	func CreateAnaLineList(m Manager, BreakTime time.Duration) []AnaLine
最后通过

	func RunManager(m Manager, WorkLineList []WorkLine, AnaLineList []AnaLine)

启动
所有线程会在所有任务完成后关闭.
