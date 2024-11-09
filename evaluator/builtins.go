package evaluator

import (
	"CompilerGolang/object"
	"fmt"
	"strconv"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got = %d, want = 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}

			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}

			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},

	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got = %d, want = 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},

	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got = %d, want = 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},

	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments, got = %d, want = 1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},

	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments, got = %d, want = 2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newELements := make([]object.Object, length+1)
			copy(newELements, arr.Elements)
			newELements[length] = args[1]

			return &object.Array{Elements: newELements}
		},
	},

	"update": {
		// update(arr, indexToBeUpdated, updatedValue)
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments, got = %d, want = 3", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `update` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			index, err := strconv.Atoi(args[1].Inspect())
			if err != nil {
				return newError("argument to `update` must be an integer, got %s", args[1].Type())
			}

			if index < 0 || index >= length {
				return newError("update index out of range, length of array = %d, got = %d", length, index)
			}

			arr.Elements[index] = args[2]

			return arr
		},
	},

	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}

/*
Test-Driving Arrays
We now have array literals, the index operator and a few built-in functions to work with arrays.

With first, rest and push we can build a map function:
let map = fn(arr, f) {
	let iter = fn(arr, accumulated) {
 		if (len(arr) == 0) {
			accumulated
		} else {
			iter(rest(arr), push(accumulated, f(first(arr))));
		}
	};

	iter(arr, []);
};

And with map we can do things like this:
 >> let a = [1, 2, 3, 4];
 >> let double = fn(x) { x * 2 };
 >> map(a, double);
	[2, 4, 6, 8]

Based on the same built-in functions we can also define a reduce function:

let reduce = fn(arr, initial, f) {
	let iter = fn(arr, result) {
		if (len(arr) == 0) {
			result
 		} else {
			iter(rest(arr), f(result, first(arr)));
		}
	};
	iter(arr, initial);
};

And reduce, in turn, can be used to define a sum function:

let sum = fn(arr) {
	reduce(arr, 0, fn(initial, el) { initial + el });
};

And it works like a charm:
>> sum([1, 2, 3, 4, 5]);
15
*/
