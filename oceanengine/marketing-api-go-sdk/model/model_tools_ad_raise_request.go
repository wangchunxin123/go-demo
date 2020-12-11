/* #
# Author: wangchunxin
# Created Time: 2020/12/8 2:44 下午
# File Name:
# Description:
# */
package model

// 启动一键起量
type ToolsAdRaiseSetRequest struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 广告主ID
	AdId         int64  `json:"ad_id,omitempty"`         // 广告计划id
	OptType      string `json:"opt_type,omitempty"`      // 操作类型，允许值: `"CLICK_RAISE"` 表示一键起量
	BudgetValue  int    `json:"budget_value,omitempty"`  // 预估值，单位元，取值大于等于0
}

// 获取起量预估值
type ToolsAdRaiseEstimateGetRequest struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 广告主ID
	AdId         int64  `json:"ad_id,omitempty"`         // 广告计划id
	OptType      string `json:"opt_type,omitempty"`      // 操作类型，允许值: `"CLICK_RAISE"` 表示一键起量
	BudgetValue  int    `json:"budget_value,omitempty"`  // 预估值，单位元，取值大于等于0
}

// 获取当前起量状态
type ToolsAdRaiseStatusGetRequest struct {
	AdvertiserId int64   `json:"advertiser_id,omitempty"` // 广告主ID
	AdIds        []int64 `json:"ad_ids,omitempty"`        // 广告计划id 列表，最多1000个
	OptType      string  `json:"opt_type,omitempty"`      // 操作类型，允许值: `"CLICK_RAISE"` 表示一键起量
}

// 获取一键起量的后验数据
type ToolsAdRaiseResultGetRequest struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 广告主ID
	AdId         int64  `json:"ad_id,omitempty"`         // 广告计划id（包含已删除广告计划）
	OptType      string `json:"opt_type,omitempty"`      // 操作类型，允许值: `"CLICK_RAISE"` 表示一键起量
}
