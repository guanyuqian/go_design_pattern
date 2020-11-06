package _6_MVC_Pattern

import (
	"testing"
)

func retrieveStudentFromDatabase() *Student {
	return &Student{"Robert", "10"}
}

//步骤 4
//使用 StudentController 方法来演示 MVC 设计模式的用法。
func TestMVCPattern(t *testing.T) {
	//从数据库获取学生记录
	model := retrieveStudentFromDatabase()
	//创建一个视图：把学生详细信息输出到控制台
	view := &StudentView{}
	wants := [2]string{"Student: Name: RobertRoll No: 10", "Student: Name: JohnRoll No: 10"}
	controller := NewStudentController(model, view)

	if got := controller.updateView(); got != wants[0] {
		t.Errorf("printStudentDetails() = %v, want %v", got, wants[0])
	}
	//更新模型数据
	controller.setStudentName("John")
	if got := controller.updateView(); got != wants[1] {
		t.Errorf("printStudentDetails() = %v, want %v", got, wants[1])
	}
}
