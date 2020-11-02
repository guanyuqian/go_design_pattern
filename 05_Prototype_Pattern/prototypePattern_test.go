package _5_Prototype_Pattern

import (
	"reflect"
	"testing"
)

//步骤 4
//PrototypePatternDemo 使用 ShapeCache 类来获取存储在 Hashtable 中的形状的克隆。
func TestGetShape(t *testing.T) {

	tests := []struct {
		name string
		argv string
		want string
	}{
		{"1", "1", "1"},
		{"2", "2", "2"},
		{"3", "3", "3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetShape(tt.argv).GetId(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShape() = %v, want %v", got, tt.want)
			}

			shape := GetShape(tt.argv)
			shape.SetId("10085")
		})
	}
}
