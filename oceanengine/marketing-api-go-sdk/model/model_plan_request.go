/* #
# Author: wangchunxin
# Created Time: 2020/12/2 3:55 下午
# File Name:
# Description:
# */
package model

//获取计划的参数
type PlanGetRequest struct {
	AdvertiserId int64          `json:"advertiser_id"`
	Fields       []string       `json:"fields,omitempty"`
	Page         int            `json:"page"`
	PageSize     int            `json:"page_size"`
	Filtering    *PlanGetFilter `json:"filtering,omitempty"`
}

//获取计划的filter
type PlanGetFilter struct {
	Ids          []int64  `json:"ids,omitempty"`
	AdName       string   `json:"ad_name,omitempty"`
	PricingList  []string `json:"pricing_list,omitempty"`
	Status       string   `json:"status,omitempty"`
	CampaignId   int64    `json:"campaign_id,omitempty"`
	AdCreateTime string   `json:"ad_create_time,omitempty"`
	AdModifyTime string   `json:"ad_modify_time,omitempty"`
}

//创建计划
type PlanAddRequest struct {
}

//更新计划状态
type PlanUpdateRequest struct {
	AdvertiserId int64   `json:"advertiser_id"`
	AdIds        []int64 `json:"ad_ids"`
	OptStatus    string  `json:"opt_status"`
}

//修改计划预算
type PlanUpdateBudgetRequest struct {
	AdvertiserId int64                         `json:"advertiser_id"`
	Data         []PlanUpdateBudgetDataRequest `json:"data"`
}

type PlanUpdateBudgetDataRequest struct {
	AdId   int64   `json:"ad_id"`
	Budget float64 `json:"budget"`
}

//修改计划出价
type PlanUpdateBidRequest struct {
	AdvertiserId int64                      `json:"advertiser_id"`
	Data         []PlanUpdateBidDataRequest `json:"data"`
}
type PlanUpdateBidDataRequest struct {
	AdId int64   `json:"ad_id"`
	Bid  float64 `json:"bid"`
}

//获取计划审核建议
type PlanRejectReasonRequest struct {
	AdvertiserId int64   `json:"advertiser_id"`
	AdIds        []int64 `json:"ad_ids"`
}
