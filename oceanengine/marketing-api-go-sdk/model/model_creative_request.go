/* #
# Author: wangchunxin
# Created Time: 2020/12/3 9:54 上午
# File Name:
# Description:
# */
package model

//获取创意的参数
type CreativeGetRequest struct {
	AdvertiserId int64                        `json:"advertiser_id"`
	Fields       []string                     `json:"fields,omitempty"`
	Page         int                          `json:"page"`
	PageSize     int                          `json:"page_size"`
	Filtering    *CreativeGetFilteringRequest `json:"filtering,omitempty"`
}

type CreativeGetFilteringRequest struct {
	CampaignId         int64   `json:"campaign_id,omitempty"`
	AdId               int64   `json:"ad_id,omitempty"`
	CreativeIds        []int64 `json:"creative_ids,omitempty"`
	CreativeTitle      string  `json:"creative_title,omitempty"`
	LandingType        string  `json:"landing_type,omitempty"`
	Pricing            string  `json:"pricing,omitempty"`
	Status             string  `json:"status,omitempty"`
	ImageMode          string  `json:"image_mode,omitempty"`
	CreativeCreateTime string  `json:"creative_create_time,omitempty"`
	CreativeModifyTime string  `json:"creative_modify_time,omitempty"`
}

//获取详情的参数
type CreativeReadRequest struct {
	AdvertiserId int64 `json:"advertiser_id"`
	AdId         int64 `json:"ad_id"`
}

//更改创意状态的参数
type CreativeUpdateStatusRequest struct {
	AdvertiserId int64   `json:"advertiser_id"`
	CreativeIds  []int64 `json:"creative_ids"`
	OptStatus    string  `json:"opt_status"`
}

//创意素材信息
type CreativeMaterialReadRequest struct {
	AdvertiserId int64    `json:"advertiser_id"`
	CreativeIds  []int64  `json:"creative_ids"`
	Fields       []string `json:"fields,omitempty"`
}

//获取创意审核建议的参数
type CreativeRejectReasonRequest struct {
	AdvertiserId int64   `json:"advertiser_id"`
	CreativeIds  []int64 `json:"creative_ids"`
}
