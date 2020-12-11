/* #
# Author: wangchunxin
# Created Time: 2020/12/8 2:28 下午
# File Name:
# Description:
# */
package model

// 创建动态创意词包
type ToolsCreativeWordCreateRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	Name         string   `json:"name,omitempty"`          // 创意词包名称
	DefaultWord  string   `json:"default_word,omitempty"`  // 默认词
	Words        []string `json:"words,omitempty"`         // 替换词
}

// 查询动态创意词包
type ToolsCreativeWordSelectRequest struct {
	AdvertiserId    int64 `json:"advertiser_id,omitempty"`     // 广告主ID
	CreativeWordIds []int `json:"creative_word_ids,omitempty"` // 创意词包id列表，如不填默认返回所有创意词包
}

// 更新动态创意词包
type ToolsCreativeWordUpdateRequest struct {
	AdvertiserId   int64    `json:"advertiser_id,omitempty"`    // 广告主ID
	CreativeWordId int      `json:"creative_word_id,omitempty"` // 创意词包id
	Name           string   `json:"name,omitempty"`             // 创意词包名称
	DefaultWord    string   `json:"default_word,omitempty"`     // 默认词
	Words          []string `json:"words,omitempty"`            // 替换词
}

// 删除动态创意词包
type ToolsCreativeWordDeleteRequest struct {
	AdvertiserId   int64 `json:"advertiser_id,omitempty"`    // 广告主ID
	CreativeWordId int   `json:"creative_word_id,omitempty"` // 创意词包id
}
