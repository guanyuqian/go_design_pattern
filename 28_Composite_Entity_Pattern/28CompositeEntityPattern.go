package _8_Composite_Entity_Pattern

//步骤 1
//创建依赖对象。

type DependentObject1 struct {
	data string
}

type DependentObject2 struct {
	data string
}

//步骤 2
//创建粗粒度对象。
type CoarseGrainedObject struct {
	do1 DependentObject1
	do2 DependentObject2
}

func (receiver *CoarseGrainedObject) getData() []string {
	return []string{receiver.do1.data, receiver.do2.data}
}

func (receiver *CoarseGrainedObject) setData(data1, data2 string) {
	receiver.do1.data = data1
	receiver.do2.data = data2
}

//步骤 3
//创建组合实体。

type CompositeEntity struct {
	coarseGrainedObject CoarseGrainedObject
}

func (receiver *CompositeEntity) getData() []string {
	return receiver.coarseGrainedObject.getData()
}

func (receiver *CompositeEntity) setData(data1, data2 string) {
	receiver.coarseGrainedObject.setData(data1, data2)
}

//步骤 4
//创建使用组合实体的客户端类。

type Client struct {
	compositeEntity CompositeEntity
}

func (receiver *Client) printData() (res string) {
	for _, v := range receiver.compositeEntity.getData() {
		res += v
	}
	return
}

func (receiver *Client) setData(data1, data2 string) {
	receiver.compositeEntity.setData(data1, data2)
}
