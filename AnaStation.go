package gocatch

import "time"

type Analysis interface{
	Analyze(u *Stack,er EleRes)
}


type AnaStation struct {
	IsConsuming bool
	IfDel bool
}

func CreateAnaLineList(num int) []AnaStation {
	var AnaLineList []AnaStation
	for i := 0;i < num;i++ {
		AnaLineList = append(AnaLineList,AnaStation{false,false})
	}
	return AnaLineList
}

func (a *AnaStation)AnaLine(
	as Analysis,
	u *Stack,
	cr chan EleRes) {
	er := <- cr
	as.Analyze(u,er)
}

func (a *AnaStation)RunStation(
	as Analysis,
	u *Stack,
	cr chan EleRes) {
	for {
		if len(cr) != 0 {
			a.IsConsuming = true
			a.AnaLine(as,u,cr)
			a.IsConsuming = false
		} else {
			time.Sleep(time.Duration(20)*time.Millisecond)
		}
		if a.IfDel {
			break
		}
	}
}
