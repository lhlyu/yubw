package code

const (
	Success                            = 0   // 成功
	UnknownError                       = 1   // 未知错误
	ServiceTemporarilyUnavailable      = 2   // 服务暂不可用
	UnsupportedOpenapiMethod           = 3   // 未知的方法
	OpenApiRequestLimitReached         = 4   // 接口调用次数已达到设定的上限
	UnauthorizedClientIpAddress        = 5   // 请求来自未经授权的IP地址
	NoPermissionToAccessUserData       = 6   // 无权限访问该用户数据
	NoPermission                       = 7   // 来自该refer的请求无访问权限
	InvalidParameter                   = 100 // 请求参数无效
	InvalidApiKey                      = 101 // api key无效
	IncorrectSignature                 = 104 // 无效签名
	TooManyParameters                  = 105 // 请求参数过多
	UnsupportedSignatureMethod         = 106 // 未知的签名方法
	InvalidTimestampParameter          = 107 // timestamp参数无效
	InvalidUserInfoField               = 109 // 无效的用户资料字段名
	AccessTokenInvalidOrNoLongerValid  = 110 // 无效的access token
	AccessTokenExpired                 = 111 // access token过期
	UserNotVisible                     = 210 // 用户不可见
	UnsupportedPermission              = 211 // 获取未授权的字段
	NoPermissionToAccessUserEmail      = 212 // 没有权限获取用户的email
	UnknownDataStoreApiError           = 800 // 未知的存储操作错误
	InvalidOperation                   = 801 // 无效的操作方法
	DataStoreAllowableQuotaWasExceeded = 802 // 数据存储空间已超过设定的上限
	SpecifiedObjectCannotBeFound       = 803 // 指定的对象不存在
	SpecifiedObjectAlreadyExists       = 804 // 指定的对象已存在
	DatabaseErrorOccurred              = 805 // 数据库操作出错
	NoSuchApplicationExists            = 900 // 访问的应用不存在
)
