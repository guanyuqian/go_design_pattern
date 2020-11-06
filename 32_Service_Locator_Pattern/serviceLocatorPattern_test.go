package _2_Service_Locator_Pattern

import (
	"reflect"
	"testing"
)

//步骤 6
//使用 ServiceLocator 来演示服务定位器设计模式。
func TestServiceLocatorPattern(t *testing.T) {
	tests := []struct {
		name     string
		jndiName string
		want     string
	}{
		{"Service1", "Service1", "Looking up and creating a new Service1 objectExecuting Service1"},
		{"Service2", "Service2", "Looking up and creating a new Service2 objectExecuting Service2"},
		{"Service1", "Service1", "Returning cached Service1 objectExecuting Service1"},
		{"Service2", "Service2", "Returning cached Service2 objectExecuting Service2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, got := getService(tt.jndiName)
			if got += service.execute(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getService() = %v, want %v", got, tt.want)
			}
		})
	}
}
