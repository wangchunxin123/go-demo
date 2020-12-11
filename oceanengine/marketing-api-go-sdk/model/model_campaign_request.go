/* #
# Author: wangchunxin
# Created Time: 2020/11/23 10:19 上午
# File Name:
# Description:
# */
package model

type CampaignGetRequest struct {
	AdvertiserId int64                        `json:"advertiser_id"`
	Fields       []string                     `json:"fields,omitempty"`
	Page         int                          `json:"page"`
	PageSize     int                          `json:"page_size"`
	Filtering    *CampaignGetFilteringRequest `json:"filtering,omitempty"`
}

type CampaignGetFilteringRequest struct {
	Ids                []int64 `json:"ids,omitempty"`
	CampaignName       string  `json:"campaign_name,omitempty"`
	LandingType        string  `json:"landing_type,omitempty"`
	Status             string  `json:"status,omitempty"`
	CampaignCreateTime string  `json:"campaign_create_time,omitempty"`
}

//添加
type CampaignAddRequest struct {
	AdvertiserId        int64   `json:"advertiser_id"`
	CampaignName        string  `json:"campaign_name"`
	Operation           string  `json:"operation,omitempty"`
	BudgetMode          string  `json:"budget_mode"`
	Budget              float64 `json:"budget,omitempty"`
	LandingType         string  `json:"landing_type"`
	DeliveryRelationNum string  `json:"delivery_relation_num,omitempty"`
}

//修改
type CampaignUpdateRequest struct {
	AdvertiserId int64   `json:"advertiser_id"`
	CampaignId   int64   `json:"campaign_id"`
	ModifyTime   string  `json:"modify_time"`
	CampaignName string  `json:"campaign_name"`
	BudgetMode   string  `json:"budget_mode"`
	Budget       float64 `json:"budget,omitempty"`
}

//修改状态
type CampaignUpdateStatusRequest struct {
	AdvertiserId int64   `json:"advertiser_id"`
	CampaignIds  []int64 `json:"campaign_ids"`
	OptStatus    string  `json:"opt_status"`
}
