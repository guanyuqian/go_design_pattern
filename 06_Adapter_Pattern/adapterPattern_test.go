package _6_Adapter_Pattern

import (
	"testing"
)

//步骤 5
//使用 AudioPlayer 来播放不同类型的音频格式。
func TestAdapterPattern(t *testing.T) {
	type args struct {
		audioType string
		fileName  string
	}
	var tests = []struct {
		name string
		args args
		want string
	}{
		{"mp3", args{"mp3", "beyond the horizon.mp3"}, "support"},
		{"mp4", args{"mp4", "alone.mp4"}, "support"},
		{"vlc", args{"vlc", "far far away.vlc"}, "support"},
		{"avi", args{"avi", "mind me.avi"}, "not support"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := AudioPlayer{}
			if got := receiver.Play(tt.args.audioType, tt.args.fileName); got != tt.want {
				t.Errorf("play() = %v, want %v", got, tt.want)
			}
		})
	}
}
