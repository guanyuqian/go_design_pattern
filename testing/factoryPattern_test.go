package testing

import (
	"awesomeProject/implement"
	"reflect"
	"testing"
)

// 测试简单方法模式
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

	shapeFactory := implement.ShapeFactory{}
	for _, tc := range tests {
		got := shapeFactory.GetShape(tc.args).Draw()
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted: %#v, got: %#v", tc.want, got)
		}

	}

}
