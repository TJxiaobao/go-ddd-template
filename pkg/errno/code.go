package errno

// code=0 请求成功
// code=4xx 客户端请求错误
// code=5xx 服务器端错误
// code=2xxxx 业务处理错误码

type Errno struct {
	Code    int
	Message string
}

func (e *Errno) Success() bool {
	return e.Code == OK.Code
}

// 全局通用错误码定义
// 2xx/4xx/5xx
var (
	OK = &Errno{Code: 200, Message: "Success"}

	ErrMissingParameter = &Errno{Code: 400, Message: "Missing parameter %s"}
	ErrParameterInvalid = &Errno{Code: 400, Message: "Invalid parameter %s"}
	ErrUnauthorized     = &Errno{Code: 401, Message: "Unauthorized"}
	ErrPermissionDeby   = &Errno{Code: 403, Message: "Permission deny"}
	ErrRouteNotFound    = &Errno{Code: 404, Message: "Route not found"}

	ErrInternalServer      = &Errno{Code: 500, Message: "Internal server error"}
	ErrInternalServerPanic = &Errno{Code: 500, Message: "Internal server panic occurred"}
	ErrDatabase            = &Errno{Code: 500, Message: "Database error"}
	ErrUnknown             = &Errno{Code: 510, Message: "Unknown error"}
)

// 主账号创建相关错误码定义
// 1101xxx
var (
	ErrUsernameExists        = &Errno{Code: 1101001, Message: "Username already exists"}
	ErrUsernameEmpty         = &Errno{Code: 1101002, Message: "Parameter username cannot be empty"}
	ErrInvalidUsernameFormat = &Errno{Code: 1101003, Message: "Parameter username is invalid"}
	ErrEmailExists           = &Errno{Code: 1101004, Message: "Email already exists"}
	ErrEmailEmpty            = &Errno{Code: 1101005, Message: "Parameter email cannot be empty"}
	ErrInvalidEmailFormat    = &Errno{Code: 1101006, Message: "Parameter email is invalid"}
	ErrPhoneExists           = &Errno{Code: 1101007, Message: "Phone already exists"}
	ErrPhoneEmpty            = &Errno{Code: 1101008, Message: "Parameter phone cannot be empty"}
	ErrInvalidPhoneFormat    = &Errno{Code: 1101009, Message: "Parameter phone is invalid"}
	ErrInvalidBizChannel     = &Errno{Code: 1101010, Message: "Parameter channelId is invalid"}
	ErrPasswordEmpty         = &Errno{Code: 1101011, Message: "Parameter password cannot be empty"}
	ErrInvalidPwdFormat      = &Errno{Code: 1101012, Message: "Parameter password is invalid"}
)

var (
	ErrUidNotFound            = &Errno{Code: 1201001, Message: "Account not found by uid"}
	ErrCreateInstanceQuantity = &Errno{Code: 1201003, Message: "Amount of instance cannot be more than 1"}
	ErrInvalidExpireDate      = &Errno{Code: 1201004, Message: "Parameter expire_date is invalid"}
)

var (
	ErrAccountNotFound = &Errno{Code: 1201001, Message: "Account not found"}
)

var (
	ErrIssueNotFound = &Errno{Code: 1301001, Message: "Issue not found"}
)
