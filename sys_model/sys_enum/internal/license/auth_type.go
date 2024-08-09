package sys_enum_license

import "github.com/kysion/base-library/utility/enum"

// 认证类型：业务层自定义（例如：认证类型：1微商、2普通商户、4企业）

type AuthTypeEnum enum.IEnumCode[int]

type authType struct{}

var AuthType = authType{}

func (e authType) New(code int, description string) AuthTypeEnum {
	return enum.New[AuthTypeEnum](code, description)
}
