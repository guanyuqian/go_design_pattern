package _3_Transfer_Object_Pattern

//步骤 1
//创建数值对象。
type StudentVO struct {
	Name, RollNo string
}

type StudentBO struct {
	students []*StudentVO
}

func NewStudentBO() *StudentBO {
	return &StudentBO{[]*StudentVO{
		&StudentVO{Name: "Robert", RollNo: "0"},
		&StudentVO{Name: "John", RollNo: "1"},
	}}
}

func (receiver *StudentBO) deleteStudent(student *StudentVO) {
	for i, v := range receiver.students {
		if v.RollNo == student.RollNo {
			receiver.students = append(receiver.students[:i-1], receiver.students[i+1:]...)
			return
		}
	}
}

//从数据库中检索学生名单
func (receiver *StudentBO) getAllStudents() []*StudentVO {
	return receiver.students
}
func (receiver *StudentBO) getStudent(RollNo string) (ret StudentVO) {
	for _, v := range receiver.students {
		if v.RollNo == RollNo {
			ret = *v
		}
	}
	return
}

func (receiver *StudentBO) updateStudent(student *StudentVO) {
	for i, v := range receiver.students {
		if v.RollNo == student.RollNo {
			receiver.students[i].Name = student.Name
		}
	}
}
