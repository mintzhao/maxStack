/*
Copyright Mojing Inc. 2016 All Rights Reserved.
Written by mint.zhao.chiu@gmail.com. github.com: https://www.github.com/mintzhao

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package maxStack

import (
	"errors"
)

var ErrEmpty = errors.New("empty stack")

type MaxStack struct {
	stack []int
	max   []int
}

func NewMaxStack() *MaxStack {
	return &MaxStack{
		stack: make([]int, 0),
		max:   make([]int, 0),
	}
}

func (ms *MaxStack) Push(v int) {
	curmax, err := ms.Max()
	if err != nil || v >= curmax {
		ms.max = append(ms.max, v)
	}

	ms.stack = append(ms.stack, v)
}

func (ms *MaxStack) Pop() (int, error) {
	if ms.Empty() {
		return 0, ErrEmpty
	}

	stackLen := len(ms.stack)
	v := ms.stack[stackLen-1]
	ms.stack = ms.stack[:stackLen-1]

	curmax, err := ms.Max()
	if err == nil && curmax == v {
		ms.max = ms.max[:len(ms.max)-1]
	}

	return v, nil
}

func (ms *MaxStack) Top() (int, error) {
	if ms.Empty() {
		return 0, ErrEmpty
	}

	return ms.stack[len(ms.stack)-1], nil
}

func (ms *MaxStack) Empty() bool {
	return len(ms.stack) == 0
}

func (ms *MaxStack) Max() (int, error) {
	if ms.Empty() {
		return 0, ErrEmpty
	}

	return ms.max[len(ms.max)-1], nil
}
