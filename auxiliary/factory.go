package auxiliary

// 下列定义了shape接口以及相关的三个shape类

type Color interface {
	Fill() string
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

// 下列定义了shape接口以及相关的三个shape类

type Shape interface {
	Draw() string
}

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


