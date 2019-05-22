package gocatch

import "reflect"

type EleRes struct {
	Res interface {}
	Data string
}

func (e EleRes)GetResValueNeedChangeTypes() interface{} {
	return reflect.ValueOf(e.Res).Interface()
}
