package _3_Transfer_Object_Pattern

import "testing"

func TestDataAccessObjectPattern(t *testing.T) {
	studentBO := NewStudentBO()
	gots := [2]string{}
	wants := [2]string{"Robert0John1", "Michael0"}

	for _, v := range studentBO.getAllStudents() {
		gots[0] += v.Name + v.RollNo
	}
	if gots[0] != wants[0] {
		t.Errorf("doTask() = %v, want %v", gots[0], wants[0])
	}

	//更新学生
	student := studentBO.getAllStudents()[0]
	student.Name = "Michael"
	studentBO.updateStudent(student)
	//获取学生
	gots[1] = studentBO.getAllStudents()[0].Name + studentBO.getAllStudents()[0].RollNo
	if gots[1] != wants[1] {
		t.Errorf("doTask() = %v, want %v", gots[1], wants[1])
	}
}
