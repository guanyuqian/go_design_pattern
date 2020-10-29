package _4_builder_pattern

import (
	"testing"
)

//步骤 7
//BuiderPatternDemo 使用 MealBuider 来演示建造者模式（Builder Pattern）。
func TestMeal_showIterms(t *testing.T) {
	var tests = []struct {
		name      string
		meal      Meal
		wantName  string
		wantPack  string
		wantPrice float64
	}{
		{"Veg Meal", PrepareVegMeal(),
			"Veg Burger" + "Coke", "Wrapper" + "Bottle", 25.0 + 30.0},
		{"Non-Veg Meal", PrepareNonVegMeal(),
			"Chicken Burger" + "Pepsi", "Wrapper" + "Bottle", 50.5 + 35.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.meal
			gotName, gotPack, gotPrice := receiver.ShowItems()
			if gotName != tt.wantName {
				t.Errorf("showIterms() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotPack != tt.wantPack {
				t.Errorf("showIterms() gotPack = %v, want %v", gotPack, tt.wantPack)
			}
			if gotPrice != tt.wantPrice {
				t.Errorf("showIterms() gotPrice = %v, want %v", gotPrice, tt.wantPrice)
			}
		})
	}
}
