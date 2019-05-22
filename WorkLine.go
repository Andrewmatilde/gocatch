package gocatch

import "time"

type WorkMachine interface {
	Work(u *StrStack) ResPipe
}

type WorkLine struct {
	BreakTime time.Duration
	IsWorking bool
	IfDel     bool
}

func CreatWorkLineList(m Manager, BreakTime time.Duration) []WorkLine {
	var DLLineList []WorkLine
	for i := 0; i < m.WorkLineNum; i++ {
		DLLineList = append(DLLineList, WorkLine{BreakTime, false, false})
	}
	return DLLineList
}

func (d *WorkLine) Worker(
	dm WorkMachine,
	u *StrStack,
	cr chan ResPipe) {
	ER := dm.Work(u)
	cr <- ER
}

func (d *WorkLine) RunWorkLine(
	dm WorkMachine,
	u *StrStack,
	cr chan ResPipe) {
	for {
		if !u.Empty() {
			d.IsWorking = true
			d.Worker(dm, u, cr)
			d.IsWorking = false
		} else {
			time.Sleep(d.BreakTime)
		}
		if d.IfDel {
			break
		}
	}
}
