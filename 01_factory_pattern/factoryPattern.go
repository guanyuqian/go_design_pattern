package _1_factory_pattern

// 工厂模式

//步骤 1
//创建一个接口:
type Shape interface {
	Draw() string
}
type Color interface {
	Fill() string
}

//步骤 2
//创建实现接口的实体类。
type Rectangle struct{}
type Square struct{}
type Circle struct{}

func (r Rectangle) Draw() string {
	return "Rectangle"
}

func (s Square) Draw() string {
	return "Square"
}

func (c Circle) Draw() string {
	return "Circle"
}

type Green struct{}
type Red struct{}
type Blue struct{}

func (r Blue) Fill() string {
	return "Blue"
}

func (s Red) Fill() string {
	return "Red"
}

func (c Green) Fill() string {
	return "Green"
}

//步骤 3
//创建一个工厂，生成基于给定信息的实体类的对象。
type ShapeFactory struct{}
type ColorFactory struct{}

func (s ShapeFactory) GetShape(shape string) Shape {
	switch shape {
	case "Rectangle":
		return new(Rectangle)
	case "Square":
		return new(Square)
	case "Circle":
		return new(Circle)
	}
	return nil
}

func (s ColorFactory) GetColor(color string) Color {
	switch color {
	case "Green":
		return new(Green)
	case "Red":
		return new(Red)
	case "Blue":
		return new(Blue)
	}
	return nil
}

func (s ShapeFactory) GetColor(color string) Color {
	return nil
}
func (s ShapeFactory) Name() string {
	return "Shape"
}

func (s ColorFactory) GetShape(shape string) Shape {
	return nil
}
func (s ColorFactory) Name() string {
	return "Color"
}
