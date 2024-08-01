package stack

import (
	"testing"
)

func TestCreate(test *testing.T) {
	var _ IStack = Create()
}

func TestEmpty(test *testing.T) {
	st := Create()
	if st.Empty() == false {
		test.Errorf("Size of stack is not empty")
	}
}

func TestAdd(test *testing.T) {
	st := Create()
	st.Add(5)
}

func TestNonEmpty(test *testing.T) {
	st := Create()
	st.Add(5)
	if st.Empty() == true {
		test.Errorf("Size of stack is non empty")
	}
}

func TestTop(test *testing.T) {
	st := Create()
	_, ok := st.Top()
	if ok == nil {
		test.Errorf("Empty stack top")
	}
}

func TestNonEmptyTop(test *testing.T) {
	st := Create()
	st.Add(4)
	st.Add(2)
	val, ok := st.Top()
	if ok != nil {
		test.Errorf("Non empty stack error %v", ok.Error())
	}
	if val != 2 {
		test.Errorf("Top value should be 2")
	}
}

func TestPop(test *testing.T) {
	st := Create()
	_, ok := st.Pop()
	if ok == nil {
		test.Errorf("Empty stack cannot pop")
	}
}

func TestSize(test *testing.T) {
	st := Create()
	st.Add(4)
	st.Add(6)
	st.Add(896)
	if st.Size() != 3 {
		test.Errorf("Stack size should be 3")
	}
}

func TestPopDecreaseSize(test *testing.T) {
	st := Create()
	for i:=1; i<=10; i++ {
		st.Add(i)
	}
	if st.Size() != 10 {
		test.Errorf("Stack size should be 10")
	}
	st.Pop()
	if st.Size() != 9 {
		test.Errorf("Stack size should be 9")
	}
	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()
	if st.Size() != 5 {
		test.Errorf("Stack size should be 5")
	}
}

func TestNonEmptyPop(test *testing.T) {
	st := Create()
	st.Add(2)
	st.Add(-65)
	val, ok := st.Pop()
	if ok != nil {
		test.Errorf("Non empty stack pop error")
	}
	if val != -65 {
		test.Errorf("Top value should be -65")
	}
	if val, ok := st.Top(); ok != nil || val != 2 {
		test.Errorf("Stack top should be 2")
	}
	if st.Size() != 1 {
		test.Errorf("Stack size should br 1")
	}
}

func TestReset(test *testing.T) {
	st := Create()
	st.Add(5)
	st.Add(6)
	st.Add(8)
	if st.Empty() == true {
		test.Errorf("Stack should not be empty")
	}
	st.Reset()
	if st.Empty() == false {
		test.Errorf("Stack should be empty")
	}
}