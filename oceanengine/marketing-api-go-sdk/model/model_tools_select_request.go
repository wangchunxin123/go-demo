/* #
# Author: wangchunxin
# Created Time: 2020/12/3 6:18 下午
# File Name:
# Description:
# */
package model

type ToolsAppSearchRequest struct {
	AdvertiserId int64  `json:"advertiser_id"`
	SearchBy     string `json:"search_by,omitempty"`
	AppName      string `json:"app_name,omitempty"`
	AppId        int    `json:"app_id,omitempty"`
}

// 查询受众预估结果
type ToolsEstimateAudienceRequest struct {
	AdvertiserId    int64  `json:"advertiser_id"`              // 广告主id
	RetargetingType string `json:"retargeting_type,omitempty"` // 定向人群包类型,详见【附录-定向人群包类型】,即将下线允许值: `"RETARGETING_INCLUDE"`,
	// `"RETARGETING_EXCLUDE"`,`"RETARGETING_NONE"`
	RetargetingTags        []int  `json:"retargeting_tags,omitempty"`         // 当选择使用人群包定向时填写，内容为人群包id，即将下线
	RetargetingTagsInclude []int  `json:"retargeting_tags_include,omitempty"` // 定向人群包列表，内容为人群包id，内容为人群包id
	RetargetingTagsExclude []int  `json:"retargeting_tags_exclude,omitempty"` // 排除人群包列表，内容为人群包id，内容为人群包id
	Gender                 string `json:"gender,omitempty"`                   // 受众性别,
	// 详见[【附录-受众性别】](https://ad.oceanengine.com/openapi/doc/index.
	// html?id=528)允许值:`"GENDER_FEMALE"`,`"GENDER_MALE"`,`"GENDER_UNLIMITED"`
	Age []string `json:"age,omitempty"` // 受众年龄区间,
	// 详见[【附录-受众年龄区间】](https://ad.oceanengine.com/openapi/doc/index.
	// html?id=528)允许值:`"AGE_BELOW_18"`,`"AGE_BETWEEN_18_23"`,`"AGE_BETWEEN_24_30"`,`"AGE_BETWEEN_31_40"`,`"AGE_BETWEEN_41_49"`,`"AGE_ABOVE_50"`
	AndroidOsv string `json:"android_osv,omitempty"` // 受众最低android版本(当推广应用下载Android时选填,其余情况不填),
	// 如果填写，对应的platform参数也要传。详见[【附录-受众android版本】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"0.0"`,`"2.0"`,`"2.1"`,`"2.2"`,`"2.3"`,`"3.0"`,`"3.1"`,`"3.2"`,`"4.0"`,`"4.1"`,`"4.2"`,`"4.3"`,`"4.4"`,`"4.5"`,`"5.0"`
	IosOsv string `json:"ios_osv,omitempty"` // 受众最低ios版本(当推广应用下载iOS时选填,其余情况不填),
	// 如果填写，对应的platform参数也要传。详见[【附录-受众ios版本】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"0.0"`,`"4.0"`,`"4.1"`,`"4.2"`,`"4.3"`,`"5.0"`,`"5.1"`,`"6.0"`,`"7.0"`,`"7.1"`,`"8.0"`,`"8.1"`,`"8.2"`,`"9.0"`
	Carrier []string `json:"carrier,omitempty"` // 受众运营商,
	// 详见[【附录-受众运营商类型】](https://ad.oceanengine.com/openapi/doc/index.
	// html?id=528)允许值:`"MOBILE"`,`"UNICOM"`,`"TELCOM"`
	Ac []string `json:"ac,omitempty"` // 受众网络类型,
	// 详见[【附录-受众网络类型】](https://ad.oceanengine.com/openapi/doc/index.
	// html?id=528)允许值:`"WIFI"`,`"2G"`,`"3G"`,`"4G"`
	DeviceBrand []string `json:"device_brand,omitempty"` // 受众手机品牌,
	// 详见[【附录-手机品牌】](https://ad.oceanengine.com/openapi/doc/index.
	// html?id=528)允许值:`"APPLE"`,`"HUAWEI"`,`"XIAOMI"`,`"SAMSUNG"`,`"OPPO"`,`"VIVO"`,`"MEIZU"`,`"GIONEE"`,`"COOLPAD"`,`"LENOVO"`,`"LETV"`,`"ZTE"`,`"CHINAMOBILE"`,`"HTC"`,`"PEPPER"`,`"NUBIA"`,`"HISENSE"`,`"QIKU"`,`"TCL"`,`"SONY"`,`"SMARTISAN"`,`"360"`,`"ONEPLUS"`,`"LG"`,`"MOTO"`,`"NOKIA"`,`"GOOGLE"`
	ArticleCategory []string `json:"article_category,omitempty"` // 受众文章分类,
	// 详见[【附录-受众文章分类】](https://ad.oceanengine.com/openapi/doc/index.
	// html?id=528)允许值:`"ENTERTAINMENT"`,`"SOCIETY"`,`"CARS"`,`"INTERNATIONAL"`,`"HISTORY"`,`"SPORTS"`,`"HEALTH"`,`"MILITARY"`,`"EMOTION"`,`"FASHION"`,`"PARENTING"`,`"FINANCE"`,`"AMUSEMENT"`,`"DIGITAL"`,`"EXPLORE"`,`"TRAVEL"`,`"CONSTELLATION"`,`"STORIES"`,`"TECHNOLOGY"`,`"GOURMET"`,`"CULTURE"`,`"HOME"`,`"MOVIE"`,`"PETS"`,`"GAMES"`,`"WEIGHT_LOSING"`,`"FREAK"`,`"EDUCATION"`,`"ESTATE"`,`"SCIENCE"`,`"PHOTOGRAPHY"`,`"REGIMEN"`,`"ESSAY"`,`"COLLECTION"`,`"ANIMATION"`,`"COMICS"`,`"TIPS"`,`"DESIGN"`,`"LOCAL"`,`"LAWS"`,`"GOVERNMENT"`,`"BUSINESS"`,`"WORKPLACE"`,`"RUMOR_CRACKER"`,`"GRADUATES"`
	ActivateType []string `json:"activate_type,omitempty"` // 用户首次激活时间,
	// 详见[【附录-用户首次激活时间】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"WITH_IN_A_MONTH"`,`"ONE_MONTH_2_THREE_MONTH"`,`"THREE_MONTH_EAILIER"`
	Platform []string `json:"platform,omitempty"` // 受众平台(当推广目的landing_type=APP时,不填,且为保证投放效果,平台类型定向PC与移动端互斥),
	// 详见[【附录-受众平台类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"ANDROID"`,`"IOS"`,`"PC"`
	City []int `json:"city,omitempty"` // 地域定向城市或者区县列表(当传递省份ID时,旗下市县ID可省略不传),
	// 详见【[附件-city.json】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)，district传CITY或COUNTY时必填】
	District []string `json:"district,omitempty"` // 地域类型，详见[【附录-地域类型】](https://ad.oceanengine.com/openapi/doc/index.
	// html?id=528)。允许值:`"CITY"`,`"COUNTY"`,`"BUSINESS_DISTRICT"`
	BusinessIds  []int  `json:"business_ids,omitempty"`  // 商圈ID数组，district传BUSINESS_DISTRICT时必填
	LocationType string `json:"location_type,omitempty"` // 受众位置类型，详见[【附录-受众位置类型】](https://ad.oceanengine.
	// com/openapi/doc/index.
	// html?id=528)
	AdTag []int `json:"ad_tag,omitempty"` // 兴趣分类,
	// 详见[【附件-ad_tag.json】](https://ad.oceanengine.com/openapi/doc/index. html?id=528)
	InterestTags      []int  `json:"interest_tags,omitempty"`       // 兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id
	AppBehaviorTarget string `json:"app_behavior_target,omitempty"` // APP行为定向,
	// 详见[【附录-APP行为定向类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"CATEGORY"`,`"APP"`,`"NONE"`
	AppCategory []int `json:"app_category,omitempty"` // APP行为定向,分类集合,
	// 详见[【附件-app_category.json】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	AppIds                 []int  `json:"app_ids,omitempty"`                  // APP行为定向,APP集合
	SuperiorPopularityType string `json:"superior_popularity_type,omitempty"` // 精选流量包，详见[【附录-精选流量包】](https://ad.oceanengine.
	// com/openapi/doc/index.
	// html?id=528)
	FlowPackage          []int       `json:"flow_package,omitempty"`           // 定向流量包ID数组
	ExcludeFlowPackage   []int       `json:"exclude_flow_package,omitempty"`   // 排除流量包ID数组
	IncludeCustomActions interface{} `json:"include_custom_actions,omitempty"` // 包含人群包((DPA推广目的特有,格式举例[{"days": 7, "code": 1001},], day可选范围:1,
	// 7,
	// 14,
	// 28, code候选范围由查询行为人群库接口得到)
}

