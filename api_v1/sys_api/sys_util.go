package sys_api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*
	第三方能力：
	 - nlp 自然语言处理
	 - 实名核身鉴权
     - 活体人脸核身
     - 人脸核身
*/

// *********************************** 工具 *****************************************

type GetNerChEcomReq struct {
	g.Meta `path:"/getNerChEcom" method:"POST" tags:"工具/中文分词" summary:"中文分词"`
	Text   string `json:"text" dc:"需要分词的内容"`
}

type GetAiSummaryReq struct {
	g.Meta     `path:"/getSummary" method:"POST" tags:"工具/文字分析" summary:"AI文字分析总结"`
	Text       string `json:"text" dc:"需要分析的内容"`
	Identifier string `json:"identifier" dc:"智能体标识符【业务层自定义】，例如：appbuilder_bot_1 家访分析总结、appbuilder_bot_2 谈心谈话分析总结、appbuilder_bot_4 学生情况分析总结、appbuilder_bot_8 学生学习情况分析、appbuilder_bot_16 学生品行情况分析"`
}
