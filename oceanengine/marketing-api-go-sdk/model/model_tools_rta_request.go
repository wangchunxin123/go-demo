/* #
# Author: wangchunxin
# Created Time: 2020/12/8 2:47 下午
# File Name:
# Description:
# */
package model

// 绑定RTA策略
type ToolsRtaBindRequest struct {
	AdvertiserId int `json:"advertiser_id ,omitempty"` // 广告主ID
	RtaId        int `json:"rta_id ,omitempty"`        // RTA策略ID
}

// 解绑RTA策略
type ToolsRtaUnbindRequest struct {
	AdvertiserId int `json:"advertiser_id ,omitempty"` // 广告主ID
	RtaId        int `json:"rta_id ,omitempty"`        // RTA策略ID
}

// 获取RTA策略数据
type ToolsRtaGetInfoRequest struct {
	AdvertiserId int `json:"advertiser_id ,omitempty"` // 广告主ID
}

// 获取可用的RTA策略
type ToolsRtaGetRequest struct {
	AdvertiserId int `json:"advertiser_id ,omitempty"` // 广告主ID
}
