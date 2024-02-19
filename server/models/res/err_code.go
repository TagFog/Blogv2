package res

type ErrorCode int

const (
	SettingsError ErrorCode = 1001 //系统错误
)

var (
	// 映射code
	ErrorMap = map[ErrorCode]string{
		SettingsError:"系统错误",
	}
)