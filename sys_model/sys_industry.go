package sys_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/kysion/base-library/base_model"
)

/*
	行业类别
*/

type SysIndustry struct {
	Id           int64   `json:"id"        dc:"ID，ID值为0时则新增菜单" v:"min:0#ID不能小于0"`
	CategoryId   *int64  `json:"categoryId" description:"行业ID"`
	CategoryName *string `json:"categoryName"       description:"行业名称"`
	CategoryDesc *string `json:"categoryDesc"       description:"行业描述"`
	Rate         *int    `json:"rate"       description:"费率"`
	ParentId     *int64  `json:"parentId"  dc:"所属父级" v:"integer|min:0#父级ID参数错误|父级ID不能小于0" default:"0"`
	Sort         *int    `json:"sort"      dc:"排序" v:"integer#排序参数错误"`
	State        *int    `json:"state"    dc:"状态：0隐藏，1显示" v:"in:0,1#请选择状态类型" default:"1"`
}

type UpdateSysIndustry struct {
	Id           int64   `json:"id"        dc:"ID，ID值为0时则新增菜单" v:"min:0#ID不能小于0"`
	CategoryId   *int64  `json:"categoryId" description:"行业ID"`
	CategoryName *string `json:"categoryName"       description:"行业名称"`
	CategoryDesc *string `json:"categoryDesc"       description:"行业描述"`
	Rate         *int    `json:"rate"       description:"费率"`
	ParentId     *int64  `json:"parentId"  dc:"所属父级" v:"integer|min:0#父级ID参数错误|父级ID不能小于0" default:"0"`
	Sort         *int    `json:"sort"      dc:"排序" v:"integer#排序参数错误"`
	State        *int    `json:"state"    dc:"状态：0隐藏，1显示" v:"in:0,1#请选择状态类型" default:"1"`
}

type SysIndustryRes sys_entity.SysIndustry
type SysIndustryListRes base_model.CollectRes[*sys_entity.SysIndustry]

type SysIndustryTreeRes struct {
	*sys_entity.SysIndustry
	Children []*SysIndustryTreeRes `json:"children" dc:"行业类别子级"`
}

type SysIndustryTreeListRes []*SysIndustryTreeRes

// 微信 -----

type Industry struct {
	BigIndustryInfo      BigIndustryInfo        `json:"big_industry_info"`
	IndustryCategoryList []IndustryCategoryList `json:"industry_category_list"`
}
type BigIndustryInfo struct {
	BigIndustryID   int    `json:"big_industry_id"`
	BigIndustryName string `json:"big_industry_name"`
}
type IndustryCategoryInfo struct {
	CategoryID               int             `json:"category_id"`
	CategoryName             string          `json:"category_name"`
	CategoryDesc             string          `json:"category_desc"`
	QualificationsType       int             `json:"qualifications_type"`
	QualificationsOptions    int             `json:"qualifications_options"`
	QualificationsGuide      string          `json:"qualifications_guide"`
	BigIndustryInfo          BigIndustryInfo `json:"big_industry_info"`
	SubjectInformationDepend int             `json:"subject_information_depend"`
	QualificationSample      string          `json:"qualification_sample"`
}
type SettlementRuleInfo struct {
	SettlementRuleID          int         `json:"settlement_rule_id"`
	MerchantCategory          int         `json:"merchant_category"`
	MerchantEntityType        int         `json:"merchant_entity_type"`
	SettlementRuleDesc        string      `json:"settlement_rule_desc"`
	CashFlow                  int         `json:"cash_flow"`
	RateTemplateID            interface{} `json:"rate_template_id"`
	RateDesc                  string      `json:"rate_desc"`
	Rate                      int         `json:"rate"`
	BillingCycleDesc          string      `json:"billing_cycle_desc"`
	LimitType                 int         `json:"limit_type"`
	CreditCardRestrictions    int         `json:"credit_card_restrictions"`
	SpecialMerchantType       int         `json:"special_merchant_type"`
	SettlementPackageID       string      `json:"settlement_package_id"`
	IndustryOverview          string      `json:"industry_overview"`
	IndustryCategory          int         `json:"industry_category"`
	RateDecimal               string      `json:"rate_decimal"`
	MerchantSpecialAttributes int         `json:"merchant_special_attributes"`
}
type IndustryCategoryList struct {
	IndustryCategoryInfo IndustryCategoryInfo `json:"industry_category_info"`
	SettlementRuleInfo   SettlementRuleInfo   `json:"settlement_rule_info"`
}
