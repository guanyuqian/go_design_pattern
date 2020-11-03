package _0_Observer_Pattern

import "strconv"

//步骤 1
//创建 Subject 类。
type Subject struct {
	obervers []Observer
	state    int64
}

func (receiver *Subject) getState() int64 {
	return receiver.state
}

func (receiver *Subject) setState(state int64) string {
	receiver.state = state
	return receiver.notifyAllObservers()
}

func (receiver *Subject) attach(observer Observer) {
	receiver.obervers = append(receiver.obervers, observer)
}

func (receiver *Subject) notifyAllObservers() (ret string) {
	for _, observer := range receiver.obervers {
		ret += observer.update()
	}
	return
}

//步骤 2
//创建 Observer 类。

type Observer interface {
	update() string
}

//步骤 3
//创建实体观察者类。

type BinaryObserver struct {
	*Subject
}

func (receiver *BinaryObserver) update() string {
	return "Binary String: " + strconv.FormatInt(receiver.getState(), 2) + "."
}

type OctalObserver struct {
	*Subject
}

func (receiver *OctalObserver) update() string {
	return "Octal String: " + strconv.FormatInt(receiver.getState(), 8) + "."
}

type HexObserver struct {
	*Subject
}

func (receiver *HexObserver) update() string {
	return "Hex String: " + strconv.FormatInt(receiver.getState(), 16) + "."
}
