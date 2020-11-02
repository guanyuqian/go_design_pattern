package _1_Facade_Pattern

import (
	"../01_Factory_Pattern"
)

//步骤 1
//创建一个接口：
type Shape = _1_Factory_Pattern.Shape

//步骤 2
//创建实现接口的实体类。
type Rectangle = _1_Factory_Pattern.Rectangle
type Circle = _1_Factory_Pattern.Circle
type Square = _1_Factory_Pattern.Square

//步骤 3
//创建一个外观类。
type ShapeMaker struct {
	circle, rectangle, square Shape
}

func (receiver *ShapeMaker) drawCircle() string {
	return receiver.circle.Draw()
}

func (receiver *ShapeMaker) drawSquare() string {
	return receiver.square.Draw()
}
func (receiver *ShapeMaker) drawRectangle() string {
	return receiver.rectangle.Draw()
}
