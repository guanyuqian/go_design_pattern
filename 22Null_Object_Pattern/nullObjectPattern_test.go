package _2Null_Object_Pattern

import (
	"reflect"
	"testing"
)

//步骤 4
//使用 CustomerFactory，基于客户传递的名字，来获取 RealCustomer 或 NullCustomer 对象。
func TestNullObjectPattern(t *testing.T) {
	tests := []struct {
		name string
		argv string
		want string
	}{
		{"Rob", "Rob", "Rob"},
		{"Bob", "Bob", "Not Available in Customer Database"},
		{"Julie", "Julie", "Julie"},
		{"Laura", "Laura", "Not Available in Customer Database"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCustomer(tt.argv).getName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}