// 建议日预算及预期成本
type ToolsBidSuggestRequest struct {
	AdvertiserId           int64       `json:"advertiser_id"`            // 广告主ID
	Pricing                string      `json:"pricing"`                  // 出价类型，查看[【附录-出价类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	CampaignId             int64       `json:"campaign_id"`              // 广告组ID
	AdId                   int64       `json:"ad_id"`                    // 广告ID，修改广告时需要传
	BidMode                string      `json:"bid_mode"`                 // 出价方式，自动&amp;手动允许值：`"SUGGEST"`、`"AUTO_BID"`- 手动获取到的是建议出价- 自动获取到的是建议日预算和预期成本
	BudgetMode             string      `json:"budget_mode"`              // 广告预算类型(创建后不可修改), 详见[【附录-预算类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"BUDGET_MODE_DAY"`,`"BUDGET_MODE_TOTAL"`
	Budget                 float64     `json:"budget"`                   // 广告预算( `出价方式为CPC、CPM、CPV时，不少于100元`；`出价方式为OCPM、OCPC时，不少于300元`。单次预算修改幅度不小于100元,日修改预算不超过20次)取值范围: `"≥ 0"`
	ScheduleType           string      `json:"schedule_type"`            // 广告投放时间类型, 详见[【附录-广告投放时间类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"SCHEDULE_FROM_NOW"`,`"SCHEDULE_START_END"`
	FlowControlMode        string      `json:"flow_control_mode"`        // 广告投放速度类型, 详见[【附录-广告投放速度类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"FLOW_CONTROL_MODE_FAST"`,`"FLOW_CONTROL_MODE_SMOOTH"`
	ConvertId              int64       `json:"convert_id"`               // 转化id，可通过【工具模块-OCPC广告创建转化查询】查询可用id
	RetargetingType        string      `json:"retargeting_type"`         // 定向人群包类型,详见[【附录-定向人群包类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528),即将下线允许值: `"RETARGETING_INCLUDE"`, `"RETARGETING_EXCLUDE"`,`"RETARGETING_NONE"`
	RetargetingTags        []int       `json:"retargeting_tags"`         // 当选择使用人群包定向时填写，内容为人群包id，即将下线
	RetargetingTagsInclude []int       `json:"retargeting_tags_include"` // 定向人群包列表，内容为人群包id，内容为人群包id
	RetargetingTagsExclude []int       `json:"retargeting_tags_exclude"` // 排除人群包列表，内容为人群包id，内容为人群包id
	Gender                 string      `json:"gender"`                   // 受众性别, 详见[【附录-受众性别】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"GENDER_FEMALE"`,`"GENDER_MALE"`,`"GENDER_UNLIMITED"`
	Age                    []string    `json:"age"`                      // 受众年龄区间, 详见[【附录-受众年龄区间】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"AGE_BELOW_18"`,`"AGE_BETWEEN_18_23"`,`"AGE_BETWEEN_24_30"`,`"AGE_BETWEEN_31_40"`,`"AGE_BETWEEN_41_49"`,`"AGE_ABOVE_50"`
	AndroidOsv             string      `json:"android_osv"`              // 受众最低android版本(当推广应用下载Android时选填,其余情况不填), 如果填写，对应的platform参数也要传。详见[【附录-受众android版本】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"0.0"`,`"2.0"`,`"2.1"`,`"2.2"`,`"2.3"`,`"3.0"`,`"3.1"`,`"3.2"`,`"4.0"`,`"4.1"`,`"4.2"`,`"4.3"`,`"4.4"`,`"4.5"`,`"5.0"`
	IosOsv                 string      `json:"ios_osv"`                  // 受众最低ios版本(当推广应用下载iOS时选填,其余情况不填),  如果填写，对应的platform参数也要传。详见[【附录-受众ios版本】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"0.0"`,`"4.0"`,`"4.1"`,`"4.2"`,`"4.3"`,`"5.0"`,`"5.1"`,`"6.0"`,`"7.0"`,`"7.1"`,`"8.0"`,`"8.1"`,`"8.2"`,`"9.0"`
	Carrier                []string    `json:"carrier"`                  // 受众运营商, 详见[【附录-受众运营商类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"MOBILE"`,`"UNICOM"`,`"TELCOM"`
	Ac                     []string    `json:"ac"`                       // 受众网络类型, 详见[【附录-受众网络类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"WIFI"`,`"2G"`,`"3G"`,`"4G"`
	DeviceBrand            []string    `json:"device_brand"`             // 受众手机品牌, 详见[【附录-手机品牌】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"APPLE"`,`"HUAWEI"`,`"XIAOMI"`,`"SAMSUNG"`,`"OPPO"`,`"VIVO"`,`"MEIZU"`,`"GIONEE"`,`"COOLPAD"`,`"LENOVO"`,`"LETV"`,`"ZTE"`,`"CHINAMOBILE"`,`"HTC"`,`"PEPPER"`,`"NUBIA"`,`"HISENSE"`,`"QIKU"`,`"TCL"`,`"SONY"`,`"SMARTISAN"`,`"360"`,`"ONEPLUS"`,`"LG"`,`"MOTO"`,`"NOKIA"`,`"GOOGLE"`
	ArticleCategory        []string    `json:"article_category"`         // 受众文章分类, 详见[【附录-受众文章分类】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"ENTERTAINMENT"`,`"SOCIETY"`,`"CARS"`,`"INTERNATIONAL"`,`"HISTORY"`,`"SPORTS"`,`"HEALTH"`,`"MILITARY"`,`"EMOTION"`,`"FASHION"`,`"PARENTING"`,`"FINANCE"`,`"AMUSEMENT"`,`"DIGITAL"`,`"EXPLORE"`,`"TRAVEL"`,`"CONSTELLATION"`,`"STORIES"`,`"TECHNOLOGY"`,`"GOURMET"`,`"CULTURE"`,`"HOME"`,`"MOVIE"`,`"PETS"`,`"GAMES"`,`"WEIGHT_LOSING"`,`"FREAK"`,`"EDUCATION"`,`"ESTATE"`,`"SCIENCE"`,`"PHOTOGRAPHY"`,`"REGIMEN"`,`"ESSAY"`,`"COLLECTION"`,`"ANIMATION"`,`"COMICS"`,`"TIPS"`,`"DESIGN"`,`"LOCAL"`,`"LAWS"`,`"GOVERNMENT"`,`"BUSINESS"`,`"WORKPLACE"`,`"RUMOR_CRACKER"`,`"GRADUATES"`
	ActivateType           []string    `json:"activate_type"`            // 用户首次激活时间, 详见[【附录-用户首次激活时间】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"WITH_IN_A_MONTH"`,`"ONE_MONTH_2_THREE_MONTH"`,`"THREE_MONTH_EAILIER"`
	Platform               []string    `json:"platform"`                 // 受众平台(当推广目的landing_type=APP时,不填,且为保证投放效果,平台类型定向PC与移动端互斥), 详见[【附录-受众平台类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"ANDROID"`,`"IOS"`,`"PC"`
	City                   []int       `json:"city"`                     // 地域定向城市或者区县列表(当传递省份ID时,旗下市县ID可省略不传), 详见[【附件-city.json】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)，district传CITY或COUNTY时必填】
	District               string      `json:"district"`                 // 地域类型，详见[【附录-地域类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)。允许值:`"CITY"`,`"COUNTY"`,`"BUSINESS_DISTRICT"`
	BusinessIds            []int       `json:"business_ids"`             // 商圈ID数组，district传BUSINESS_DISTRICT时必填
	LocationType           string      `json:"location_type"`            // 受众位置类型，详见[【附录-受众位置类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	AdTag                  []int       `json:"ad_tag"`                   // 兴趣分类, 详见【附件-ad_tag.json】
	InterestTags           []int       `json:"interest_tags"`            // 兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id
	AppBehaviorTarget      string      `json:"app_behavior_target"`      // APP行为定向, 详见[【附录-APP行为定向类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值:`"CATEGORY"`,`"APP"`,`"NONE"`
	AppCategory            []int       `json:"app_category"`             // APP行为定向,分类集合, 详见[【附件-app_category.json】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	AppIds                 []int       `json:"app_ids"`                  // APP行为定向,APP集合
	SuperiorPopularityType string      `json:"superior_popularity_type"` // 精选流量包，详见[【附录-精选流量包】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	FlowPackage            []int       `json:"flow_package"`             // 定向流量包ID数组
	ExcludeFlowPackage     []int       `json:"exclude_flow_package"`     // 排除流量包ID数组
	IncludeCustomActions   interface{} `json:"include_custom_actions"`   // 包含人群包((DPA推广目的特有,格式举例[{"days": 7, "code": 1001},], day可选范围:1, 7, 14, 28, code候选范围由查询行为人群库接口得到)
	ExcludeCustomActions   interface{} `json:"exclude_custom_actions"`   // 排除人群包((DPA推广目的特有,格式举例[{"days": 7, "code": 1002},], day可选范围:1, 7, 14, 28, code候选范围由查询行为人群库接口得到)
}

