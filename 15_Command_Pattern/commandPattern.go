package _5_Command_Pattern

import "strconv"

//步骤 1
//创建一个命令接口。
type Order interface {
	execute() string
}

//步骤 2
//创建一个请求类。
type Stock struct {
	name     string
	quantity int
}

func (receiver Stock) buy() string {
	return "Stock [ name: " + receiver.name + ",Quantity: " + strconv.Itoa(receiver.quantity) + " ] bought."
}
func (receiver Stock) sell() string {
	return "Stock [ name: " + receiver.name + ",Quantity: " + strconv.Itoa(receiver.quantity) + " ] sold."
}

//步骤 3
//创建实现了 Order 接口的实体类。
type BuyStock struct {
	abcStock *Stock
}

func (receiver *BuyStock) execute() string {
	return receiver.abcStock.buy()
}

type SellStock struct {
	abcStock *Stock
}

func (receiver *SellStock) execute() string {
	return receiver.abcStock.sell()
}

//步骤 4
//创建命令调用类。
type Broker struct {
	orderList []Order
}

func (receiver *Broker) takeOrder(order Order) {
	receiver.orderList = append(receiver.orderList, order)
}

func (receiver *Broker) placeOrders() (ret string) {
	for _, order := range receiver.orderList {
		ret += order.execute()
	}
	return
}
