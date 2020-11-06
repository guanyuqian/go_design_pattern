package _8_Composite_Entity_Pattern

import (
	"testing"
)

//步骤 5
//使用 Client 来演示组合实体设计模式的用法。
func TestBusinessDelegatePattern(t *testing.T) {

	client := &Client{}
	client.setData("Test", "Data")
	tests := []struct {
		name string
		arg1 string
		arg2 string
		want string
	}{
		{"Test", "Test", "Data", "TestData"},
		{"Second Test", "Second Test", "Data1", "Second TestData1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client.setData(tt.arg1, tt.arg2)
			if got := client.printData(); got != tt.want {
				t.Errorf("doTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
