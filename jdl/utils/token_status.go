package utils

type TokenStatus int

const (
	TokenFailed     TokenStatus = iota
	TokenValid                  // token 有效，时间待确认
	TokenValidGtTen             // token 有效且剩余时间大于十天
	TokenValidLtTen             // token有效但剩余时间小于十天
	TokenExpired                // token 已过期，重新获取code并授权
	TokenMemoryNil
	TokenRedisNil // token 为空。获取code并授权
)

func (t TokenStatus) String() string {
	switch t {
	case TokenFailed:
		return "TokenFailed"
	case TokenValid:
		return "TokenValid"
	case TokenValidGtTen:
		return "TokenValidGtTen"
	case TokenValidLtTen:
		return "TokenValidLtTen"
	case TokenExpired:
		return "TokenExpired"
	case TokenMemoryNil:
		return "TokenMemoryNil"
	case TokenRedisNil:
		return "TokenRedisNil"
	default:
		return "TokenOtherFailed"
	}
}
