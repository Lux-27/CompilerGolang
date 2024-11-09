package object

import (
	"CompilerGolang/ast"
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	QUOTE_OBJ        = "QUOTE"
	MACRO_OBJ        = "MACRO"
)

/*****************INTEGER OBJECT*****************/
// type Integer struct {
// 	Value int64
// }

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

/*****************BOOLEAN OBJECT*****************/
// type Boolean struct {
// 	Value bool
// }

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

/*****************NULL OBJECT*****************/
type Null struct{}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

func (n *Null) Inspect() string {
	return "null"
}

/*****************RETURN VALUE OBJECT*****************/
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

/*****************ERROR OBJECT*****************/
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "Error: " + e.Message }

/*****************FUNCTION OBJECT*****************/
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

/*****************STRING OBJECT*****************/
// type String struct {
// 	Value string
// }

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

/*****************BUILT IN FUNCTIONS*****************/
type BuiltinFunctions func(args ...Object) Object
type Builtin struct {
	Fn BuiltinFunctions
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "Built-in function" }

/*****************ARRAY OBJECT*****************/
type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

/*****************HASHKEY OBJECT*****************/
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// optimizing the performance of the HashKey() methods by caching their return values inside structs
type Boolean struct {
	Value   bool
	hashKey *HashKey
}

type Integer struct {
	Value   int64
	hashKey *HashKey
}

type String struct {
	Value   string
	hashKey *HashKey
}

func (b *Boolean) HashKey() HashKey {
	if b.hashKey != nil {
		return *b.hashKey
	}

	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}

	b.hashKey = &HashKey{Type: b.Type(), Value: value}
	return *b.hashKey
}

func (i *Integer) HashKey() HashKey {
	if i.hashKey != nil {
		return *i.hashKey
	}

	i.hashKey = &HashKey{Type: i.Type(), Value: uint64(i.Value)}
	return *i.hashKey
}

func (s *String) HashKey() HashKey {
	if s.hashKey != nil {
		return *s.hashKey
	}

	h := fnv.New64a()
	h.Write([]byte(s.Value))
	s.hashKey = &HashKey{Type: s.Type(), Value: h.Sum64()}
	return *s.hashKey
}

/*****************HASH OBJECT*****************/
type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }

func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

/*****************HASHABLE OBJECT*****************/
type Hashable interface {
	HashKey() HashKey
}

/*****************QUOTE OBJECT*****************/
type Quote struct {
	Node ast.Node
}

func (q *Quote) Type() ObjectType { return QUOTE_OBJ }
func (q *Quote) Inspect() string {
	return "QUOTE(" + q.Node.String() + ")"
}

/*****************MACRO OBJECT*****************/
type Macro struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (m *Macro) Type() ObjectType { return MACRO_OBJ }
func (m *Macro) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range m.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("macro")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(m.Body.String())
	out.WriteString("\n}")

	return out.String()
}
