package implement

import (
	"awesomeProject/auxiliary"
)

//步骤 5
//创建一个 Meal 类，带有上面定义的 Item 对象。
type Meal struct {
	Items []auxiliary.Item
}

func (receiver *Meal) addItem(item auxiliary.Item) {
	receiver.Items = append(receiver.Items, item)
}

//func (receiver Meal) getCost() (cost float64) {
//	for _, item := range receiver.Items {
//		cost += item.Price()
//	}
//	return
//}

func (receiver Meal) ShowItems() (name, pack string, price float64) {
	for _, item := range receiver.Items {
		name += item.Name()
		price += item.Price()
		pack += item.Packing().Pack()
	}
	return
}

//步骤 6
//创建一个 MealBuilder 类，实际的 builder 类负责创建 Meal 对象。
//go中直接实现接口函数
func PrepareVegMeal() (meal Meal) {
	meal.addItem(new(auxiliary.VegBurger))
	meal.addItem(new(auxiliary.Coke))
	return
}

func PrepareNonVegMeal() (meal Meal) {
	meal.addItem(new(auxiliary.ChickenBurger))
	meal.addItem(new(auxiliary.Pepsi))
	return
}
