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

import "testing"

func TestMaxStack_Empty(t *testing.T) {
	s := NewMaxStack()
	s.Push(1)

	if s.Empty() {
		t.Fatal("stack isn't empty")
	}
}

func TestMaxStack_Top(t *testing.T) {
	s := NewMaxStack()
	s.Push(1)

	if v, err := s.Top(); v != 1 || err != nil {
		t.Fatal("s.Top() must be 1")
	}
}

func TestMaxStack_Push(t *testing.T) {
	s := NewMaxStack()
	s.Push(1)
	s.Push(2)
}

func TestMaxStack_Pop(t *testing.T) {
	s := NewMaxStack()
	s.Push(1)
	s.Push(2)

	if v, err := s.Pop(); v != 2 || err != nil {
		t.Fatal("s.Pop() wrong")
	}
}

func TestMaxStack_Max(t *testing.T) {
	s := NewMaxStack()
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)

	for i := 0; i< 4; i++ {
		if v, err := s.Max(); v != 1 || err != nil {
			t.Fatal("s.Max() must be 1")
		}

		s.Pop()
	}
}