package stack

import "errors"

type Stack struct {
	arr         []byte
	initialised bool
}

func (this *Stack) New(size int) error {
	if this.initialised == true {
		return errors.New("Called \"New()\" on already initialized Stack")
	}
	this.arr = make([]byte, 0, 20)
	this.initialised = true
	return nil
}

func (this *Stack) Top() (byte, error) {
	if this.initialised == false {
		return 0, errors.New("Called \"Top()\" on not initialized Stack")
	}
	length := len(this.arr)
	if length == 0 {
		return 0, errors.New("Called \"Top()\" on empty Stack")
	}
	return this.arr[length-1], nil
}

func (this *Stack) Pop() error {
	if this.initialised == false {
		return errors.New("Called \"Pop()\" on not initialized Stack")
	}
	length := len(this.arr)
	if length == 0 {
		return errors.New("Called \"Pop()\" on empty Stack")
	}
	this.arr = this.arr[0 : length-1]
	return nil
}

func (this *Stack) Push(item byte) error {
	if this.initialised == false {
		return errors.New("Called \"Push()\" on not initialized Stack")
	}
	length := len(this.arr)
	if (length + 1) > cap(this.arr) {
		this.arr = increaseCap(this.arr)
	}
	this.arr = this.arr[0 : length+1]
	this.arr[length] = item
	return nil
}

func (this *Stack) Len() int {
	return len(this.arr)
}

func increaseCap(arr []byte) []byte {
	newSize := cap(arr) * 2
	newArr := make([]byte, len(arr), newSize)
	for i := range arr {
		newArr[i] = arr[i]
	}
	return newArr
}
