/* #
# Author: wangchunxin
# Created Time: 2020/12/7 6:20 下午
# File Name:
# Description:
# */
package model

// 获取计划诊断详情
type ToolsDiagnosisAdGetRequest struct {
	AdvertiserId int64   `json:"advertiser_id,omitempty"` // 广告主ID
	AdIds        []int64 `json:"ad_ids,omitempty"`        // 广告计划ID数组，最多100个
}

// 获取计划诊断预估变化趋势
type ToolsDiagnosisAdCurveRequest struct {
	AdvertiserId  int64  `json:"advertiser_id,omitempty"`  // 广告主ID
	AdId          int64  `json:"ad_id,omitempty"`          // 广告计划ID
	InventoryType string `json:"inventory_type,omitempty"` // 投放位置，INVENTORY_ALL为多平台聚合，其他详见[【附录-投放位置】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值：`"INVENTORY_ALL"`,`"INVENTORY_FEED"`,`"INVENTORY_HOTSOON_FEED"`,`"INVENTORY_AWEME_FEED"`,`"INVENTORY_VIDEO_FEED"`
	Type          string `json:"type,omitempty"`           // 诊断类型，ctr:创意质量、cvr:落地页质量、ecpm:ECPM、bid:预算type为ctr：可展示最近48小时内该计划和标杆计划的预估点击率变化趋势，以反映该计划的创意质量度在相竞争的计划中是否具有优势。type为cvr：展示最近48小时内该计划和标杆计划的预估转化率变化趋势，以反映该计划的落地页质量度在相竞争的计划中是否具有优势。type为ecpm：展示该计划所在行业的ECPM均值与预估展示量之间的关系，可根据该曲线预估不同的ECPM值对应的展示量。type为bid：展示最近48小时内该计划和标杆计划的实时出价变化趋势，以反映该计划的出价在相竞争的计划中是否具有优势。允许值: "ctr", "cvr", "ecpm", "bid"
}
