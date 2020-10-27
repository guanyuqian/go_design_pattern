package testing

import (
	"awesomeProject/implement"
	"testing"
)

func TestMeal_showIterms(t *testing.T) {
	var tests = []struct {
		name      string
		meal      implement.Meal
		wantName  string
		wantPack  string
		wantPrice float64
	}{
		{"Veg Meal", implement.PrepareVegMeal(),
			"Veg Burger" + "Coke", "Wrapper" + "Bottle", 25.0 + 30.0},
		{"Non-Veg Meal", implement.PrepareNonVegMeal(),
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
