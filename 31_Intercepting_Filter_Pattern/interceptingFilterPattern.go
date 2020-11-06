package _1_Intercepting_Filter_Pattern

//步骤 1
//创建过滤器接口 Filter。
type Filter interface {
	execute(request string) string
}

//步骤 2
//创建实体过滤器。

type AuthenticationFilter struct {
}

func (receiver AuthenticationFilter) execute(request string) string {
	return "Authenticating request: " + request
}

type DebugFilter struct {
}

func (receiver DebugFilter) execute(request string) string {
	return "request log: " + request
}

//步骤 3
//创建 Target。

type Target struct {
}

func (receiver *Target) execute(request string) string {
	return "Executing request: " + request
}

//步骤 4
//创建过滤器链。
type FilterChain struct {
	filters []Filter
	target  Target
}

func (receiver *FilterChain) addFilter(filter Filter) {
	receiver.filters = append(receiver.filters, filter)
}

func (receiver *FilterChain) execute(request string) (ret string) {
	for _, filter := range receiver.filters {
		ret += filter.execute(request)
	}
	ret += receiver.target.execute(request)
	return
}

//步骤 5
//创建过滤管理器。

type FilterManager struct {
	filterChain FilterChain
}

func NewFilterManager(target Target) *FilterManager {
	return &FilterManager{FilterChain{target: target}}
}
func (receiver *FilterManager) setFilter(filter Filter) {
	receiver.filterChain.addFilter(filter)
}

func (receiver *FilterManager) filterRequest(request string) string {
	return receiver.filterChain.execute(request)
}

//步骤 6
//创建客户端 Client。
type Client struct {
	filterManager *FilterManager
}

func (receiver *Client) sendRequest(request string) string {
	return receiver.filterManager.filterRequest(request)
}

func (receiver *Client) setFilterManager(filterManager *FilterManager) {
	receiver.filterManager = filterManager
}
