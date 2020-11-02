package _7_Bridge_Pattern

// 桥接模式

//步骤 1
//创建桥接实现接口。
type DrawAPI interface {
	drawCircle(radius, x, y int) string
}

//步骤 2
//创建实现了 DrawAPI 接口的实体桥接实现类。
type RedCircle struct{}

func (receiver RedCircle) drawCircle(radius, x, y int) string {
	return "RedCircle"
}

type GreenCircle struct{}

func (receiver GreenCircle) drawCircle(radius, x, y int) string {
	return "GreenCircle"
}

//步骤 3
//使用 DrawAPI 接口创建抽象类 Shape。
// 此处DrawCircle实现Shape接口

//步骤 4
//创建实现了 Shape 接口的实体类。
type DrawCircle struct {
	Radius, X, Y int
	DrawAPI      DrawAPI
}

func (receiver DrawCircle) Draw() string {
	return receiver.DrawAPI.drawCircle(receiver.Radius, receiver.X, receiver.Y)
}
