package _9_composite_pattern

import "strconv"

//步骤 1
//创建 Employee 类，该类带有 Employee 对象的列表。
type Employee struct {
	name, dept   string
	salary       int
	subordinates []Employee
}

func newEmployee(name, dept string, salary int) *Employee {
	return &Employee{name, dept, salary, []Employee{}}
}
func (receiver *Employee) Compare(employee *Employee) bool {
	return receiver.name == employee.name && receiver.dept == employee.dept && receiver.salary == employee.salary
}
func (receiver *Employee) add(employee *Employee) {
	receiver.subordinates = append(receiver.subordinates, *employee)
}
func (receiver *Employee) remove(employee *Employee) {
	for i, v := range receiver.subordinates {
		if v.Compare(employee) {
			receiver.subordinates = append(receiver.subordinates[:i], receiver.subordinates[i+1:]...)
			break
		}
	}
}

func (receiver *Employee) toString() string {
	return "Employee :[ Name : " +
		receiver.name +
		", dept : " +
		receiver.dept +
		", salary :" +
		strconv.Itoa(receiver.salary) +
		" ]"
}
