package _5_Command_Pattern

import "testing"

//步骤 5
//使用 Broker 类来接受并执行命令。
func TestCommandPattern(t *testing.T) {
	abcStock := Stock{"ABC", 10}
	buyStockOrder := BuyStock{&abcStock}
	sellStockOrder := SellStock{&abcStock}
	broker := Broker{}
	broker.takeOrder(&buyStockOrder)
	broker.takeOrder(&sellStockOrder)
	got := broker.placeOrders()
	want := "Stock [ name: ABC,Quantity: 10 ] bought.Stock [ name: ABC,Quantity: 10 ] sold."
	if got != want {
		t.Errorf("placeOrders() = %v, want %v", got, want)
	}
}
