/* #
# Author: wangchunxin
# Created Time: 2020/12/7 5:16 下午
# File Name:
# Description:
# */
package model

// 查询抖音帐号和类目信息
type ToolsAwemeInfoSearchRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	QueryWord    string   `json:"query_word,omitempty"`    // 搜索词
	Behaviors    []string `json:"behaviors,omitempty"`     // 抖音用户行为类型, 详见[【附录-抖音用户行为类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"FOLLOWED_USER"`,`"COMMENTED_USER"`,`"LIKED_USER"`,`"SHARED_USER"`
}

// 查询抖音类似帐号
type ToolsAwemeSimilarAuthorSearchRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	AwemeId      string   `json:"aweme_id,omitempty"`      // 抖音id
	Behaviors    []string `json:"behaviors,omitempty"`     // 抖音用户行为类型, 详见[【附录-抖音用户行为类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"FOLLOWED_USER"`,`"COMMENTED_USER"`,`"LIKED_USER"`,`"SHARED_USER"`
}

// 查询抖音类目列表
type ToolsAwemeMultiLevelCategoryGetRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	Behaviors    []string `json:"behaviors,omitempty"`     // 抖音用户行为类型, 详见[【附录-抖音用户行为类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"FOLLOWED_USER"`,`"COMMENTED_USER"`,`"LIKED_USER"`,`"SHARED_USER"` 默认为空，仅影响覆盖人群数
}

// 查询抖音类目下的推荐达人
type ToolsAwemeCategoryTopAuthorGetRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	CategoryId   int      `json:"category_id,omitempty"`   // 类目id，一级，二级，三级类目id均可
	Behaviors    []string `json:"behaviors,omitempty"`     // 抖音用户行为类型, 详见[【附录-抖音用户行为类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"FOLLOWED_USER"`,`"COMMENTED_USER"`,`"LIKED_USER"`,`"SHARED_USER"` 默认为空,仅影响覆盖人群数
}

// 查询抖音号id对应的达人信息
type ToolsAwemeAuthorInfoGetRequest struct {
	AdvertiserId int64    `json:"advertiser_id,omitempty"` // 广告主ID
	LabelIds     []int    `json:"label_ids,omitempty"`     // 抖音号id列表，取值大小：1～50；label_id即计划中设置的抖音达人账号的id
	Behaviors    []string `json:"behaviors,omitempty"`     // 抖音用户行为类型, 详见[【附录-抖音用户行为类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"FOLLOWED_USER"`,`"COMMENTED_USER"`,`"LIKED_USER"`,`"SHARED_USER"`
}

