package _1_State_Pattern

//步骤 1
//创建一个接口。
type State interface {
	doAction(context *Context)
	toString() string
}

//步骤 2
//创建实现接口的实体类。

type StartState struct {
}

func (receiver *StartState) doAction(context *Context) {
	context.setState(receiver)
}

func (receiver *StartState) toString() string {
	return "Start State"
}

type StopState struct {
}

func (receiver *StopState) doAction(context *Context) {
	context.setState(receiver)
}

func (receiver *StopState) toString() string {
	return "Start State"
}

//步骤 3
//创建 Context 类。

type Context struct {
	state State
}

func (receiver *Context) getState() State {
	return receiver.state
}

func (receiver *Context) setState(state State) {
	receiver.state = state
}
