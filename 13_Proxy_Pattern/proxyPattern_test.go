package _3_Proxy_Pattern

import "testing"

//步骤 3
//当被请求时，使用 ProxyImage 来获取 RealImage 类的对象。
func TestProxyImage_display(t *testing.T) {
	fileName := "test_10mb.jpg"
	tests := []struct {
		name string
		want string
	}{
		{"图像将从磁盘加载", "Loading test_10mb.jpg, Displaying test_10mb.jpg"},
		{"图像不需要从磁盘加载", "Displaying test_10mb.jpg"},
		{"图像不需要从磁盘加载", "Displaying test_10mb.jpg"},
	}
	receiver := &ProxyImage{
		fileName: fileName,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := receiver.display(); got != tt.want {
				t.Errorf("display() = %v, want %v", got, tt.want)
			}
		})
	}
}
