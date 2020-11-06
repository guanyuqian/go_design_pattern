package _4_Template_Pattern

import "testing"

//步骤 3
//使用 Game 的模板方法 play() 来演示游戏的定义方式。
func TestTemplatePattern(t *testing.T) {

	tests := []struct {
		name    string
		game    GameImplement
		wantRes string
	}{
		{"Cricket Game", newCricketGame(), "Cricket Game Initialized! Start playing.Cricket Game Started. Enjoy the game!Cricket Game Finished!"},
		{"Football Game", newFootballGame(), "Football Game Initialized! Start playing.Football Game Started. Enjoy the game!Football Game Finished!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tt.game.play(); gotRes != tt.wantRes {
				t.Errorf("play() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