// 查询广告质量度
type ToolsAdQualityGetRequest struct {
	AdvertiserId int64   `json:"advertiser_id"` // 广告主ID
	AdIds        []int64 `json:"ad_ids"`        // 广告id列表
}

// 获取行业列表
type ToolsIndustryGetRequest struct {
	Level int    `json:"level,omitempty"` // 只获取某级别数据，1:第一级,2:第二级,3:第三级，默认都返回
	Type  string `json:"type,omitempty"`  // 可选值：`"ADVERTISER"`，`"AGENT"`，`"ADVERTISER"`为原有广告3.0行业, `"AGENT"`为代理商行业获取，代理商行业level都为1
}

//获取地域列表
type ToolsRegionGetRequest struct {
	RegionType  string `json:"region_type,omitempty"`  // 地域类型，目前只支持：BUSINESS_DISTRICT(商圈)允许值:`"BUSINESS_DISTRICT"`
	RegionLevel string `json:"region_level,omitempty"` // 只获取某层级数据，详见【附录-地域层级】
}

//获取鲁班商品列表
type ToolsAdvertiserGoodsListRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	Page         int   `json:"page,omitempty"`          // 门页数默认值:  `1`
	PageSize     int   `json:"page_size,omitempty"`     // 页面大小默认值: `10`
	Filtering    struct {
		ProductId   string `json:"product_id,omitempty"`   // 商品编号
		ProductName string `json:"product_name,omitempty"` // 商品名称
		StartTime   string `json:"start_time,omitempty"`   // 查询起始时间
		EndTime     string `json:"end_time,omitempty"`     // 查询结束时间
		Status      string `json:"status,omitempty"`       // 上下架状态,详见附录[【商品上下架状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
		CheckStatus string `json:"check_status,omitempty"` // 审核状态,详见附录[【商品审核状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	} `json:"filtering,omitempty"` // 过滤条件，若此字段不传，或传空则视为无限制条件
}

