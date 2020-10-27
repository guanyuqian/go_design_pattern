package auxiliary

// 步骤 1
// 创建一个表示食物条目和食物包装的接口。
type Packing interface {
	Pack() string
}

type Item interface {
	Name() string
	Packing() Packing
	Price() float64
}

// 步骤 2
// 创建实现 Packing 接口的实体类。
type Wrapper struct{}
type Bottle struct{}

func (w Wrapper) Pack() string {
	return "Wrapper"
}
func (b Bottle) Pack() string {
	return "Bottle"
}

// 步骤 3
// 创建实现 Item 接口的抽象类，该类提供了默认的功能。
type Burger struct{}
type ColdDrink struct{}

func (b Burger) Packing() Packing {
	return new(Wrapper)
}
func (b ColdDrink) Packing() Packing {
	return new(Bottle)
}

//步骤 4
//创建扩展了 Burger 和 ColdDrink 的实体类。

type VegBurger struct {
	Burger
}

func (receiver VegBurger) Price() float64 {
	return 25.0
}

func (receiver VegBurger) Name() string {
	return "Veg Burger"
}

type ChickenBurger struct {
	Burger
}

func (receiver ChickenBurger) Price() float64 {
	return 50.5
}

func (receiver ChickenBurger) Name() string {
	return "Chicken Burger"
}

type Coke struct {
	ColdDrink
}

func (receiver Coke) Price() float64 {
	return 30.0
}

func (receiver Coke) Name() string {
	return "Coke"
}

type Pepsi struct {
	ColdDrink
}

func (receiver Pepsi) Price() float64 {
	return 35.0
}

func (receiver Pepsi) Name() string {
	return "Pepsi"
}
