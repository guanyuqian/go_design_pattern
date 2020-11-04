package _1_State_Pattern

import "testing"

//步骤 4
//使用 Context 来查看当状态 State 改变时的行为变化。
func TestStatePattern(t *testing.T) {
	wants := [...]string{"Start State", "Stop State"}
	context := new(Context)
	startState := new(StartState)
	startState.doAction(context)
	if got := context.getState().toString(); got != wants[0] {
		t.Errorf("doAction() = %v, want %v", got, wants[0])
	}
	stopState := new(StopState)
	stopState.doAction(context)
	if got := context.getState().toString(); got != wants[0] {
		t.Errorf("doAction() = %v, want %v", got, wants[0])
	}
}
