package _0_Observer_Pattern

import "testing"

//步骤 4
//使用 Subject 和实体观察者对象。
func TestObserverPattern(t *testing.T) {
	subject := Subject{}
	tests := []struct {
		name     string
		observer Observer
		state    int64
		want     string
	}{
		{"HexaObserver", &HexObserver{&subject}, 15, "Hex String: f."},
		{"HexaObserver", &OctalObserver{&subject}, 15, "Hex String: f.Octal String: 17."},
		{"HexaObserver", &BinaryObserver{&subject}, 10, "Hex String: a.Octal String: 12.Binary String: 1010."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subject.attach(tt.observer)
			if got := subject.setState(tt.state); got != tt.want {
				t.Errorf("setState() = %v, want %v", got, tt.want)
			}
		})
	}
}
