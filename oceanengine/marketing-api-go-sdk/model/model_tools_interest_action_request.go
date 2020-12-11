/* #
# Author: wangchunxin
# Created Time: 2020/12/8 2:23 下午
# File Name:
# Description:
# */
package model

// 行为类目查询
type ToolsInterestActionActionCategoryRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	ActionScene  []string `json:"action_scene,omitempty"`  // 行为场景，详见[【附录-行为场景】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"E-COMMERCE"`,`"NEWS"`,`"APP"`
	ActionDays   int      `json:"action_days,omitempty"`   // 行为天数默认值: `7`,`15`,`30`,`60`,`90`,`180`,`365`
}

// 行为关键词查询
type ToolsInterestActionActionKeywordRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	QueryWords   string   `json:"query_words,omitempty"`   // 关键词
	ActionScene  []string `json:"action_scene,omitempty"`  // 行为场景，详见[【附录-行为场景】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值：`E-COMMERCE`, `NEWS`, `APP`
	ActionDays   int      `json:"action_days,omitempty"`   // 行为天数默认值: `7`,`15`,`30`,`60`,`90`,`180`,`365`
}

// 兴趣类目查询
type ToolsInterestActionInterestCategoryRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
}

// 兴趣关键词查询
type ToolsInterestActionInterestKeywordRequest struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 广告主ID
	QueryWords   string `json:"query_words,omitempty"`   // 关键词
}

// 兴趣行为类目关键词id转词
type ToolsInterestActionId2WordRequest struct {
	AdvertiserId  int64    `json:"advertiser_id,omitempty"`  // 广告主ID
	Ids           []int    `json:"ids,omitempty"`            // 类目或关键词id列表
	TagType       string   `json:"tag_type,omitempty"`       // 查询类型：类目还是关键词允许值：`CATEGORY`（类目）、`KEYWORD`（关键词）
	TargetingType string   `json:"targeting_type,omitempty"` // 查询目标：兴趣还是行为允许值：`ACTION`（行为）、`INTEREST`（兴趣）
	ActionScene   []string `json:"action_scene,omitempty"`   // 行为场景，查询目标为行为时必填，兴趣不生效允许值：`E-COMMERCE`、`NEWS`、`APP`
	ActionDays    int      `json:"action_days,omitempty"`    // 行为天数，查询目标为行为时必填，兴趣不生效允许值：`7`, `15`, `30`, `60`, `90`, `180`, `365`
}

// 获取行为兴趣推荐关键词
type ToolsInterestActionKeywordSuggestRequest struct {
	AdvertiserId  int64    `json:"advertiser_id,omitempty"`  // 广告主ID
	Id            int      `json:"id,omitempty"`             // 类目或关键词id
	TagType       string   `json:"tag_type,omitempty"`       // 查询类型：类目还是关键词允许值：`CATEGORY`（类目）、`KEYWORD`（关键词）
	TargetingType string   `json:"targeting_type,omitempty"` // 查询目标：兴趣还是行为允许值：`ACTION`（行为）、`INTEREST`（兴趣）
	ActionScene   []string `json:"action_scene,omitempty"`   // 行为场景，查询目标为行为时必填，兴趣不生效允许值：`E-COMMERCE`、`NEWS`、`APP`
	ActionDays    int      `json:"action_days,omitempty"`    // 行为天数，查询目标为行为时必填，兴趣不生效允许值：`7`, `15`, `30`, `60`, `90`, `180`, `365`
}
