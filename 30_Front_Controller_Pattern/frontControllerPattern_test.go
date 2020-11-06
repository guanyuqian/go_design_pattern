package _0_Front_Controller_Pattern

import (
	"testing"
)

//步骤 4
//使用 FrontController 来演示前端控制器设计模式。
func TestFrontControllerPattern(t *testing.T) {
	frontController := FrontController{}
	tests := []struct {
		name string
		args string
		want string
	}{
		{"HOME", "HOME", "Displaying Home Page"},
		{"STUDENT", "STUDENT", "Displaying Student Page"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frontController.dispatch(tt.args); got != tt.want {
				t.Errorf("dispatchRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
