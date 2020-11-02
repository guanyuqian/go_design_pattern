package _3_Proxy_Pattern

//步骤 1
//创建一个接口。
type Image interface {
	display() string
}

//步骤 2
//创建实现接口的实体类。
type RealImage struct {
	fileName string
}

func (receiver *RealImage) display() string {
	return "Displaying " + receiver.fileName
}

func loadFromDisk(fileName string) (*RealImage, string) {
	return &RealImage{fileName}, "Loading " + fileName + ", "
}

type ProxyImage struct {
	realImage *RealImage
	fileName  string
}

func (receiver *ProxyImage) display() (msg string) {
	if receiver.realImage == nil {
		receiver.realImage, msg = loadFromDisk(receiver.fileName)
	}
	msg += receiver.realImage.display()
	return
}
