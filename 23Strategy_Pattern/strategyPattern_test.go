package _3Strategy_Pattern

import "testing"

func TestStrategyPattern(t *testing.T) {
	type fields struct {
		strategy Strategy
	}
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"10 + 5 = ", fields{&OperationAdd{}}, args{10, 5}, 15},
		{"10 - 5 = ", fields{&OperationSubtract{}}, args{10, 5}, 5},
		{"10 * 5 = ", fields{&OperationMultiply{}}, args{10, 5}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := Context{
				strategy: tt.fields.strategy,
			}
			if got := receiver.executeStrategy(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("executeStrategy() = %v, want %v", got, tt.want)
			}
		})
	}
}
