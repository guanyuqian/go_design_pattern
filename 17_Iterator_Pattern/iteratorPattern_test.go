package _7_Iterator_Pattern

import (
	"reflect"
	"testing"
)

//步骤 3
//使用 NameRepository 来获取迭代器，并打印名字。
func TestIteratorPattern(t *testing.T) {
	nameRepository := getNameRepository()
	want := nameRepository.names
	i := 0
	for iter := nameRepository.getIterator(); iter.hasNext(); i++ {
		got := iter.next().(string)
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("next() = %v, want %v", got, want[i])
		}
	}
}
