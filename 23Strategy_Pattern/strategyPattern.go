package _3Strategy_Pattern

//步骤 1
//创建一个接口。
type Strategy interface {
	doOperation(num1, num2 int) int
}

//步骤 2
//创建实现接口的实体类。
type OperationAdd struct{}

func (receiver *OperationAdd) doOperation(num1, num2 int) int {
	return num1 + num2
}

type OperationSubtract struct{}

func (receiver *OperationSubtract) doOperation(num1, num2 int) int {
	return num1 - num2
}

type OperationMultiply struct{}

func (receiver *OperationMultiply) doOperation(num1, num2 int) int {
	return num1 * num2
}

//步骤 3
//创建 Context 类。

type Context struct {
	strategy Strategy
}

func (receiver Context) executeStrategy(num1, num2 int) int {
	return receiver.strategy.doOperation(num1, num2)
}
