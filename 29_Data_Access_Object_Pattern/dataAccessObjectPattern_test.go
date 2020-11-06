package _9_Data_Access_Object_Pattern

import "testing"

//步骤 4
//使用 StudentDao 来演示数据访问对象模式的用法。
func TestDataAccessObjectPattern(t *testing.T) {
	var studentDao StudentDao = NewStudentDaoImpl()
	gots := [2]string{}
	wants := [2]string{"Robert0John1", "Michael0"}

	for _, v := range studentDao.getAllStudents() {
		gots[0] += v.Name + v.RollNo
	}
	if gots[0] != wants[0] {
		t.Errorf("doTask() = %v, want %v", gots[0], wants[0])
	}

	//更新学生
	student := studentDao.getAllStudents()[0]
	student.Name = "Michael"
	studentDao.updateStudent(student)
	//获取学生
	gots[1] = studentDao.getAllStudents()[0].Name + studentDao.getAllStudents()[0].RollNo
	if gots[1] != wants[1] {
		t.Errorf("doTask() = %v, want %v", gots[1], wants[1])
	}
}
