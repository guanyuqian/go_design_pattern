package _7_Bridge_Pattern

import (
	"testing"
)

//步骤 5
//使用 Shape 和 DrawAPI 类画出不同颜色的圆。
func TestBridgePattern(t *testing.T) {
	type fields struct {
		radius  int
		x       int
		y       int
		drawAPI DrawAPI
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"RedCircle", fields{100, 100, 10, new(RedCircle)}, "RedCircle"},
		{"GreenCircle", fields{100, 100, 10, new(GreenCircle)}, "GreenCircle"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := DrawCircle{
				Radius:  tt.fields.radius,
				X:       tt.fields.x,
				Y:       tt.fields.y,
				DrawAPI: tt.fields.drawAPI,
			}
			if got := receiver.Draw(); got != tt.want {
				t.Errorf("Draw() = %v, want %v", got, tt.want)
			}
		})
	}
}
