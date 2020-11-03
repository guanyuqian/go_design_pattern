package _9_Memento_Pattern

//步骤 1
//创建 Memento 类。

type Memento struct {
	state string
}

//步骤 2
//创建 Originator 类。
type Originator struct {
	state string
}

func (receiver *Originator) setState(state string) {
	receiver.state = state
}

func (receiver *Originator) getState() string {
	return receiver.state
}

func (receiver *Originator) saveStateToMemento() *Memento {
	return &Memento{receiver.state}
}

func (receiver *Originator) getStateFromMemento(Memento *Memento) {
	if Memento != nil {
		receiver.state = Memento.state
	}
}

//步骤 3
//创建 CareTaker 类。
type CareTaker struct {
	mementoList []*Memento
}

func (receiver *CareTaker) add(state *Memento) {
	receiver.mementoList = append(receiver.mementoList, state)
}

func (receiver *CareTaker) get(index int) (ret *Memento) {
	if index >= 0 && index < len(receiver.mementoList) {
		ret = receiver.mementoList[index]
	}
	return
}
