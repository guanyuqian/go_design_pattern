package _9_Data_Access_Object_Pattern

import (
	"../26_MVC_Pattern"
	"strconv"
)

//步骤 1
//创建数值对象。
type Student _6_MVC_Pattern.Student

type StudentDao interface {
	getAllStudents() []*Student
	getStudent(RollNo int) Student
	updateStudent(student *Student)
	deleteStudent(student *Student)
}

type StudentDaoImpl struct {
	students []*Student
}

func NewStudentDaoImpl() *StudentDaoImpl {
	return &StudentDaoImpl{[]*Student{
		&Student{Name: "Robert", RollNo: "0"},
		&Student{Name: "John", RollNo: "1"},
	}}
}

func (receiver *StudentDaoImpl) deleteStudent(student *Student) {
	for i, v := range receiver.students {
		if v.RollNo == student.RollNo {
			receiver.students = append(receiver.students[:i-1], receiver.students[i+1:]...)
			return
		}
	}
}

//从数据库中检索学生名单
func (receiver *StudentDaoImpl) getAllStudents() []*Student {
	return receiver.students
}
func (receiver *StudentDaoImpl) getStudent(RollNo int) (ret Student) {
	for _, v := range receiver.students {
		if v.RollNo == strconv.Itoa(RollNo) {
			ret = *v
		}
	}
	return
}

func (receiver *StudentDaoImpl) updateStudent(student *Student) {
	for i, v := range receiver.students {
		if v.RollNo == student.RollNo {
			receiver.students[i].Name = student.Name
		}
	}
}
