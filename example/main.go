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
package main

import "github.com/mintzhao/maxStack"

func main() {
	stack := maxStack.NewMaxStack()
	stack.Push(4)
	stack.Push(5)

	if max, _ := stack.Max(); max != 5 {
		panic("max must be 5")
	}

	stack.Push(8)
	stack.Push(8)
	if max, _ := stack.Max(); max != 8 {
		panic("max must be 8")
	}
}