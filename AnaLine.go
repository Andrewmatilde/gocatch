package gocatch

import "time"

type AnalyzeMachine interface {
	Analyze(u *StrStack, er ResPipe)
}

type AnaLine struct {
	BreakTime   time.Duration
	IsConsuming bool
	IfDel       bool
}

func CreateAnaLineList(m Manager, BreakTime time.Duration, IsConsuming bool, IfDel bool) []AnaLine {
	var AnaLineList []AnaLine
	for i := 0; i < m.AnaLineNum; i++ {
		AnaLineList = append(AnaLineList, AnaLine{BreakTime, IsConsuming, IfDel})
	}
	return AnaLineList
}

func (a *AnaLine) Analysis(
	as AnalyzeMachine,
	u *StrStack,
	cr chan ResPipe) {
	er := <-cr
	as.Analyze(u, er)
}

func (a *AnaLine) RunAnaLine(
	as AnalyzeMachine,
	u *StrStack,
	cr chan ResPipe) {
	for {
		if len(cr) != 0 {
			a.IsConsuming = true
			a.Analysis(as, u, cr)
			a.IsConsuming = false
		} else {
			time.Sleep(a.BreakTime)
		}
		if a.IfDel {
			break
		}
	}
}
