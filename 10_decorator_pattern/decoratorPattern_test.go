package _0_decorator_pattern

import "testing"

func TestRedShapeDecorator_Draw(t *testing.T) {
	type fields struct {
		decoratedShape Shape
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"redCircle", fields{&RedShapeDecorator{Circle{}}}, "Circle| RedBorder"},
		{"redCircle", fields{&RedShapeDecorator{Rectangle{}}}, "Rectangle| RedBorder"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.decoratedShape.Draw(); got != tt.want {
				t.Errorf("Draw() = %v, want %v", got, tt.want)
			}
		})
	}
}
