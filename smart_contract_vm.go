package main

import "errors"

type OpCode byte

const (
	OpAdd OpCode = iota
	OpSub
	OpMul
	OpStore
	OpLoad
)

type VM struct {
	Memory map[string]int
	Stack  []int
	Code   []OpCode
	Cursor int
}

func NewVM(code []OpCode) *VM {
	return &VM{
		Memory: make(map[string]int),
		Stack:  []int{},
		Code:   code,
		Cursor: 0,
	}
}

func (vm *VM) Run() error {
	for vm.Cursor < len(vm.Code) {
		op := vm.Code[vm.Cursor]
		switch op {
		case OpAdd:
			if len(vm.Stack) < 2 {
				return errors.New("stack underflow")
			}
			a := vm.Stack[len(vm.Stack)-1]
			b := vm.Stack[len(vm.Stack)-2]
			vm.Stack = vm.Stack[:len(vm.Stack)-2]
			vm.Stack = append(vm.Stack, a+b)
		case OpStore:
			key := "storage_" + string(vm.Cursor)
			val := vm.Stack[len(vm.Stack)-1]
			vm.Stack = vm.Stack[:len(vm.Stack)-1]
			vm.Memory[key] = val
		case OpLoad:
			key := "storage_" + string(vm.Cursor)
			vm.Stack = append(vm.Stack, vm.Memory[key])
		}
		vm.Cursor++
	}
	return nil
}