//获取广告预览二维码
type ToolsAdPreviewQrcodeGetRequest struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 广告主ID
	IdType       string `json:"id_type,omitempty"`       // 查询条件，可选值： `"ID_TYPE_AD"`,`"ID_TYPE_CREATIVE"``"ID_TYPE_AD"`:按广告计划的ID搜索`"ID_TYPE_CREATIVE"`:按广告创意的ID搜索
	Id           int    `json:"id,omitempty"`            // 广告计划或者广告创意的ID，根据id_type进行相应id的填写
}

//获取绑定的抖音号
type ToolsIesAccountSearchRequest struct {
	AdvertiserId int64 `json:"advertiser_id"` // 广告主ID
}

//查询极速下载列表
type ToolsAdvertiserPackageGetRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	Page         int   `json:"page,omitempty"`          // 页数默认值: `1`
	PageSize     int   `json:"page_size,omitempty"`     // 页面大小默认值: `10`
}

//日志查询
type ToolsLogSearchRequest struct {
	AdvertiserId int64   `json:"advertiser_id,omitempty"` // 广告主ID
	ObjectId     []int64 `json:"object_id,omitempty"`     // 操作对象ID, 1 &lt;= len &lt;= 20 , 可以为campaign_id、ad_id、creative_id， 各种id可以随意组合
	StartTime    string  `json:"start_time,omitempty"`    // 日志查询开始时间，格式 "2019-07-24 21:46:57"
	EndTime      string  `json:"end_time,omitempty"`      // 日志查询开始时间，格式 "2019-07-24 21:46:57"
	Page         int     `json:"page,omitempty"`          // 页码默认值: `1`
	PageSize     int     `json:"page_size,omitempty"`     // 页面大小默认值: `10`允许值:`1~20`
}

