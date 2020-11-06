package _0_Front_Controller_Pattern

//步骤 1
//创建视图。
type HomeView struct {
}

func (receiver *HomeView) show() string {
	return "Displaying Home Page"
}

type StudentView struct {
}

func (receiver *StudentView) show() string {
	return "Displaying Student Page"
}

//步骤 2
//创建调度器 Dispatcher。

type Dispatcher struct {
	studentView StudentView
	homeView    HomeView
}

func (receiver *Dispatcher) dispatch(request string) string {
	if request == "STUDENT" {
		return receiver.studentView.show()
	} else {
		return receiver.homeView.show()
	}
}

//步骤 3
//创建前端控制器 FrontController。

type FrontController struct {
	Dispatcher
}

//对用户进行身份验证
func (receiver FrontController) isAuthenticUser() bool {
	return true
}

//记录每一个请求
func (receiver FrontController) trackRequest() {
}

func (receiver FrontController) dispatchRequest(request string) string {
	receiver.trackRequest()
	if receiver.isAuthenticUser() {
		return receiver.dispatch(request)
	}
	return ""
}
