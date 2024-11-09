# CompilerGolang

**CompilerGolang** is a simple interpreter written in Go. It supports basic arithmetic operations, variable assignments, functions, arrays, and built-in functions, making it a great resource for understanding the fundamentals of building an interpreter.

## Features

- **Arithmetic Operations**: `+`, `-`, `*`, `/`
- **Variable Assignments**
- **Functions** and **Function Calls**
- **Arrays** and **Array Operations**
- **Built-in Functions**:
  - `len`: Returns the length of an array or string.
  - `first`: Returns the first element of an array.
  - `rest`: Returns the array without the first element.
  - `push`: Adds an element to the end of an array.
  - `puts`: Prints arguments to the console.
  - `update`: Updates an element in an array at a specified index.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/Lux-27/Interpreter_Golang
   cd CompilerGolang
   ```

2. Run the project:
   ```sh
   go run main.go
   ```
3. OR Build the project:
   ```sh
   go build
   ```
4. Run the interpreter:
   ```sh
   ./CompilerGolang
   ```

### Usage

You can use the interpreter to evaluate expressions, define functions, and work with arrays. Here are some examples:

#### Arithmetic Operations

```sh
>> 3 + 4
7
>> 10 - 2 * 5
0
```

#### Variable Assignments

```sh
>> let x = 10;
>> x
10
```

#### Functions

```sh
>> let add = fn(a, b) { a + b };
>> add(2, 3)
5
```

#### Nested Functions

```sh
 >> let makeGreeter = fn(greeting) { fn(name) { greeting + " " + name + "!" } };
 >> let hello = makeGreeter("Hello");
 >> hello("Luxy");
 Hello Luxy!
 >> let heythere = makeGreeter("Hey there");
 >> heythere("Luxy");
 Hey there Luxy!
```

#### Arrays

```sh
>> let arr = [1, 2, 3, 4];
>> len(arr)
4
>> first(arr)
1
>> rest(arr)
[2, 3, 4]
>> push(arr, 5)
[1, 2, 3, 4, 5]
>> update(arr, 1, 10)
[1, 10, 3, 4]
```

## Project Structure

- lexer/: Contains the lexer implementation.
- parser/: Contains the parser implementation.
- evaluator/: Contains the evaluator implementation and built-in functions.
- object/: Contains the object definitions and hash key implementations.
- repl/: Contains the REPL (Read-Eval-Print Loop) implementation.
- token/: Contains the token definitions.

## Acknowledgments

Inspired by the book Writing an Interpreter in Go by Thorsten Ball.
