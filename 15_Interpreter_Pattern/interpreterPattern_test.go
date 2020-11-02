package _5_Interpreter_Pattern

import "testing"

func TestTerminalExpression_interpret(t *testing.T) {
	//规则：Robert 和 John 是男性
	robert := TerminalExpression{"Robert"}
	john := TerminalExpression{"John"}
	isMale := OrExpression{robert, john}

	//规则：Julie 是一个已婚的女性
	julie := TerminalExpression{"Julie"}
	married := TerminalExpression{"Married "}
	isMarriedWoman := OrExpression{julie, married}

	tests := []struct {
		name       string
		expression Expression
		args       string
		want       bool
	}{
		{"John is male?", isMale, "John", true},
		{"Julie is a married women?", isMarriedWoman, "Married Julie", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.expression.interpret(tt.args); got != tt.want {
				t.Errorf("interpret() = %v, want %v", got, tt.want)
			}
		})
	}
}
