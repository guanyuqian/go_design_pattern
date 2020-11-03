package _8_Mediator_Pattern

import "testing"

//步骤 3
//使用 User 对象来显示他们之间的通信。
func TestMediatorPattern(t *testing.T) {
	chatRoom := ChatRoom{"room1"}
	robert := User{chatRoom, "Robert"}
	john := User{chatRoom, "John"}
	wants := [...]string{"room1 [Robert]: Hi! John!", "room1 [John]: Hello! Robert!"}
	if got := robert.sendMessage("Hi! John!"); got != wants[0] {
		t.Errorf("sendMessage() = %v, want %v", got, wants[0])
	}
	if got := john.sendMessage("Hello! Robert!"); got != wants[1] {
		t.Errorf("sendMessage() = %v, want %v", got, wants[1])

	}
}
