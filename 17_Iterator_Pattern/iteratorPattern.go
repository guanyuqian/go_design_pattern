package _7_Iterator_Pattern

//步骤 1
//创建接口
type Object interface{}
type Iterator interface {
	hasNext() bool
	next() Object
}
type Container interface {
	getIterator() Iterator
}

//步骤 2
//创建实现了 Container 接口的实体类。该类有实现了 Iterator 接口的内部类 NameIterator。
type NameRepository struct {
	names []string
}

func getNameRepository() *NameRepository {
	return &NameRepository{[]string{"Robert", "John", "Julie", "Lora"}}
}

func (receiver *NameRepository) getIterator() Iterator {
	return &NameIterator{}
}

type NameIterator struct {
	NameRepository
	index int
}

func (receiver *NameIterator) hasNext() bool {
	return receiver.index < len(receiver.names)
}

func (receiver *NameIterator) next() (ret Object) {
	if receiver.hasNext() {
		ret = &(receiver.names[receiver.index])
		receiver.index++
	}
	return
}
