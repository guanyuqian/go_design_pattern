package _1_Facade_Pattern

import "testing"

//步骤 4
//使用该外观类画出各种类型的形状。
func TestShapeMaker_drawCircle(t *testing.T) {
	want := [3]string{"Circle", "Rectangle", "Square"}
	receiver := &ShapeMaker{Circle{}, Rectangle{}, Square{}}
	if got := receiver.drawCircle(); got != want[0] {
		t.Errorf("Draw() = %v, want %v", got, want[0])
	}
	if got := receiver.drawRectangle(); got != want[1] {
		t.Errorf("Draw() = %v, want %v", got, want[1])
	}
	if got := receiver.drawSquare(); got != want[2] {
		t.Errorf("Draw() = %v, want %v", got, want[2])
	}
}
