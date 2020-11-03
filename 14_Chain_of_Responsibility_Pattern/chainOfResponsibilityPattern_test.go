package _4_Chain_of_Responsibility_Pattern

import (
	"testing"
)

//步骤 3
//创建不同类型的记录器。赋予它们不同的错误级别，并在每个记录器中设置下一个记录器。每个记录器中的下一个记录器代表的是链的一部分。
func TestChainOfResponsibilityPattern(t *testing.T) {
	errorLogger := ErrorLogger{level: ERROR}
	fileLogger := FileLogger{level: DEBUG}
	consoleLogger := ConsoleLogger{level: INFO}
	errorLogger.nextLogger = &fileLogger
	fileLogger.nextLogger = &consoleLogger
	loggerChain := &errorLogger

	tests := []struct {
		name    string
		level   int
		message string
		wantRet string
	}{
		{"information", INFO, "This is an information.",
			"Console::Logger: This is an information."},
		{"debug", DEBUG, "This is a debug level information.",
			"File::Logger: This is a debug level information." +
				"Console::Logger: This is a debug level information."},
		{"error", ERROR, "This is an error information.",
			"Error::Logger: This is an error information." +
				"File::Logger: This is an error information." +
				"Console::Logger: This is an error information."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := loggerChain.logMessage(tt.level, tt.message); gotRet != tt.wantRet {
				t.Errorf("logMessage() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
