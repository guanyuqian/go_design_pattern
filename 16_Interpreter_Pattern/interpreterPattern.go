package _6_Interpreter_Pattern

import "strings"

//步骤 1
//创建一个表达式接口。
type Expression interface {
	interpret(context string) bool
}

//步骤 2
//创建实现了上述接口的实体类。
type TerminalExpression struct {
	data string
}

func (receiver TerminalExpression) interpret(context string) bool {
	return strings.Contains(context, receiver.data)
}

type OrExpression struct {
	expr1, expr2 Expression
}

func (receiver OrExpression) interpret(context string) bool {
	return receiver.expr1.interpret(context) || receiver.expr2.interpret(context)
}

type AndExpression struct {
	expr1, expr2 Expression
}

func (receiver AndExpression) interpret(context string) bool {
	return receiver.expr1.interpret(context) && receiver.expr2.interpret(context)
}
