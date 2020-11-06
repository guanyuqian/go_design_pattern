package _2_Service_Locator_Pattern

//步骤 1
//创建服务接口 Service。
type Service interface {
	getName() string
	execute() string
}

type Service1 struct {
}

func (receiver Service1) getName() string {
	return "Service1"
}

func (receiver Service1) execute() string {
	return "Executing Service1"
}

type Service2 struct {
}

func (receiver Service2) getName() string {
	return "Service2"
}

func (receiver Service2) execute() string {
	return "Executing Service2"
}

//步骤 3
//为 JNDI 查询创建 InitialContext。
type Object interface {
}

type InitialContext struct {
}

func (receiver InitialContext) lookup(jndiName string) Object {
	if jndiName == "Service1" {
		return Service1{}
	}
	if jndiName == "Service2" {
		return Service2{}
	}
	return nil
}

//步骤 4
//创建缓存 Cache。

type Cache struct {
	services []Service
}

func (receiver Cache) getService(serviceName string) Service {
	for _, service := range receiver.services {
		if service.getName() == serviceName {
			return service
		}
	}
	return nil
}

func (receiver *Cache) addService(newService Service) {
	for _, service := range receiver.services {
		if service.getName() == newService.getName() {
			return
		}
	}
	receiver.services = append(receiver.services, newService)
}

//步骤 5
//创建服务定位器。
var cache Cache

func getService(jndiName string) (Service, string) {
	service := cache.getService(jndiName)
	if service != nil {
		return service, "Returning cached " + jndiName + " object"
	}
	initialContext := InitialContext{}
	service1 := initialContext.lookup(jndiName).(Service)
	if service1 != nil {
		cache.addService(service1)

	}
	return service1, "Looking up and creating a new " + jndiName + " object"
}

func init() {
	cache = Cache{}
}
