package compiler

import (
	"CompilerGolang/ast"
	"CompilerGolang/code"
	"CompilerGolang/lexer"
	"CompilerGolang/object"
	"CompilerGolang/parser"
	"fmt"
	"testing"
)

type compilerTestCase struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithematic(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "1 + 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpAdd),
			},
		},
	}
	runCompilerTests(t, tests)
}

func runCompilerTests(t *testing.T, tests []compilerTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)

		compiler := New()
		err := compiler.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		bytecode := compiler.Bytecode()

		err = testInstructions(tt.expectedInstructions, bytecode.Instructions)
		if err != nil {
			t.Fatalf("testInstructions failed: %s", err)
		}

		err = testConstants(t, tt.expectedConstants, bytecode.Constants)
		if err != nil {
			t.Fatalf("testConstants failed: %s", err)
		}
	}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testInstructions(expectedInstructions []code.Instructions, actualInstructions code.Instructions) error {
	// We need concatInstructions because the expectedInstructions field
	// in compilerTestCase is not just a slice of bytes, but a slice of slices of bytes.
	concatted := concatInstructions(expectedInstructions)

	if len(actualInstructions) != len(concatted) {
		return fmt.Errorf("wrong instructions length.\nwant=%q\ngot =%q", concatted.MiniDisassembler(), actualInstructions.MiniDisassembler())
	}

	for i, ins := range concatted {
		if actualInstructions[i] != ins {
			return fmt.Errorf("wrong instruction at %d.\nwant=%q\ngot =%q", i, concatted.MiniDisassembler(), actualInstructions.MiniDisassembler())
		}
	}

	return nil
}

func concatInstructions(expectedInstructions []code.Instructions) code.Instructions {
	out := code.Instructions{}

	for _, ins := range expectedInstructions {
		out = append(out, ins...)
	}

	return out
}

func testConstants(t *testing.T, expectedConstants []interface{}, actualConstants []object.Object) error {
	if len(expectedConstants) != len(actualConstants) {
		return fmt.Errorf("wrong number of constants. got=%d, want=%d", len(actualConstants), len(expectedConstants))
	}

	for i, constant := range expectedConstants {
		switch constant := constant.(type) {
		case int:
			err := testIntegerObject(int64(constant), actualConstants[i])
			if err != nil {
				return fmt.Errorf("constant %d- testIntegerObject failed: %s", i, err)
			}

		}
	}

	return nil
}

func testIntegerObject(expectedConstant int64, actualConstant object.Object) error {
	result, ok := actualConstant.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)", actualConstant, actualConstant)
	}

	if result.Value != expectedConstant {
		return fmt.Errorf("object has wrong value, got = %d, expected = %d", result.Value, expectedConstant)
	}

	return nil
}
