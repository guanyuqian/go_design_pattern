package _9_Memento_Pattern

import (
	"reflect"
	"testing"
)

//步骤 4
//使用 CareTaker 和 Originator 对象。
func TestMementoPattern(t *testing.T) {
	state := [...]string{"State #1", "State #2", "State #3", "State #4"}
	originator := Originator{}
	careTaker := CareTaker{}
	originator.setState(state[0])
	originator.setState(state[1])
	careTaker.add(originator.saveStateToMemento())
	originator.setState(state[2])
	careTaker.add(originator.saveStateToMemento())
	originator.setState(state[3])

	tests := [...]struct {
		name    string
		index   int
		wantRet string
	}{
		{"Current State", -1, state[3]},
		{"First saved State", 0, state[1]},
		{"Second saved State", 1, state[2]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originator.getStateFromMemento(careTaker.get(tt.index))
			if gotRet := originator.getState(); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("got = %v, want = %v", gotRet, tt.wantRet)
			}
		})
	}
}
