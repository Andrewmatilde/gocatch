package gocatch

import "sync"

type StrStack interface {
	Empty() bool
	PushO(s string)
	PushA(s []string)
	SafePushO(s string)
	SafePushA(s []string)
	PopO() string
	PopA(size int) []string
	SafePopO() string
	SafePopA(size int) []string
}

type Stack struct {
	Urls []string
	mu sync.Mutex
}

func (us *Stack) Empty() bool {
	return len(us.Urls) == 0
}

func (us *Stack) PushO(s string) {
	us.Urls = append(us.Urls,s)
}

func (us *Stack) PushA(s []string) {
	for _,v := range s {
		us.Urls = append(us.Urls,v)
	}
}

func (us *Stack) SafePushO(s string) {
	us.mu.Lock()
	us.PushO(s)
	defer us.mu.Unlock()
}

func (us *Stack) SafePushA(s []string) {
	us.mu.Lock()
	us.PushA(s)
	defer us.mu.Unlock()
}

func (us *Stack) PopO() string {
	s := us.Urls[0]
	us.Urls = us.Urls[1:]
	return s
}

func (us *Stack) PopA(size int) []string {
	s := us.Urls[:size]
	us.Urls = us.Urls[size:]
	return s
}

func (us *Stack) SafePopO() string {
	us.mu.Lock()
	defer us.mu.Unlock()
	return us.PopO()
}

func (us *Stack) SafePopA(size int) []string {
	us.mu.Lock()
	defer us.mu.Unlock()
	return us.PopA(size)
}

func (us *Stack) CreatUrlStack(s []string) Stack {
	return Stack{s,sync.Mutex{}}
}