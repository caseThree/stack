package stack

import (
	"errors"
	"sync"
)

type intStack struct {
	list []int
	mux sync.Mutex
}

type IStack interface {
	Empty() bool
	Add(int)
	Top() (int, error)
	Pop() (int, error)
	Size() int
	Reset()
}

func (st *intStack) Empty() bool {
	st.mux.Lock()
	defer st.mux.Unlock()
	return len(st.list) == 0
}

func (st *intStack) intEmpty() bool {
	return len(st.list) == 0
}


func (st *intStack) Add(val int) {
	st.mux.Lock()
	defer st.mux.Unlock()
	st.list = append(st.list, val)
}

func (st *intStack) Top() (int, error) {
	st.mux.Lock()
	defer st.mux.Unlock()
	if st.intEmpty() {
		return 0, errors.New("stack is empty. cannot top")
	}
	return st.list[len(st.list)-1], nil
}

func (st *intStack) Pop() (int, error) {
	st.mux.Lock()
	defer st.mux.Unlock()
	if st.intEmpty() {
		return 0, errors.New("stack is empty. cannot pop")
	}
	val := st.list[len(st.list)-1]
	st.list = st.list[:len(st.list)-1]
	return val, nil
}

func (st *intStack) Size() int {
	st.mux.Lock()
	defer st.mux.Unlock()
	return len(st.list)
}

func (st *intStack) Reset() {
	st.mux.Lock()
	defer st.mux.Unlock()
	st.list = make([]int, 0)
}

func Create() IStack {
	return &intStack{
		list: make([]int, 0),
		mux: sync.Mutex{},
	}
}