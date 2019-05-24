package i18n

import (
	"errorCode/code"
)

// 用脚本一键生成的，绝不手打。。
var ZHCNCodeMap = map[int]string{
	code.Success:                            "成功",
	code.UnknownError:                       "未知错误",
	code.ServiceTemporarilyUnavailable:      "服务暂不可用",
	code.UnsupportedOpenapiMethod:           "未知的方法",
	code.OpenApiRequestLimitReached:         "接口调用次数已达到设定的上限",
	code.UnauthorizedClientIpAddress:        "请求来自未经授权的IP地址",
	code.NoPermissionToAccessUserData:       "无权限访问该用户数据",
	code.NoPermission:                       "来自该refer的请求无访问权限",
	code.InvalidParameter:                   "请求参数无效",
	code.InvalidApiKey:                      "api key无效",
	code.IncorrectSignature:                 "无效签名",
	code.TooManyParameters:                  "请求参数过多",
	code.UnsupportedSignatureMethod:         "未知的签名方法",
	code.InvalidTimestampParameter:          "timestamp参数无效",
	code.InvalidUserInfoField:               "无效的用户资料字段名",
	code.AccessTokenInvalidOrNoLongerValid:  "无效的access token",
	code.AccessTokenExpired:                 "access token过期",
	code.UserNotVisible:                     "用户不可见",
	code.UnsupportedPermission:              "获取未授权的字段",
	code.NoPermissionToAccessUserEmail:      "没有权限获取用户的email",
	code.UnknownDataStoreApiError:           "未知的存储操作错误",
	code.InvalidOperation:                   "无效的操作方法",
	code.DataStoreAllowableQuotaWasExceeded: "数据存储空间已超过设定的上限",
	code.SpecifiedObjectCannotBeFound:       "指定的对象不存在",
	code.SpecifiedObjectAlreadyExists:       "指定的对象已存在",
	code.DatabaseErrorOccurred:              "数据库操作出错",
	code.NoSuchApplicationExists:            "访问的应用不存在",
}
