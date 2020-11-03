package _2_Abstract_Factory_Pattern

import (
	"reflect"
	"testing"
)

//步骤 8
//使用 FactoryProducer 来获取 AbstractFactory，通过传递类型信息来获取实体类的对象。
func TestAbstractFactoryPattern(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{"color", "color", "color"},
		{"Shape", "Shape", "Shape"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 测试抽象工厂
			factory := GetFactory(tt.args)
			if !reflect.DeepEqual(factory.Name(), tt.want) {
				t.Errorf("GetFactory() = %v, want %v", factory.Name(), tt.want)
			}

			// 测试 color 工厂
			colorTests := []struct {
				name string
				args string
				want string
			}{
				{"Red", "Red", "Red"},
				{"Blue", "Blue", "Blue"},
				{"Green", "Green", "Green"},
			}
			for _, ctt := range colorTests {
				t.Run(ctt.name, func(t *testing.T) {
					color := factory.GetColor(ctt.args)
					if factory.Name() != "color" {
						if !reflect.DeepEqual(color, nil) {
							t.Errorf("GetFactory() = %v, want %v", color, nil)
						}
					} else {
						if !reflect.DeepEqual(color.Fill(), ctt.want) {
							t.Errorf("GetFactory() = %v, want %v", color, ctt.want)
						}
					}
				})
			}

			// 测试 shape 工厂
			shapeTests := []struct {
				name string
				args string
				want string
			}{
				{"Rectangle", "Rectangle", "Rectangle"},
				{"Square", "Square", "Square"},
				{"Circle", "Circle", "Circle"},
			}
			for _, stt := range shapeTests {
				t.Run(stt.name, func(t *testing.T) {
					shape := factory.GetShape(stt.args)
					if factory.Name() != "Shape" {
						if !reflect.DeepEqual(shape, nil) {
							t.Errorf("GetFactory() = %v, want %v", shape, nil)
						}
					} else {
						if !reflect.DeepEqual(shape.Draw(), stt.want) {
							t.Errorf("GetFactory() = %v, want %v", shape, stt.want)
						}
					}
				})
			}
		})
	}

}
