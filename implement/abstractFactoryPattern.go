package implement

import "awesomeProject/auxiliary"

type AbstractFactory interface {
	Name() string
	GetColor(string) auxiliary.Color
	GetShape(string) auxiliary.Shape
}

type FactoryProducer struct{}

func GetFactory(choice string) AbstractFactory {
	switch choice {
	case "Color":
		return new(ColorFactory)
	case "Shape":
		return new(ShapeFactory)
	}
	return nil
}

func (s ShapeFactory) GetColor(color string) auxiliary.Color {
	return nil
}
func (s ShapeFactory) Name() string {
	return "Shape"
}

func (s ColorFactory) GetShape(shape string) auxiliary.Shape {
	return nil
}
func (s ColorFactory) Name() string {
	return "Color"
}
