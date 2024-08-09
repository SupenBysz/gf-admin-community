package sys_api

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetIndustryByIdReq struct {
	g.Meta `path:"/getIndustryById" method:"post" summary:"根据行业ID获取行业|信息" tags:"行业类别"`
	Id     int64 `json:"id" v:"required#行业ID校验失败" dc:"行业ID"`
}

type CreateIndustryReq struct {
	g.Meta `path:"/createIndustry" method:"post" summary:"创建行业" tags:"行业类别"`
	sys_model.SysIndustry
}

type UpdateIndustryReq struct {
	g.Meta `path:"/updateIndustry" method:"post" summary:"更新行业" tags:"行业类别"`
	sys_model.UpdateSysIndustry
}

type DeleteIndustryReq struct {
	g.Meta `path:"/deleteIndustry" method:"post" summary:"删除行业" tags:"行业类别"`
	Id     int64 `json:"id" v:"required#行业ID校验失败" dc:"行业ID"`
}

type GetIndustryTreeReq struct {
	g.Meta `path:"/getIndustryTree" method:"post" summary:"获取行业树" tags:"行业类别"`
	Id     int64 `json:"id" v:"required#行业ID校验失败" dc:"行业ID"`
}

type ImportIndustryReq struct {
	g.Meta `path:"/importIndustry" method:"post" summary:"导入行业" tags:"行业类别"`
	List   []Info `json:"list"`
}

//type List []Info

type Info struct {
	BigIndustryInfo struct {
		BigIndustryId   int    `json:"big_industry_id"`
		BigIndustryName string `json:"big_industry_name"`
	} `json:"big_industry_info"`
	IndustryCategoryList []struct {
		IndustryCategoryInfo struct {
			CategoryId            int    `json:"category_id"`
			CategoryName          string `json:"category_name"`
			CategoryDesc          string `json:"category_desc"`
			QualificationsType    int    `json:"qualifications_type"`
			QualificationsOptions int    `json:"qualifications_options"`
			QualificationsGuide   string `json:"qualifications_guide"`
			BigIndustryInfo       struct {
				BigIndustryId   int         `json:"big_industry_id"`
				BigIndustryName interface{} `json:"big_industry_name"`
			} `json:"big_industry_info"`
			SubjectInformationDepend int    `json:"subject_information_depend"`
			QualificationSample      string `json:"qualification_sample"`
		} `json:"industry_category_info"`
		SettlementRuleInfo struct {
			SettlementRuleId          int         `json:"settlement_rule_id"`
			MerchantCategory          int         `json:"merchant_category"`
			MerchantEntityType        int         `json:"merchant_entity_type"`
			SettlementRuleDesc        string      `json:"settlement_rule_desc"`
			CashFlow                  int         `json:"cash_flow"`
			RateTemplateId            interface{} `json:"rate_template_id"`
			RateDesc                  string      `json:"rate_desc"`
			Rate                      int         `json:"rate"`
			BillingCycleDesc          string      `json:"billing_cycle_desc"`
			LimitType                 int         `json:"limit_type"`
			CreditCardRestrictions    int         `json:"credit_card_restrictions"`
			SpecialMerchantType       int         `json:"special_merchant_type"`
			SettlementPackageId       string      `json:"settlement_package_id"`
			IndustryOverview          string      `json:"industry_overview"`
			IndustryCategory          int         `json:"industry_category"`
			RateDecimal               string      `json:"rate_decimal"`
			MerchantSpecialAttributes int         `json:"merchant_special_attributes"`
		} `json:"settlement_rule_info"`
	} `json:"industry_category_list"`
}
