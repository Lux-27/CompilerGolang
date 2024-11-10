package vm

import (
	"CompilerGolang/code"
	"CompilerGolang/compiler"
	"CompilerGolang/object"
	"fmt"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions

	stack        []object.Object
	stackPointer int // Always points to the next free value, Top of stack is (stackPointer - 1)
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		constants:    bytecode.Constants,
		instructions: bytecode.Instructions,

		stack:        make([]object.Object, StackSize),
		stackPointer: 0,
	}
}

// gives the top element of stack
func (vm *VM) StackTop() object.Object {
	if vm.stackPointer == 0 {
		return nil
	}

	return vm.stack[vm.stackPointer-1]
}

func (vm *VM) Run() error {
	for instructionPointer := 0; instructionPointer < len(vm.instructions); instructionPointer++ {
		// turn first byte of instruction into opcode
		// we do not use Lookup here as it is too slow
		//  It costs time to move the byte around, lookup the opcode’s definition, return it and take it apart.
		opCode := code.Opcode(vm.instructions[instructionPointer])

		switch opCode {
		case code.OpConstant:
			// decode operands in the bytecode, starting with byte positioned right after instructionPointer
			constIndex := code.ReadUint16(vm.instructions[instructionPointer+1:])
			// increment instructionPointer by 2 (2 is the size of each operand in OpConstant - SIZE OF INT)
			instructionPointer += 2

			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}

		case code.OpAdd:
			right := vm.pop()
			left := vm.pop()
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			result := leftValue + rightValue
			vm.push(&object.Integer{Value: result})
		}
	}

	return nil
}

func (vm *VM) push(obj object.Object) error {
	if vm.stackPointer >= StackSize {
		return fmt.Errorf("STACK OVERFLOW")
	}

	vm.stack[vm.stackPointer] = obj
	vm.stackPointer++

	return nil
}

func (vm *VM) pop() object.Object {
	obj := vm.stack[vm.stackPointer-1]
	vm.stackPointer--

	return obj
}
