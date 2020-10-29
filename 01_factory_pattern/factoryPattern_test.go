package _1_factory_pattern

import (
	"reflect"
	"testing"
)

//步骤 4
//使用该工厂，通过传递类型信息来获取实体类的对象。
func TestFactoryPattern(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{name: "Rectangle", args: "Rectangle", want: "Rectangle"},
		{name: "Square", args: "Square", want: "Square"},
		{name: "Circle", args: "Circle", want: "Circle"},
	}

	shapeFactory := ShapeFactory{}
	for _, tc := range tests {
		got := shapeFactory.GetShape(tc.args).Draw()
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted: %#v, got: %#v", tc.want, got)
		}

	}

}
