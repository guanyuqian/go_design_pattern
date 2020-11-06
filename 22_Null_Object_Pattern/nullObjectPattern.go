package _2_Null_Object_Pattern

//步骤 1
//创建一个抽象类。
type AbstractCustomer interface {
	isNil() bool
	getName() string
}

//步骤 2
//创建扩展了上述类的实体类。

type RealCustomer struct {
	name string
}

func (receiver *RealCustomer) isNil() bool {
	return false
}

func (receiver *RealCustomer) getName() string {
	return receiver.name
}

type NullCustomer struct{}

func (receiver *NullCustomer) isNil() bool {
	return true
}

func (receiver *NullCustomer) getName() string {
	return "Not Available in Customer Database"
}

//步骤 3
//创建 CustomerFactory 类。

var names = [3]string{"Rob", "Joe", "Julie"}

func getCustomer(name string) AbstractCustomer {
	for _, v := range names {
		if v == name {
			return &RealCustomer{name}
		}
	}
	return &NullCustomer{}
}
