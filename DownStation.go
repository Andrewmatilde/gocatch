package gocatch

import "time"

type DownMachine interface {
	Down(u *Stack) EleRes
}

type DownStation struct {
	IsWorking bool
	IfDel     bool
}

func CreatDLLineList(num int) []DownStation {
	var DLLineList []DownStation
	for i := 0;i < num;i++ {
		DLLineList = append(DLLineList,DownStation{false,false})
	}
	return DLLineList
}

func (d *DownStation) DownLine(
	dm DownMachine,
	u *Stack,
	cr chan EleRes) {
	ER := dm.Down(u)
	cr <- ER
}

func (d *DownStation) RunStation(
	dm DownMachine,
	u *Stack,
	cr chan EleRes) {
	for {
		if !u.Empty() {
			d.IsWorking = true
			d.DownLine(dm,u,cr)
			d.IsWorking = false
		} else {
			time.Sleep(time.Duration(20)*time.Millisecond)
		}
		if d.IfDel {
			break
		}
	}
}
