package gocatch

import "reflect"

type ResPipe struct {
	Res interface {}
	Data string
}

func (e ResPipe) GetResValueInterface() interface{} {
	return reflect.ValueOf(e.Res).Interface()
}
