package _2_Abstract_Factory_Pattern

import "../01_Factory_Pattern"

// 抽象工厂模式

//步骤 1
//为形状创建一个接口。
type Shape = _1_Factory_Pattern.Shape

//步骤 2
//创建实现接口的实体类。
type ShapeFactory = _1_Factory_Pattern.ShapeFactory

//步骤 3
//为颜色创建一个接口。
type Color = _1_Factory_Pattern.Color

//步骤4
//创建实现接口的实体类。
type ColorFactory = _1_Factory_Pattern.ColorFactory

//步骤 5
//为 color 和 Shape 对象创建抽象类来获取工厂。

//步骤 6
//创建扩展了 AbstractFactory 的工厂类，基于给定的信息生成实体类的对象。
type AbstractFactory interface {
	Name() string
	GetColor(string) Color
	GetShape(string) Shape
}

//步骤 7
//创建一个工厂创造器/生成器类，通过传递形状或颜色信息来获取工厂。
type FactoryProducer struct{}

func GetFactory(choice string) AbstractFactory {
	switch choice {
	case "color":
		return new(ColorFactory)
	case "Shape":
		return new(ShapeFactory)
	}
	return nil
}
