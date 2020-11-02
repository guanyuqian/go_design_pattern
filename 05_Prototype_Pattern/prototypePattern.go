package _5_Prototype_Pattern

// 原型模式

//步骤 1
//创建一个实现了 Cloneable 接口的抽象类。
type CloneableShape interface {
	Draw() string
	Clone() CloneableShape
	SetId(string)
	GetId() string
}

//步骤 2
//创建扩展了上面抽象类的实体类。
type CloneableRectangle struct {
	id string
}
type CloneableSquare struct {
	id string
}
type CloneableCircle struct {
	id string
}

func (receiver *CloneableRectangle) Draw() string {
	return "CloneableRectangle"
}

func (receiver *CloneableSquare) Draw() string {
	return "CloneableSquare"
}

func (receiver *CloneableCircle) Draw() string {
	return "CloneableCircle"
}

func (receiver *CloneableRectangle) Clone() CloneableShape {
	return &CloneableRectangle{receiver.id}
}

func (receiver *CloneableSquare) Clone() CloneableShape {
	return &CloneableSquare{receiver.id}
}

func (receiver *CloneableCircle) Clone() CloneableShape {
	return &CloneableCircle{receiver.id}
}

func (receiver *CloneableRectangle) GetId() string {
	return receiver.id
}

func (receiver *CloneableSquare) GetId() string {
	return receiver.id
}
func (receiver *CloneableCircle) GetId() string {
	return receiver.id
}

func (receiver *CloneableCircle) SetId(id string) {
	receiver.id = id
}
func (receiver *CloneableSquare) SetId(id string) {
	receiver.id = id
}
func (receiver *CloneableRectangle) SetId(id string) {
	receiver.id = id
}

//步骤 3
//创建一个类，从数据库获取实体类，并把它们存储在一个 Hashtable 中。
var cloneableShape map[string]CloneableShape

func init() {
	cloneableShape = map[string]CloneableShape{}
	loadCache()
}

func loadCache() {
	cloneableCircle := CloneableCircle{}
	cloneableCircle.SetId("1")
	cloneableShape[cloneableCircle.GetId()] = &cloneableCircle

	cloneableSquare := CloneableSquare{}
	cloneableSquare.SetId("2")
	cloneableShape[cloneableSquare.GetId()] = &cloneableSquare

	cloneableRectangle := CloneableRectangle{}
	cloneableRectangle.SetId("3")
	cloneableShape[cloneableRectangle.GetId()] = &cloneableRectangle
}

func GetShape(shapeId string) CloneableShape {
	return cloneableShape[shapeId].Clone()
}
