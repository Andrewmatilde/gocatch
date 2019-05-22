package gocatch

import "sync"

type StrStack struct {
	Strings []string
	mu      sync.Mutex
}

func (us *StrStack) Empty() bool {
	return len(us.Strings) == 0
}

func (us *StrStack) PushO(s string) {
	us.Strings = append(us.Strings, s)
}

func (us *StrStack) PushA(s []string) {
	for _, v := range s {
		us.Strings = append(us.Strings, v)
	}
}

func (us *StrStack) SafePushO(s string) {
	us.mu.Lock()
	us.PushO(s)
	defer us.mu.Unlock()
}

func (us *StrStack) SafePushA(s []string) {
	us.mu.Lock()
	us.PushA(s)
	defer us.mu.Unlock()
}

func (us *StrStack) PopO() string {
	s := us.Strings[0]
	us.Strings = us.Strings[1:]
	return s
}

func (us *StrStack) PopA(size int) []string {
	s := us.Strings[:size]
	us.Strings = us.Strings[size:]
	return s
}

func (us *StrStack) SafePopO() string {
	us.mu.Lock()
	defer us.mu.Unlock()
	return us.PopO()
}

func (us *StrStack) SafePopA(size int) []string {
	us.mu.Lock()
	defer us.mu.Unlock()
	return us.PopA(size)
}

func (us *StrStack) CreatUrlStack(s []string) StrStack {
	return StrStack{s, sync.Mutex{}}
}
