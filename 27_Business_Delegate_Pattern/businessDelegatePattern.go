package _7_Business_Delegate_Pattern

//步骤 1
//创建 BusinessService 接口。

type BusinessService interface {
	doProcessing() string
}

//步骤 2
//创建实体服务类。

type EJBService struct {
}

func (receiver *EJBService) doProcessing() string {
	return "Processing task by invoking EJB Service"
}

type JMSService struct {
}

func (receiver *JMSService) doProcessing() string {
	return "Processing task by invoking JMS Service"
}

//步骤 3
//创建业务查询服务。
type BusinessLookUp struct {
}

func (receiver *BusinessLookUp) getBusinessService(serviceType string) BusinessService {
	if serviceType == "EJB" {
		return &EJBService{}
	} else {
		return &JMSService{}
	}
}

//步骤 4
//创建业务代表。
type BusinessDelegate struct {
	lookupService   BusinessLookUp
	businessService BusinessService
	serviceType     string
}

func (receiver *BusinessDelegate) doTask() string {
	receiver.businessService = receiver.lookupService.getBusinessService(receiver.serviceType)
	return receiver.businessService.doProcessing()
}

//步骤 5
//创建客户端。

type Client struct {
	*BusinessDelegate
}
