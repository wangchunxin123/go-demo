/* #
# Author: wangchunxin
# Created Time: 2020/12/8 2:32 下午
# File Name:
# Description:
# */
package model

// 获取定向包
type AudiencePackageGetRequest struct {
	AdvertiserId  int64  `json:"advertiser_id,omitempty"`  // 广告主ID
	Filtering     interface{} `json:"filtering,omitempty"`      // 过滤字段
	LandingType   string `json:"landing_type,omitempty"`   // 定向包类型[【附录-定向包推广类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528#custom-anchor-定向包推广类型)
	DeliveryRange string `json:"delivery_range,omitempty"` // 广告投放范围[【附录-广告投放范围】](https://ad.oceanengine.com/openapi/doc/index.html?id=528#custom-anchor-广告投放范围)
	Page          int    `json:"page,omitempty"`           // 页码默认值:`1`
	PageSize      int    `json:"page_size,omitempty"`      // 页面数据量默认值: `10`，取值范围1～100
}

// 创建定向包
type AudiencePackageCreateRequest struct {
	AdvertiserId           int64    `json:"advertiser_id,omitempty"`            // 广告主ID
	Name                   string   `json:"name,omitempty"`                     // 定向包名称
	Description            string   `json:"description,omitempty"`              // 定向包描述
	LandingType            string   `json:"landing_type,omitempty"`             // 定向包类型[【附录-定向包推广类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528#custom-anchor-定向包推广类型)
	DeliveryRange          string   `json:"delivery_range,omitempty"`           // 广告投放范围[【附录-广告投放范围】](https://ad.oceanengine.com/openapi/doc/index.html?id=528#custom-anchor-广告投放范围)
	RetargetingTags        []int    `json:"retargeting_tags,omitempty"`         // 定向人群包列表，内容为人群包id
	RetargetingTagsExclude []int    `json:"retargeting_tags_exclude,omitempty"` // 排除人群包列表，内容为人群包id
	Gender                 string   `json:"gender,omitempty"`                   // 受众性别, 详见[【附录-受众性别](https://ad.oceanengine.com/openapi/doc/index.html?id=528)】允许值: `GENDER_FEMALE`,`GENDER_MALE` , `NONE`
	Age                    []string `json:"age,omitempty"`                      // 受众年龄区间, 详见[【附录-受众年龄区间】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `AGE_BETWEEN_18_23`,`AGE_BETWEEN_24_30`,`AGE_BETWEEN_31_40`, `AGE_BETWEEN_41_49`,`AGE_ABOVE_50`
	AndroidOsv             string   `json:"android_osv,omitempty"`              // 受众最低android版本(当推广应用下载Android时选填,其余情况不填), 详见[【附录-受众android版本】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: "0.0", "2.0", "2.1", "2.2", "2.3", "3.0", "3.1", "3.2", "4.0","4.1","4.2", "4.3", "4.4", "4.5", "5.0", "NONE"
	IosOsv                 string   `json:"ios_osv,omitempty"`                  // 受众最低ios版本(当推广应用下载iOS时选填,其余情况不填), 详见[【附录-受众ios版本】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: "0.0", "4.0", "4.1", "4.2", "4.3", "5.0", "5.1", "6.0", "7.0","7.1","8.0", "8.1", "8.2", "9.0", "NONE"
	Carrier                []string `json:"carrier,omitempty"`                  // 受众运营商, 详见[【附录-受众运营商类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: "MOBILE", "UNICOM", "TELCOM"
	Ac                     []string `json:"ac,omitempty"`                       // 受众网络类型, 详见[【附录-受众网络类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	DeviceBrand            []string `json:"device_brand,omitempty"`             // 受众手机品牌, 详见[【附录-手机品牌】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	ArticleCategory        []string `json:"article_category,omitempty"`         // 受众文章分类, 详见[【附录-受众文章分类】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	ActivateType           []string `json:"activate_type,omitempty"`            // 用户首次激活时间, 详见[【附录-用户首次激活时间】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	Platform               []string `json:"platform,omitempty"`                 // 受众平台(当推广目的landing_type=APP时,不填,且为保证投放效果,平台类型定向PC与移动端互斥)，android_osv或ios_osv非空时，必填。详见[【附录-受众平台类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	AutoExtendEnabled      int      `json:"auto_extend_enabled,omitempty"`      // 是否启动智能放量允许值：0表示关闭，1表示开启
	AutoExtendTargets      []string `json:"auto_extend_targets,omitempty"`      // 智能放量定向允许值：`REGION`（地域）,`GENDER`（性别）,`AGE`（年龄）,`CUSTOM_AUDIENCE`（自定义人群）,`INTEREST_ACTION`（行为兴趣）
	LaunchPrice            []int    `json:"launch_price,omitempty"`             // 手机价格定向,传入价格区间，最高传入11000（表示1w以上）传值示例 "launch_price": [2000, 11000]，表示2000元以上
	InterestActionMode     string   `json:"interest_action_mode,omitempty"`     // 行为兴趣选择允许值：`UNLIMITED`（不限）,`CUSTOM`（自定义）,`RECOMMEND`（系统推荐）
	ActionScene            []string `json:"action_scene,omitempty"`             // 行为场景，详见【附录-行为场景】，当interest_action_mode传CUSTOM时有效允许值：`E-COMMERCE`,`NEWS`, `APP`
	ActionDays             int      `json:"action_days,omitempty"`              // 行为天数,当interest_action_mode传CUSTOM时有效允许值：7, 15, 30, 60, 90, 180, 365
	ActionCategories       []int    `json:"action_categories,omitempty"`        // 行为类目,当interest_action_mode传CUSTOM时有效
	ActionWords            []int    `json:"action_words,omitempty"`             // 行为关键词,当interest_action_mode传CUSTOM时有效
	InterestCategories     []int    `json:"interest_categories,omitempty"`      // 兴趣分类,如果传空数组 [] 表示不限，如果只传[0]表示系统推荐,如果按兴趣类型传表示自定义，通过“查询工具”-兴趣类目/兴趣关键词查询。当interest_action_mode传CUSTOM时有效
	InterestWords          []int    `json:"interest_words,omitempty"`           // 兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。当interest_action_mode传CUSTOM时有效
	City                   []int    `json:"city,omitempty"`                     // 地域定向城市或者区县列表(当传递省份ID时,旗下市县ID可省略不传), 详见[【附件-city.json】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	BusinessIds            []int    `json:"business_ids,omitempty"`             // 商圈ID数组
	District               interface{}     `json:"district,omitempty"`                 // 地域类型，前者为省市，后者为区县。当city有数据时，必填。允许值: `"CITY"`, `"COUNTY"`,
	// `"BUSINESS_DISTRICT"`,`"NONE"`](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	LocationType           string   `json:"location_type,omitempty"`            // 受众位置类型,当city和district有值时，该字段必填，详见[【附录-受众位置类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	Geolocation            struct {
		Province     string `json:"province,omitempty"`      // 省份名
		City         string `json:"city,omitempty"`          // 城市名
		Name         string `json:"name,omitempty"`          // 地点名字
		Long         float64  `json:"long,omitempty"`          // 经度lat
		Lat          float64  `json:"lat,omitempty"`           // 纬度
		Street       string `json:"street,omitempty"`        // 街道名
		District     string `json:"district,omitempty"`      // 区域名
		Radius       int    `json:"radius,omitempty"`        // 半径范围
		StreetNumber string `json:"street_number,omitempty"` // 街道号
	} `json:"geolocation,omitempty"`                                                   // 地图位置，district为BUSINESS_DISTRICT才有效
	AwemeFansNumbers          []int    `json:"aweme_fans_numbers,omitempty"`           // （抖音推广特有）账号粉丝相似人群（添加抖音账号，会将广告投放给对应账号的相似人群粉丝）
	FilterAwemeAbnormalActive int      `json:"filter_aweme_abnormal_active,omitempty"` // （抖音推广特有）过滤高活跃用户允许值：`0`:表示不过滤，`1`:表示过滤
	FilterAwemeFansCount      int      `json:"filter_aweme_fans_count,omitempty"`      // （抖音推广特有）过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户
	FilterOwnAwemeFans        int      `json:"filter_own_aweme_fans,omitempty"`        // （抖音推广特有）过滤自己的粉丝，允许值：0表示不过滤，1表示过滤
	AwemeFanAccounts          []int    `json:"aweme_fan_accounts,omitempty"`           // 抖音号id，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanCategories        []int    `json:"aweme_fan_categories,omitempty"`         // 抖音类目id，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanBehaviors         []string `json:"aweme_fan_behaviors,omitempty"`          // 抖音用户行为类型, 详见[【附录-抖音用户行为类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"FOLLOWED_USER"`,`"COMMENTED_USER"`,`"LIKED_USER"`,`"SHARED_USER"`（抖音达人定向）
	HideIfExists              int      `json:"hide_if_exists,omitempty"`               // 已安装用户，0表示不限，1表示过滤，2表示定向；过滤表示投放时不给安装客户展示广告，支持应用推广；定向表示投放时给安装客户展示广告；投放时优先获取直达链接，无直达链接时使用应用包名进行投放；如果无直达链接或应用包名，定向安装选项实际不生效；定向仅对Android链接生效。默认值:`0`允许值: `0`, `1`,`2`
}

