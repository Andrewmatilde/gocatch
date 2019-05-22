package gocatch

import (
	"sync"
	"time"
)

type Manager struct {
	WorkLineNum int
	AnaLineNum  int
	Worker      WorkMachine
	Analyst     AnalyzeMachine
	Stack       *StrStack
}

func (m *Manager)CreateResPipe() chan ResPipe {
	return  make(chan ResPipe, m.WorkLineNum*m.AnaLineNum*16)
}

func (m *Manager)GoWorkLine(wg *sync.WaitGroup,WorkLineList []WorkLine,cr chan ResPipe) {
	for i := 0; i < m.WorkLineNum;i++ {
		wg.Add(1)
		go func(int0 int) {
			WorkLineList[int0].RunWorkLine(m.Worker,m.Stack,cr)
			wg.Done()
		}(i)
	}
}

func (m *Manager)GoAnaLine(wg *sync.WaitGroup,AnaLineList []AnaLine,cr chan ResPipe) {
	for i := 0; i < m.AnaLineNum;i++ {
		wg.Add(1)
		go func(int0 int) {
			AnaLineList[int0].RunAnaLine(m.Analyst,m.Stack,cr)
			wg.Done()
		}(i)
	}
}

func RunEngine(m Manager,WorkLineList []WorkLine,AnaLineList []AnaLine) {
	wg := sync.WaitGroup{}
	cr := m.CreateResPipe()
	m.GoWorkLine(&wg,WorkLineList,cr)
	m.GoAnaLine(&wg,AnaLineList,cr)
	go func() {
		for {
			IfOneWorking := false
			for i := 0; i < m.WorkLineNum;i++ {
				IfOneWorking = WorkLineList[i].IsWorking||IfOneWorking
			}
			for j:=0;j< m.AnaLineNum ;j++ {
				IfOneWorking = AnaLineList[j].IsConsuming||IfOneWorking
			}
			if (IfOneWorking==false)&& m.Stack.Empty()&&(len(cr)==0) {
				for i := 0; i < m.WorkLineNum;i++ {
					WorkLineList[i].IfDel = true
				}
				for j:=0;j< m.AnaLineNum ;j++ {
					AnaLineList[j].IfDel = true
				}
				break
			} else {
				time.Sleep(time.Duration(m.WorkLineNum*m.AnaLineNum*100)*time.Millisecond)
			}
		}
	}()
	wg.Wait()
}
