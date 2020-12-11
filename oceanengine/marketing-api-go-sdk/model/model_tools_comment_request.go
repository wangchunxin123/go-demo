/* #
# Author: wangchunxin
# Created Time: 2020/12/8 11:59 上午
# File Name:
# Description:
# */
package model

// 查询评论
type ToolsCommentGetRequest struct {
	AdvertiserId  int64   `json:"advertiser_id,omitempty"`  // 广告主ID
	AdIds         []int64 `json:"ad_ids,omitempty"`         // 计划id列表，一次最多10个
	InventoryType string  `json:"inventory_type,omitempty"` // 广告位允许值：`"INVENTORY_AWEME_FEED"`（抖音）
	StartTime     string  `json:"start_time,omitempty"`     // 查询起始时间，格式：yyyy-MM-dd，若不填，默认6天前（即获取最近七天的内容）
	EndTime       string  `json:"end_time,omitempty"`       // 查询截止时间，格式：yyyy-MM-dd，若不填，默认当天
	Page          int     `json:"page,omitempty"`           // 页数默认值: `1`
	PageSize      int     `json:"page_size,omitempty"`      // 页面大小默认值: `10`
}

// 评论操作
type ToolsCommentOperateRequest struct {
	AdvertiserId int64   `json:"advertiser_id,omitempty"` // 广告主ID
	CommentIds   []int64 `json:"comment_ids,omitempty"`   // 评论id列表operate_type类型为`"REPLY_TO_REPLY"`、`"STICK"`、`"CANCEL_STICK"`时只允许传入1个评论ID；operate_type
	// 类型为`"REPLY时"`时，只允许传入小于等于20个评论ID;operate_type类型为`"HIDE"`时，只允许传入小于等于50个评论ID
	InventoryType string `json:"inventory_type,omitempty"` // 广告位，允许值：`"INVENTORY_AWEME_FEED"`（抖音）
	OperateType   string `json:"operate_type,omitempty"`   // 操作类型，详情见[【附录-评论操作类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值：`"REPLY"`、`"REPLY_TO_REPLY"`、`"STICK_ON_TOP"`、`"CANCEL_STICK"`、`"HIDE"`、`"BLOCK_USERS"`（屏蔽用户目前只支持抖音）
	ReplyText     string `json:"reply_text,omitempty"`     // 回复内容，当operate_type为`"REPLY"`，`"REPLY_TO_REPLY"`时必填
	ReplyId       int    `json:"reply_id,omitempty"`       // 回复的id，回复二级评论时传入, 用户可从reply_infos中获取，当operate_type为`"REPLY_TO_REPLY"`时必填
}

// 评论回复获取
type ToolsCommentReplyGetRequest struct {
	AdvertiserId  int64  `json:"advertiser_id,omitempty"`  // 广告主ID
	CommentId     int64  `json:"comment_id,omitempty"`     // 父评论id
	InventoryType string `json:"inventory_type,omitempty"` // 广告位允许值：`"INVENTORY_AWEME_FEED"`（抖音）
	Page          int    `json:"page,omitempty"`           // 页数默认值: `1`
	PageSize      int    `json:"page_size,omitempty"`      // 页面大小默认值: `10`
}

// 删除屏蔽词
type ToolsCommentTermsBannedDeleteRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	Terms        []string `json:"terms,omitempty"`         // 待删除的屏蔽词列表如果删除的屏蔽词不存在也会显示成功一次最多操作`500`个词，单个屏蔽词长度范围为`0-100`
}

// 获取屏蔽词列表
type ToolsCommentTermsBannedGetRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	Page         int   `json:"page,omitempty"`          // 页数默认值: 1
	PageSize     int   `json:"page_size,omitempty"`     // 页面大小范围：`1-500`，默认值: `10`
}

// 添加屏蔽词
type ToolsCommentTermsBannedAddRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	Terms        []string `json:"terms,omitempty"`         // 待添加的屏蔽词列表若添加的屏蔽词已存在，不会再次新增，同一个屏蔽词只会在屏蔽词中记录一次一次最多添加`500`个，单个屏蔽词长度范围为`0-100`
}

// 更新屏蔽词
type ToolsCommentTermsBannedUpdateRequest struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 广告主ID
	OriginTerms  string `json:"origin_terms,omitempty"`  // 待更新的屏蔽词屏蔽词长度范围为`0-100`
	NewTerms     string `json:"new_terms,omitempty"`     // 更新后的屏蔽词如果new_terms已存在，则等同于删除origin_terms；如果origin_terms不存在，则等同于新增new_terms屏蔽词长度范围为`0-100`
}

// 添加屏蔽用户
type ToolsAwemeBannedCreateRequest struct {
	AdvertiserId     int64    `json:"advertiser_id ,omitempty"`    // 广告主id
	BannedType       string   `json:"banned_type ,omitempty"`      // 屏蔽类型，允许值：`CUSTOM_TYPE`：自定义规则，根据昵称关键词屏蔽；`AWEME_TYPE`：根据抖音id屏蔽。
	AwemeIds         []string `json:"aweme_ids,omitempty"`         // 抖音id列表，抖音id为抖音页面展示id；当banned_type为`AWEME_TYPE`时必填；每次最多传50个抖音id，抖音id限制长度20，纯数字id不能以0开头；最多屏蔽5000个抖音ID。
	NicknameKeywords []string `json:"nickname_keywords,omitempty"` // 抖音昵称关键词列表，当banned_type为`CUSTOM_TYPE`时必填；关键词长度不大于40个字符，中文算2个字符；每次最多传50个关键词；最多屏蔽5000个关键词。
}

// 删除屏蔽用户
type ToolsAwemeBannedDeleteRequest struct {
	AdvertiserId     int64    `json:"advertiser_id ,omitempty"`    // 广告主id
	BannedType       string   `json:"banned_type ,omitempty"`      // 屏蔽类型，允许值：`CUSTOM_TYPE`：自定义规则，根据昵称关键词屏蔽；`AWEME_TYPE`：根据抖音id屏蔽。
	AwemeIds         []string `json:"aweme_ids,omitempty"`         // 抖音id列表，抖音id为抖音页面展示id；当banned_type为`AWEME_TYPE`时必填；每次最多传50个抖音id，抖音id限制长度20，纯数字id不能以0开头。
	NicknameKeywords []string `json:"nickname_keywords,omitempty"` // 抖音昵称关键词列表，当banned_type为`CUSTOM_TYPE`时必填；关键词长度不大于40个字符，中文算2个字符；每次最多传50个关键词。
}

// 获取屏蔽用户列表
type ToolsAwemeBannedListRequest struct {
	AdvertiserId int64 `json:"advertiser_id ,omitempty"` // 广告主id
	Filtering    struct {
		BannedType      string `json:"banned_type,omitempty"`      // 屏蔽类型，允许值：`CUSTOM_TYPE`：自定义规则，根据昵称关键词屏蔽；`AWEME_TYPE`：根据抖音id屏蔽。
		AwemeId         string `json:"aweme_id,omitempty"`         // 抖音id，抖音id限制长度20，纯数字id不能以0开头。
		NicknameKeyword string `json:"nickname_keyword,omitempty"` // 昵称关键词，关键词长度不大于40个字符，中文算2个字符。
	} `json:"filtering,omitempty"`            // 筛选条件
	Page     int `json:"page,omitempty"`      // 页码，默认值：`1`
	PageSize int `json:"page_size,omitempty"` // 页面大小，取值范围`1~50`，默认值：`10`
}
