/* #
# Author: wangchunxin
# Created Time: 2020/12/8 2:29 下午
# File Name:
# Description:
# */
package model

// 获取穿山甲流量包
type ToolsUnionFlowPackageGetRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	Filtering    struct {
		FlowPackageIds []int `json:"flow_package_ids,omitempty"` // 按照流量包ID过滤
	} `json:"filtering,omitempty"`            // 过滤字段
	Page     int `json:"page,omitempty"`      // 页数默认值: `1`
	PageSize int `json:"page_size,omitempty"` // 页面大小默认值: `10`
}

// 创建穿山甲流量包
type ToolsUnionFlowPackageCreateRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	Name         string   `json:"name,omitempty"`          // 流量包名称，20 个字符以内
	Rit          []string `json:"rit,omitempty"`           // 穿山甲广告位，最多 500 个
}

// 修改穿山甲流量包
type ToolsUnionFlowPackageUpdateRequest struct {
	AdvertiserId  int64    `json:"advertiser_id,omitempty"`   // 广告主ID
	Name          string   `json:"name,omitempty"`            // 流量包名称，20 个字符以内
	FlowPackageId int      `json:"flow_package_id,omitempty"` // 流量包名称，20 个字符以内
	Rit           []string `json:"rit,omitempty"`             // 穿山甲广告位，最多 500 个
}

// 删除穿山甲流量包
type ToolsUnionFlowPackageDeleteRequest struct {
	AdvertiserId  int64 `json:"advertiser_id,omitempty"`   // 广告主ID
	FlowPackageId int   `json:"flow_package_id,omitempty"` // 流量包ID，20 个字符以内
}

// 查看rit数据
type ToolsUnionFlowPackageReportRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	Filter       struct {
		AdIds     []int64 `json:"ad_ids,omitempty"`     // 广告 id 列表
		StartTime string  `json:"start_time,omitempty"` // 开始时间。格式是"2019-07-24"。范围是 100 天以内，默认是 30 天前。
		EndTime   string  `json:"end_time,omitempty"`   // 结束时间。格式是"2019-07-24",默认是今天。
	} `json:"filter,omitempty"`               // 筛选条件
	Page     int `json:"page,omitempty"`      // 页数默认值: `1`
	PageSize int `json:"page_size,omitempty"` // 页面大小,范围是[1, 100]默认值: `10`
}
