/* #
# Author: wangchunxin
# Created Time: 2020/12/8 2:25 下午
# File Name:
# Description:
# */
package model

// 获取兴趣关键词词包
type ToolsInterestTagsSelectRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
}

// 创建兴趣词包
type ToolsInterestTagsCreateRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	Name         string   `json:"name,omitempty"`          // 兴趣词包名称
	Words        []string `json:"words,omitempty"`         // 兴趣词包具体内容
}

// 更新兴趣词包
type ToolsInterestTagsUpdateRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	Id           int      `json:"id,omitempty"`            // 兴趣词包id
	Name         string   `json:"name,omitempty"`          // 兴趣词包名称
	Words        []string `json:"words,omitempty"`         // 兴趣词包具体内容
}

// 删除兴趣词包
type ToolsInterestTagsDeleteRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	Id           int   `json:"id,omitempty"`            // 兴趣词包id
}

// 兴趣关键词ID转词
type ToolsInterestTagsId2WordRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	WordIds      []string `json:"word_ids,omitempty"`      // 兴趣关键词id
}

// 兴趣关键词词转ID
type ToolsInterestTagsWord2IdRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	Words        []string `json:"words,omitempty"`         // 兴趣关键词
}
