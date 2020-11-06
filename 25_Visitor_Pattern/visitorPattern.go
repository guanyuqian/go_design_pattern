package _5_Visitor_Pattern

//步骤 1
//定义一个表示元素的接口。

type ComputerPart interface {
	accept(computerPartVisitor ComputerPartVisitor) string
}

//步骤 2
//创建扩展了上述类的实体类。
type Keyboard struct {
}

func (receiver *Keyboard) accept(computerPartVisitor ComputerPartVisitor) string {
	return computerPartVisitor.visit(receiver)
}

type Monitor struct {
}

func (receiver *Monitor) accept(computerPartVisitor ComputerPartVisitor) string {
	return computerPartVisitor.visit(receiver)
}

type Mouse struct {
}

func (receiver *Mouse) accept(computerPartVisitor ComputerPartVisitor) string {
	return computerPartVisitor.visit(receiver)
}

type Computer struct {
	parts []ComputerPart
}

func (receiver *Computer) accept(computerPartVisitor ComputerPartVisitor) string {
	ret := ""
	for _, part := range receiver.parts {
		ret += part.accept(computerPartVisitor)
	}
	ret += computerPartVisitor.visit(receiver)
	return ret
}

func newComputer() *Computer {
	return &Computer{[]ComputerPart{&Mouse{}, &Monitor{}, &Keyboard{}}}
}

//步骤 3
//定义一个表示访问者的接口。
type ComputerPartVisitor interface {
	visit(ComputerPart) string
}

//步骤 4
//创建实现了上述类的实体访问者。
type ComputerPartDisplayVisitor struct {
}

func (receiver *ComputerPartDisplayVisitor) visit(computerPart ComputerPart) (ret string) {
	switch computerPart.(type) {
	case *Mouse:
		ret = "Displaying Mouse."
	case *Monitor:
		ret = "Displaying Monitor."
	case *Keyboard:
		ret = "Displaying Keyboard."
	case *Computer:
		ret = "Displaying Computer."
	}
	return
}
