package invite_id

import (
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"strconv"
)

func TestFunc() {
	var id int64 = 7063052439322693
	code := InviteIdToCode(id)

	inviteId := CodeToInviteId(code)
	fmt.Println(inviteId == id)
}

// InviteIdToCode 邀约id转换为code: id --> 32位编码 --> 全部转为大写 --> code
func InviteIdToCode(id int64) string {
	// id编码成32进制
	id32 := strconv.FormatInt(id, 32)

	// 转为大写
	idStr := gstr.ToUpper(id32)

	return idStr
}

// CodeToInviteId 邀约code转为原始id：code --> 全部转小写 --> 32位解码 --> id
func CodeToInviteId(idStr string) int64 {
	// 大写转小写
	idLower := gstr.ToLower(idStr)

	// 32进制解码成id
	id64, _ := strconv.ParseInt(idLower, 32, 0)

	return id64
}
