package _9_Composite_Pattern

import "testing"

var CEO, headSales, headMarketing, clerk1, clerk2, salesExecutive1, salesExecutive2 *Employee

func init() {
	CEO = newEmployee("John", "CEO", 30000)

	headSales = newEmployee("Robert", "Head Sales", 20000)

	headMarketing = newEmployee("Michel", "Head Marketing", 20000)

	clerk1 = newEmployee("Laura", "Marketing", 10000)
	clerk2 = newEmployee("Bob", "Marketing", 10000)

	salesExecutive1 = newEmployee("Richard", "Sales", 10000)
	salesExecutive2 = newEmployee("Rob", "Sales", 10000)

	CEO.add(headSales)
	CEO.add(headMarketing)

	headSales.add(salesExecutive1)
	headSales.add(salesExecutive2)

	headMarketing.add(clerk1)
	headMarketing.add(clerk2)
}

//步骤 2
//使用 Employee 类来创建和打印员工的层次结构。

func TestEmployee_toString(t *testing.T) {
	want := "Employee :[ Name : John, dept : CEO, salary :30000 ]"
	want2 := [2]string{"Employee :[ Name : Robert, dept : Head Sales, salary :20000 ]", "Employee :[ Name : Michel, dept : Head Marketing, salary :20000 ]"}
	if got := CEO.toString(); got != want {
		t.Errorf("toString() = %v, want %v", got, want)
	}

	for i, tt := range CEO.subordinates {
		if got := tt.toString(); got != want2[i] {
			t.Errorf("toString() = %v, want %v", got, want2[i])
		}
	}
}
