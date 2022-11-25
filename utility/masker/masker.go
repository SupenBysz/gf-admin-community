package masker

import (
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
)

// 数据脱敏

type MaskType interface {
	MaskType() string
}

type maskRule struct {
	matchString string
}

func (mr maskRule) MaskType() string {
	return mr.matchString
}

var (
	MaskPhone     = maskRule{matchString: "required|phone-loose"}
	MaskIPv4      = maskRule{matchString: "required|ip"}
	MaskEmail     = maskRule{matchString: "required|email"}
	MaskTelephone = maskRule{matchString: "required|telephone"}
	Password      = maskRule{matchString: "required|length:1,128"}
	IDCard        = maskRule{matchString: "required|size:18"}
)

func MaskString(in string, maskType MaskType) string {
	if maskType.MaskType() == MaskPhone.MaskType() {
		if len(in) >= 5 {
			return gstr.SubStr(in, 0, 3) + "******" + gstr.SubStr(in, len(in)-2)
		}
		return ""
	}
	if maskType.MaskType() == MaskIPv4.MaskType() {
		err := g.Validator().Bail().Data(in).Rules(maskType.MaskType()).Run(gctx.New())
		if err == nil {
			ipArr := gstr.Split(in, ".")
			return ipArr[0] + ".***.***." + ipArr[3]
		}
		return "---.---.---.---"
	}
	if maskType.MaskType() == MaskEmail.MaskType() {
		if gstr.ContainsI(in, "@") {
			targetArr := strings.Split(in, "@")
			return targetArr[0][0:1] + "******" + targetArr[0][len(targetArr[0])-1:1] + "@" + targetArr[1]
		}
		return "------@***.***"
	}
	if maskType.MaskType() == MaskTelephone.MaskType() {
		if len(in) > 2 {
			return "******" + in[len(in)-2:2]
		}
		return ""
	}
	if maskType.MaskType() == Password.MaskType() {
		return ""
	}
	if maskType.MaskType() == IDCard.MaskType() {
		if nil == g.Validator().Bail().Data(in).Rules(maskType.MaskType()).Run(gctx.New()) {
			return in[0:6] + "**********" + in[len(in)-2:len(in)]
		}
		return ""
	}
	return ""
}
