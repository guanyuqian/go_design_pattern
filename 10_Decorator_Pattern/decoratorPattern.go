package _0_Decorator_Pattern

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

//步骤 3
//创建实现了 Shape 接口的抽象装饰类。
type ShapeDecorator interface {
	Draw() string
	setDecoratedShape(decoratedShape Shape)
}

//步骤 4
//创建扩展了 ShapeDecorator 类的实体装饰类。
type RedShapeDecorator struct {
	decoratedShape Shape
}

func (receiver *RedShapeDecorator) setDecoratedShape(decoratedShape Shape) {
	receiver.decoratedShape = decoratedShape
}
func (receiver *RedShapeDecorator) Draw() string {
	return receiver.decoratedShape.Draw() + "| " + receiver.setRedBorder()
}
func (receiver *RedShapeDecorator) setRedBorder() string {
	return "RedBorder"
}