// 更新定向包
type AudiencePackageUpdateRequest struct {
	AdvertiserId           int64    `json:"advertiser_id,omitempty"`            // 广告主ID
	Name                   string   `json:"name,omitempty"`                     // 定向包名称
	Description            string   `json:"description,omitempty"`              // 定向包描述
	AudiencePackageId      string   `json:"audience_package_id,omitempty"`      // 修改的定向包ID
	RetargetingTags        []int    `json:"retargeting_tags,omitempty"`         // 定向人群包列表，内容为人群包id
	RetargetingTagsExclude []int    `json:"retargeting_tags_exclude,omitempty"` // 排除人群包列表，内容为人群包id
	Gender                 string   `json:"gender,omitempty"`                   // 受众性别, 详见[【附录-受众性别】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `GENDER_FEMALE`,`GENDER_MALE` , `NONE`
	Age                    []string `json:"age,omitempty"`                      // 受众年龄区间, 详见[【附录-受众年龄区间】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `AGE_BETWEEN_18_23`,`AGE_BETWEEN_24_30`,`AGE_BETWEEN_31_40`, `AGE_BETWEEN_41_49`,`AGE_ABOVE_50`
	AndroidOsv             string   `json:"android_osv,omitempty"`              // 受众最低android版本(当推广应用下载Android时选填,其余情况不填), 详见[【附录-受众android版本】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: "0.0", "2.0", "2.1", "2.2", "2.3", "3.0", "3.1", "3.2", "4.0","4.1","4.2", "4.3", "4.4", "4.5", "5.0", "NONE"
	IosOsv                 string   `json:"ios_osv,omitempty"`                  // 受众最低ios版本(当推广应用下载iOS时选填,其余情况不填), 详见[【附录-受众ios版本】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: "0.0", "4.0", "4.1", "4.2", "4.3", "5.0", "5.1", "6.0", "7.0","7.1","8.0", "8.1", "8.2", "9.0", "NONE"
	Carrier                []string `json:"carrier,omitempty"`                  // 受众运营商, 详见[【附录-受众运营商类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: "MOBILE", "UNICOM", "TELCOM"
	Ac                     []string `json:"ac,omitempty"`                       // 受众网络类型, 详见[【附录-受众网络类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	DeviceBrand            []string `json:"device_brand,omitempty"`             // 受众手机品牌, 详见[【附录-手机品牌】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	ArticleCategory        []string `json:"article_category,omitempty"`         // 受众文章分类, 详见[【附录-受众文章分类】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	ActivateType           []string `json:"activate_type,omitempty"`            // 用户首次激活时间, 详见[【附录-用户首次激活时间】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	Platform               []string `json:"platform,omitempty"`                 // 受众平台(当推广目的landing_type=APP时,不填,且为保证投放效果,平台类型定向PC与移动端互斥)，android_osv或ios_osv非空时，必填。详见[【附录-受众平台类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	AutoExtendEnabled      int      `json:"auto_extend_enabled,omitempty"`      // 是否启动智能放量允许值：0表示关闭，1表示开启
	AutoExtendTargets      []string `json:"auto_extend_targets,omitempty"`      // 智能放量定向允许值：`REGION`（地域）,`GENDER`（性别）,`AGE`（年龄）,`CUSTOM_AUDIENCE`（自定义人群）
	LaunchPrice            []int    `json:"launch_price,omitempty"`             // 手机价格定向,传入价格区间，最高传入11000（表示1w以上）传值示例 "launch_price": [2000, 11000]，表示2000元以上
	InterestActionMode     string   `json:"interest_action_mode,omitempty"`     // 行为兴趣选择允许值：`UNLIMITED`（不限）,`CUSTOM`（自定义）,`RECOMMEND`（系统推荐）
	ActionScene            []string `json:"action_scene,omitempty"`             // 行为场景，详见[【附录-行为场景】](https://ad.oceanengine.com/openapi/doc/index.html?id=528),当interest_action_mode传CUSTOM时有效允许值：`E-COMMERCE`,`NEWS`, `APP`
	ActionDays             int      `json:"action_days,omitempty"`              // 行为天数,当interest_action_mode传CUSTOM时有效允许值：7, 15, 30, 60, 90, 180, 365
	ActionCategories       []int    `json:"action_categories,omitempty"`        // 行为类目,当interest_action_mode传CUSTOM时有效
	ActionWords            []int    `json:"action_words,omitempty"`             // 行为关键词,当interest_action_mode传CUSTOM时有效
	InterestCategories     []int    `json:"interest_categories,omitempty"`      // 兴趣分类,如果传空数组 [] 表示不限，如果只传[0]表示系统推荐,如果按兴趣类型传表示自定义，通过“查询工具”-兴趣类目/兴趣关键词查询。当interest_action_mode传CUSTOM时有效
	InterestWords          []int    `json:"interest_words,omitempty"`           // 兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。当interest_action_mode传CUSTOM时有效
	City                   []int    `json:"city,omitempty"`                     // 地域定向城市或者区县列表(当传递省份ID时,旗下市县ID可省略不传), 详见[【附件-city.json】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	BusinessIds            []int    `json:"business_ids,omitempty"`             // 商圈ID数组
	District               []string `json:"district,omitempty"`                 // 地域类型，前者为省市，后者为区县。当city有数据时，必填。允许值: `"CITY"`, `"COUNTY"`, `"BUSINESS_DISTRICT"`, `"NONE"`](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	LocationType string `json:"location_type,omitempty"` // 受众位置类型,当city和district有值时，该字段必填，详见[【附录-受众位置类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	Geolocation  struct {
		Province     string  `json:"province,omitempty"`      // 省份名
		City         string  `json:"city,omitempty"`          // 城市名
		Name         string  `json:"name,omitempty"`          // 地点名字
		Long         float64 `json:"long,omitempty"`          // 经度lat
		Lat          float64 `json:"lat,omitempty"`           // 纬度
		Street       string  `json:"street,omitempty"`        // 街道名
		District     string  `json:"district,omitempty"`      // 区域名
		Radius       int     `json:"radius,omitempty"`        // 半径范围
		StreetNumber string  `json:"street_number,omitempty"` // 街道号
	} `json:"geolocation,omitempty"`                                                   // 地图位置，district为BUSINESS_DISTRICT才有效
	AwemeFansNumbers          []int    `json:"aweme_fans_numbers,omitempty"`           // （抖音推广特有）账号粉丝相似人群（添加抖音账号，会将广告投放给对应账号的相似人群粉丝）
	FilterAwemeAbnormalActive int      `json:"filter_aweme_abnormal_active,omitempty"` // （抖音推广特有）过滤高活跃用户允许值：`0`:表示不过滤，`1`:表示过滤
	FilterAwemeFansCount      int      `json:"filter_aweme_fans_count,omitempty"`      // （抖音推广特有）过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户
	FilterOwnAwemeFans        int      `json:"filter_own_aweme_fans,omitempty"`        // 过滤自己的粉丝，允许值：0表示不过滤，1表示过滤
	AwemeFanAccounts          []int    `json:"aweme_fan_accounts,omitempty"`           // 抖音号id，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanCategories        []int    `json:"aweme_fan_categories,omitempty"`         // 抖音类目id，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanBehaviors         []string `json:"aweme_fan_behaviors,omitempty"`          // 抖音用户行为类型, 详见[【附录-抖音用户行为类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"FOLLOWED_USER"`,`"COMMENTED_USER"`,`"LIKED_USER"`,`"SHARED_USER"`（抖音达人定向）
	HideIfExists              int      `json:"hide_if_exists,omitempty"`               // 已安装用户，0表示不限，1表示过滤，2表示定向；过滤表示投放时不给安装客户展示广告，支持应用推广；定向表示投放时给安装客户展示广告；投放时优先获取直达链接，无直达链接时使用应用包名进行投放；如果无直达链接或应用包名，定向安装选项实际不生效；定向仅对Android链接生效。允许值: `0`, `1`,`2`
}

// 删除定向包
type AudiencePackageDeleteRequest struct {
	AdvertiserId      int64 `json:"advertiser_id,omitempty"`       // 广告主ID
	AudiencePackageId int   `json:"audience_package_id,omitempty"` // 删除的定向包id
}

// 计划绑定定向包
type AudiencePackageAdBindRequest struct {
	AdvertiserId      int64 `json:"advertiser_id,omitempty"`       // 广告主ID
	AudiencePackageId int   `json:"audience_package_id,omitempty"` // 定向包id
	AdIds             []int `json:"ad_ids,omitempty"`              // 绑定的计划id列表，上限1000个
}

// 定向包解绑
type AudiencePackageAdUnbindRequest struct {
	AdvertiserId      int64 `json:"advertiser_id,omitempty"`       // 广告主ID
	AudiencePackageId int   `json:"audience_package_id,omitempty"` // 删除的定向包id
	AdIds             []int `json:"ad_ids,omitempty"`              // 绑定的计划id列表，上限1000个
}
