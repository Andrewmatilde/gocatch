package gocatch

import (
	"sync"
	"time"
)

type Manager struct {
	DlLineNum int
	AnaLineNum int
}

func (Eg Manager)RunEngine(
	s *Stack,
	dm DownMachine,
	as Analysis) {
	wg := sync.WaitGroup{}
	DLLineList := CreatDLLineList(Eg.DlLineNum)
	AnaLineList := CreateAnaLineList(Eg.AnaLineNum)
	cr := make(chan EleRes,Eg.DlLineNum*Eg.AnaLineNum*16)
	for i := 0; i <Eg.DlLineNum;i++ {
		wg.Add(1)
		go func(int0 int) {
			DLLineList[int0].RunStation(dm,s,cr)
			wg.Done()
		}(i)
	}
	for j:=0;j< Eg.AnaLineNum ;j++ {
		wg.Add(1)
		go func(int1 int) {
			AnaLineList[int1].RunStation(as,s,cr)
			wg.Done()
		}(j)
	}
	go func() {
		for {
			IfOneWorking := false
			for i := 0; i <Eg.DlLineNum;i++ {
				IfOneWorking = DLLineList[i].IsWorking||IfOneWorking
			}
			for j:=0;j< Eg.AnaLineNum ;j++ {
				IfOneWorking = AnaLineList[j].IsConsuming||IfOneWorking
			}
			if (IfOneWorking==false)&&s.Empty()&&(len(cr)==0) {
				for i := 0; i <Eg.DlLineNum;i++ {
					DLLineList[i].IfDel = true
				}
				for j:=0;j< Eg.AnaLineNum ;j++ {
					AnaLineList[j].IfDel = true
				}
				break
			} else {
				time.Sleep(time.Duration(Eg.DlLineNum*Eg.AnaLineNum*20)*time.Millisecond)
			}
		}
	}()
	wg.Wait()
}
