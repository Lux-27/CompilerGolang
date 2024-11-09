package ast

type ModifierFunc func(Node) Node

func Modify(node Node, modfier ModifierFunc) Node {
	switch node := node.(type) {
	case *Program:
		for i, statement := range node.Statements {
			node.Statements[i], _ = Modify(statement, modfier).(Statement)
		}

	case *ExpressionStatement:
		node.Expression, _ = Modify(node.Expression, modfier).(Expression)

	case *InfixExpression:
		node.Left, _ = Modify(node.Left, modfier).(Expression)
		node.Right, _ = Modify(node.Right, modfier).(Expression)

	case *PrefixExpression:
		node.Right, _ = Modify(node.Right, modfier).(Expression)

	case *IndexExpression:
		node.Left, _ = Modify(node.Left, modfier).(Expression)
		node.Index, _ = Modify(node.Index, modfier).(Expression)

	case *IfExpression:
		node.Condition, _ = Modify(node.Condition, modfier).(Expression)
		node.Consequence, _ = Modify(node.Consequence, modfier).(*BlockStatement)
		if node.Alternative != nil {
			node.Alternative, _ = Modify(node.Alternative, modfier).(*BlockStatement)
		}

	case *BlockStatement:
		for i := range node.Statements {
			node.Statements[i], _ = Modify(node.Statements[i], modfier).(Statement)
		}

	case *ReturnStatement:
		node.ReturnValue, _ = Modify(node.ReturnValue, modfier).(Expression)

	case *LetStatement:
		node.Value, _ = Modify(node.Value, modfier).(Expression)

	case *FunctionLiteral:
		for i := range node.Parameters {
			node.Parameters[i], _ = Modify(node.Parameters[i], modfier).(*Identifier)
		}
		node.Body, _ = Modify(node.Body, modfier).(*BlockStatement)

	case *ArrayLiteral:
		for i := range node.Elements {
			node.Elements[i], _ = Modify(node.Elements[i], modfier).(Expression)
		}

	case *HashLiteral:
		newPairs := make(map[Expression]Expression)
		for key, val := range node.Pairs {
			newKey, _ := Modify(key, modfier).(Expression)
			newVal, _ := Modify(val, modfier).(Expression)

			newPairs[newKey] = newVal
		}

		node.Pairs = newPairs

	}

	return modfier(node)
}
