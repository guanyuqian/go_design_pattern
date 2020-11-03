package _8_Mediator_Pattern

//步骤 1
//创建中介类。

type ChatRoom struct {
	roomName string
}

func (receiver ChatRoom) showMessage(user *User, message *string) string {
	return receiver.roomName + " [" + user.name + "]: " + *message
}

//步骤 2
//创建 user 类。

type User struct {
	ChatRoom
	name string
}

func (receiver *User) sendMessage(message string) string {
	return receiver.showMessage(receiver, &message)
}
