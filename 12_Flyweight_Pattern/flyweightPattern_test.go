package _2_Flyweight_Pattern

import (
	"testing"
	"unsafe"
)

func TestFlyweightPattern(t *testing.T) {
	tests := []struct {
		name, color, want    string
		x, y, radius         int
		name2, color2, want2 string
		x2, y2, radius2      int
	}{
		{"Red,36,71,100", "Red", "Color :Red, x :36, y :71, radius :100", 36, 71, 100,
			"Red,71,36,100", "Red", "Color :Red, x :71, y :36, radius :100", 71, 36, 100},
		{"Black,36,71,100", "Black", "Color :Black, x :36, y :71, radius :100", 36, 71, 100,
			"Black,71,36,100", "Black", "Color :Black, x :71, y :36, radius :100", 71, 36, 100},
		{"Green,36,71,100", "Green", "Color :Green, x :36, y :71, radius :100", 36, 71, 100,
			"Green,36,100,71", "Green", "Color :Green, x :36, y :100, radius :71", 36, 100, 71},
	}
	for _, tt := range tests {
		got := (*Circle)(unsafe.Pointer(getCircle(tt.color)))
		if got.setArgs(tt.x, tt.y, tt.radius); got.Draw() != tt.want {
			t.Errorf("Draw() = %v, want %v", got.Draw(), tt.want)
		}
		got2 := (*Circle)(unsafe.Pointer(getCircle(tt.color2)))
		if got2.setArgs(tt.x2, tt.y2, tt.radius2); got2.Draw() != tt.want2 {
			t.Errorf("Draw() = %v, want %v", got2.Draw(), tt.want2)
		}
		if got != got2 {
			t.Errorf("Draw() = %v, want %v", got2.Draw(), tt.want2)
		}
	}
}
