package implement

import "awesomeProject/auxiliary"

type ShapeFactory struct{}
type ColorFactory struct{}

func (s ShapeFactory) GetShape(shape string) auxiliary.Shape {
	switch shape {
	case "Rectangle":
		return new(auxiliary.Rectangle)
	case "Square":
		return new(auxiliary.Square)
	case "Circle":
		return new(auxiliary.Circle)
	}
	return nil
}

func (s ColorFactory) GetColor(color string) auxiliary.Color {
	switch color {
	case "Green":
		return new(auxiliary.Green)
	case "Red":
		return new(auxiliary.Red)
	case "Blue":
		return new(auxiliary.Blue)
	}
	return nil
}