//查询程序化创意包
type ToolsProceduralPackageGetRequest struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 广告主ID
	AdId         int64  `json:"ad_id,omitempty"`         // 计划ID和审核相关，不传ad_id返回的package信息里面status=1, 如果传入ad_id，当审核拒绝时就返回status=2
	Keyword      string `json:"keyword,omitempty"`       // 搜索关键词，根据关键词对创意包名进行模糊查询
	Page         int    `json:"page,omitempty"`          // 页数默认值: 1
	PageSize     int    `json:"page_size,omitempty"`     // 页面大小默认值: 10
}

//行动号召字段内容获取
type ToolsActionTextGetRequest struct {
	AdvertiserId         int64  `json:"advertiser_id,omitempty"`          // 广告主ID
	LandingType          string `json:"landing_type,omitempty"`           // 推广类型允许值：APP，SHOP，LINK
	AdvancedCreativeType string `json:"advanced_creative_type,omitempty"` // 附加创意类型，详见枚举
	Industry             int    `json:"industry,omitempty"`               // 广告主行业id，可以从[获取行业接口](https://ad.oceanengine.com/openapi/doc/index.html?id=370)进行获取
}

//查询广告计划学习期状态
type ToolsAdStatExtraInfoGetRequest struct {
	AdvertiserId int64   `json:"advertiser_id,omitempty"` // 广告主ID
	AdIds        []int64 `json:"ad_ids,omitempty"`        // 广告计划 id，最多传100个广告计划id。
}

