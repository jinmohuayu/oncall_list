package web

// Result 结果
type Result struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var (
	// ResultOK 成功
	ResultOK = Result{Success: true}
)

// successResult 成功
func successResult(data interface{}) Result {
	return Result{Success: true, Data: data}
}

// errorResult 错误
func errorResult(err error) Result {
	return Result{Success: false, Message: err.Error()}
}

// errorMessageResult 错误
func errorMessageResult(message string) Result {
	return Result{Success: false, Message: message}
}
