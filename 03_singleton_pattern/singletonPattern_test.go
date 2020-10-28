package _3_singleton_pattern

import (
	"awesomeProject/implement"
	"reflect"
	"testing"
)

//步骤 2
//从 singleton 类获取唯一的对象。
func TestGetHungryInstance(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"1", 1},
		{"2", 1},
		{"3", 1},
		{"4", 1},
		{"5", 1},
		{"6", 1},
		{"7", 1},
		{"8", 1},
		{"9", 1},
		{"10", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := implement.GetHungryInstance(); !reflect.DeepEqual(got.Counter, tt.want) {
				t.Errorf("GetHungryInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLazyInstance(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"1", 1},
		{"2", 1},
		{"3", 1},
		{"4", 1},
		{"5", 1},
		{"6", 1},
		{"7", 1},
		{"8", 1},
		{"9", 1},
		{"10", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := implement.GetLazyInstance(); !reflect.DeepEqual(got.Counter, tt.want) {
				t.Errorf("GetLazyInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