//查询推广卡片推荐内容
type ToolsPromotionCardRecommendGetRequest struct {
	AdvertiserId  int64  `json:"advertiser_id,omitempty"`  // 广告主id
	AdId          int64  `json:"ad_id,omitempty"`          // 广告计划ID。请传入属于当前广告主的计划ID，否则会导致获得的推荐内容不准确或不可用
	RecommendType string `json:"recommend_type,omitempty"` // 文案推荐类型，详见[【附录-文案推荐类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528#custom-anchor-文案推荐类型)允许值: `PRODUCT_DESCRIPTION`, `SELLING_POINTS`,`CALL_TO_ACTION`
	TitleList     struct {
		Title string `json:"title,omitempty"` // 创意标题，文案推荐类型为`PRODUCT_DESCRIPTION`、`SELLING_POINTS`时必填
	} `json:"title_list,omitempty"`                                       // 创意标题列表，文案推荐类型为`PRODUCT_DESCRIPTION`、`SELLING_POINTS`时必填
	AdvancedCreativeType string `json:"advanced_creative_type,omitempty"` // 附加创意类型，文案推荐类型为`CALL_TO_ACTION`时必填，详见[【附录-附加创意类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528#custom-anchor-附加创意类型)允许值:`ATTACHED_CREATIVE_NONE`,`ATTACHED_CREATIVE_FORM`,`ATTACHED_CREATIVE_CONSULTANT`,`ATTACHED_CREATIVE_COUPON`,`ATTACHED_CREATIVE_SMART_PHONE`,`ATTACHED_CREATIVE_GAME_FORM`,`ATTACHED_CREATIVE_GAME_SUBSCRIBE`,`ATTACHED_CREATIVE_GAME_PACKAGE`
	DownloadType         string `json:"download_type,omitempty"`          // 应用下载类型允许值:`DOWNLOAD_URL` 下载链接，`EXTERNAL_URL` 落地页链接
}
