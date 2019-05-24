package i18n

import (
	"errorCode/code"
)

// 用脚本一键生成的，绝不手打。。
var ESUNCodeMap = map[int]string{
	code.Success:                            "Success",
	code.UnknownError:                       "Unknown error",
	code.ServiceTemporarilyUnavailable:      "Service temporarily unavailable",
	code.UnsupportedOpenapiMethod:           "Unsupported openapi method",
	code.OpenApiRequestLimitReached:         "Open api request limit reached",
	code.UnauthorizedClientIpAddress:        "Unauthorized client IP address",
	code.NoPermissionToAccessUserData:       "No permission to access user data",
	code.NoPermission:                       "No permission",
	code.InvalidParameter:                   "Invalid parameter",
	code.InvalidApiKey:                      "Invalid API key",
	code.IncorrectSignature:                 "Incorrect signature",
	code.TooManyParameters:                  "Too many parameters",
	code.UnsupportedSignatureMethod:         "Unsupported signature method",
	code.InvalidTimestampParameter:          "Invalid timestamp parameter",
	code.InvalidUserInfoField:               "Invalid user info field",
	code.AccessTokenInvalidOrNoLongerValid:  "Access token invalid or no longer valid",
	code.AccessTokenExpired:                 "Access token expired",
	code.UserNotVisible:                     "User not visible",
	code.UnsupportedPermission:              "Unsupported permission",
	code.NoPermissionToAccessUserEmail:      "No permission to access user email",
	code.UnknownDataStoreApiError:           "Unknown data store API error",
	code.InvalidOperation:                   "Invalid operation",
	code.DataStoreAllowableQuotaWasExceeded: "Data store allowable quota was exceeded",
	code.SpecifiedObjectCannotBeFound:       "Specified object cannot be found",
	code.SpecifiedObjectAlreadyExists:       "Specified object already exists",
	code.DatabaseErrorOccurred:              "database error occurred",
	code.NoSuchApplicationExists:            "No such application exists",
}
