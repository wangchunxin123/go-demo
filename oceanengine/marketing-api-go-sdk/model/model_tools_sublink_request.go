/* #
# Author: wangchunxin
# Created Time: 2020/12/8 2:49 下午
# File Name:
# Description:
# */
package model

// 获取子链列表
type ToolsSublinkListRequest struct {
	AdvertiserId int64  `json:"advertiser_id ,omitempty"` // 广告主ID
	Content      string `json:"content,omitempty"`        // 子链内容或ID（模糊搜索），不传递则不进行筛选
	Page         int    `json:"page,omitempty"`           // 当前页，默认值为1
	PageSize     int    `json:"page_size,omitempty"`      // 页大小，默认值10，超过50则取默认值
}

// 创建子链
type ToolsSublinkCreateRequest struct {
	SubLinks struct {
		Tag         string `json:"tag,omitempty"`            // 子链标签，标签长子链必须有值，长度限制2-4字
		Info        string `json:"info ,omitempty"`          // 子链内容，长度限制：1.短子链：4-6字；2.长子链：13-27字；3.标签长子链：13-22汉字。
		SubLinkType string `json:"sub_link_type ,omitempty"` // 子链类型: `SHORT` 短子链, `LONG` 长子链, `TAG_LONG` 标签长子链
	} `json:"sub_links ,omitempty"`                      // 新增子链信息，单次创建三种类型子链总和不超过45个
	AdvertiserId int64 `json:"advertiser_id ,omitempty"` // 广告主id
}

// 编辑子链
type ToolsSublinkUpdateRequest struct {
	AdvertiserId int64 `json:"advertiser_id ,omitempty"` // 广告主id
	SubLinks     struct {
		Info      string `json:"info,omitempty"`         // 子链内容，长度限制：1.短子链：4-6字；2.长子链：13-27字；3.标签长子链：13-22汉字。
		SubLinkId int64  `json:"sub_link_id ,omitempty"` // 子链id
		Tag       string `json:"tag,omitempty"`          // 子链标签，长度限制2-4字。tag和info不能同时为空，两者必须一个有值，且tag只能在标签长子链的时候可修改。
	} `json:"sub_links ,omitempty"` // 编辑子链信息，子链编辑的数量、字段长度和格式要求同子链创建
}

// 删除子链
type ToolsSublinkDeleteRequest struct {
	AdvertiserId int64   `json:"advertiser_id ,omitempty"` // 广告主id
	Ids          []int64 `json:"ids ,omitempty"`           // 子链ID列表，最多可以传100个子链ID
}
