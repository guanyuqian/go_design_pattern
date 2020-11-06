package _6_MVC_Pattern

//步骤 1
//创建模型。
type Student struct {
	Name, RollNo string
}

//步骤 2
//创建视图。
type StudentView struct {
}

func (receiver StudentView) printStudentDetails(studentName, studentRollNo string) string {
	return "Student: " + "Name: " + studentName + "Roll No: " + studentRollNo
}

//步骤 3
//创建控制器。
type StudentController struct {
	model *Student
	view  *StudentView
}

func NewStudentController(model *Student, view *StudentView) *StudentController {
	return &StudentController{model: model, view: view}
}

func (receiver *StudentController) updateView() string {
	return receiver.view.printStudentDetails(receiver.model.Name, receiver.model.RollNo)
}

func (receiver StudentController) setStudentName(name string) {
	receiver.model.Name = name
}

func (receiver StudentController) getStudentName() string {
	return receiver.model.Name
}
