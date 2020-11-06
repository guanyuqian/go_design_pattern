package _5_Visitor_Pattern

import "testing"

//步骤 5
//使用 ComputerPartDisplayVisitor 来显示 Computer 的组成部分。
func TestVisitorPattern(t *testing.T) {
	computer := newComputer()
	wantRet := "Displaying Mouse.Displaying Monitor.Displaying Keyboard.Displaying Computer."
	if gotRet := computer.accept(&ComputerPartDisplayVisitor{}); gotRet != wantRet {
		t.Errorf("visit() = %v, want %v", gotRet, wantRet)
	}
}
