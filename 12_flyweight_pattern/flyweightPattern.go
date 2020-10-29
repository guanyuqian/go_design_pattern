package _2_flyweight_pattern

import (
	"../01_factory_pattern"
	"strconv"
	"unsafe"
)

//步骤 1
//创建一个接口。
type Shape = _1_factory_pattern.Shape

//步骤 2
//创建实现接口的实体类。
type Circle _1_factory_pattern.Circle

func (receiver *Circle) setArgs(x, y, radius int) {
	receiver.X = x
	receiver.Y = y
	receiver.Radius = radius
}
func (receiver Circle) Draw() string {
	return "Color :" + receiver.Color + ", " + "x :" + strconv.Itoa(receiver.X) + ", " +
		"y :" + strconv.Itoa(receiver.Y) + ", " +
		"radius :" + strconv.Itoa(receiver.Radius)
}

var circleMap map[string]*Shape

func init() {
	circleMap = make(map[string]*Shape, 3)
}
func getCircle(color string) *Shape {
	_, ok := circleMap[color]
	if !ok {
		circleMap[color] = (*Shape)(unsafe.Pointer(&Circle{Color: color}))
	}
	return circleMap[color]
}
