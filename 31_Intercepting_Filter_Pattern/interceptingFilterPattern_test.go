package _1_Intercepting_Filter_Pattern

import "testing"

//步骤 7
//使用 Client 来演示拦截过滤器设计模式。
func TestInterceptingFilterPattern(t *testing.T) {
	filterManager := NewFilterManager(Target{})
	filterManager.setFilter(AuthenticationFilter{})
	filterManager.setFilter(DebugFilter{})
	client := Client{}
	client.setFilterManager(filterManager)
	want := "Authenticating request: HOMErequest log: HOMEExecuting request: HOME"
	if got := client.sendRequest("HOME"); got != want {
		t.Errorf("dispatchRequest() = %v, want %v", got, want)
	}

}
