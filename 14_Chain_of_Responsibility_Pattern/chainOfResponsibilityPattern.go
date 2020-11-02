package _4_Chain_of_Responsibility_Pattern

//步骤 1
//创建抽象的记录器类。
const (
	INFO  = 1
	DEBUG = 2
	ERROR = 3
)

type AbstractLogger interface {
	logMessage(int, string) string
}

//步骤 2
//创建扩展了该记录器类的实体类。
type ConsoleLogger struct {
	nextLogger AbstractLogger
	level      int
}

func (receiver *ConsoleLogger) logMessage(level int, message string) (ret string) {
	if receiver.level <= level {
		ret += "Console::Logger: " + message
	}
	if receiver.nextLogger != nil {
		ret += receiver.nextLogger.logMessage(level, message)
	}
	return
}

type FileLogger struct {
	nextLogger AbstractLogger
	level      int
}

func (receiver *FileLogger) logMessage(level int, message string) (ret string) {
	if receiver.level <= level {
		ret += "File::Logger: " + message
	}
	if receiver.nextLogger != nil {
		ret += receiver.nextLogger.logMessage(level, message)
	}
	return
}

type ErrorLogger struct {
	nextLogger AbstractLogger
	level      int
}

func (receiver *ErrorLogger) logMessage(level int, message string) (ret string) {
	if receiver.level <= level {
		ret += "Error::Logger: " + message
	}
	if receiver.nextLogger != nil {
		ret += receiver.nextLogger.logMessage(level, message)
	}
	return
}
