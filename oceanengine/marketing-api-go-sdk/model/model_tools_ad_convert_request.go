/* #
# Author: wangchunxin
# Created Time: 2020/12/8 2:31 下午
# File Name:
# Description:
# */
package model

// 创建转化ID
type ToolsAdConvertCreateRequest struct {
	AdvertiserId               int64    `json:"advertiser_id,omitempty"`                  // 广告主ID
	Name                       string   `json:"name,omitempty"`                           // 转化名称
	ConvertSourceType          string   `json:"convert_source_type,omitempty"`            // 转化来源，详见[【附录-转化工具来源】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)，默认为`"AD_CONVERT_SOURCE_TYPE_APP_DOWNLOAD"`
	DownloadUrl                string   `json:"download_url,omitempty"`                   // 下载链接
	AppType                    string   `json:"app_type,omitempty"`                       // 应用类型允许值: `"APP_ANDROID"`, `"APP_IOS"`，当传app_id时，app_type必填
	ActionTrackUrl             string   `json:"action_track_url,omitempty"`               // 点击监测链接
	DisplayTrackUrl            string   `json:"display_track_url,omitempty"`              // 展示监测链接
	VideoPlayEffectiveTrackUrl string   `json:"video_play_effective_track_url,omitempty"` // 视频有效播放监测链接
	VideoPlayDoneTrackUrl      string   `json:"video_play_done_track_url ,omitempty"`     // 视频播放完毕监测链接
	VideoPlayTrackUrl          string   `json:"video_play_track_url ,omitempty"`          // 视频播放监测链接
	PackageName                string   `json:"package_name,omitempty"`                   // 包名，当转化类型和规则为应用下载SDK/应用下载API时必填
	ConvertType                string   `json:"convert_type,omitempty"`                   // 转化类型，详见[【附录-转化类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	ConvertXpathValue          string   `json:"convert_xpath_value,omitempty"`            // XPath转化路径
	ConvertXpathUrl            string   `json:"convert_xpath_url,omitempty"`              // XPath转化页面地址
	ExternalUrl                string   `json:"external_url,omitempty"`                   // 落地页地址
	AppId                      string   `json:"app_id,omitempty"`                         // APP ID
	AppName                    string   `json:"app_name,omitempty"`                       // app name
	OpenUrl                    string   `json:"open_url,omitempty"`                       // 直达链接
	XpathIgnoreParams          []string `json:"xpath_ignore_params,omitempty"`            // 匹配规则字段(xpath下可传)允许值:`"UTM_ID"`,`"CID"`，`"ADID"`
	ForceActive                int      `json:"force_active,omitempty"`                   // 是否启用自动激活免联调，启用后无需前往投放平台进行联调激活，将会自动激活。枚举值：0（不启用），1（启用）；默认不启用注：该功能为白名单功能，需要联系管理员给应用添加免联调白名单后才可使用该字段。
	ConvertDataType            string   `json:"convert_data_type,omitempty"`              // 转化统计方式，针对每次付费广告，投放范围是站内和穿山甲，转化来源是应用下载SDK/API这两种方式，广告平台统计该转化目标是否发生的方式，默认“仅一次”（即，每个用户最多仅统计一次转化行为）。允许值：1. ONLY_ONE（仅一次）：对于每位转化的用户，仅统计其首次“目标事件”的转化行为，即每位用户最多仅记录一次转化。 2. EVERY_ONE（每一次）：对于每位转化的用户，统计其每次“目标事件”的发生次数，即每位用户可记录多次发生的转化；该统计方式下，创建广告计划时deep_bid_type须为BID_PER_ACTION。注：如果广告主ID不在白名单里面，且统计方式选择EVERY_ONE，请求会失败，报错信息“convertDataType not in whiteList”。
}

// 修改转化ID
type ToolsAdConvertUpdateRequest struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 广告主ID
	ConvertId    int    `json:"convert_id,omitempty"`    // 转化id
	DownloadUrl  string `json:"download_url,omitempty"`  // 下载链接
}

// 更新转化状态
type ToolsAdConvertUpdateStatusRequest struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 广告主ID
	ConvertId    int    `json:"convert_id,omitempty"`    // 转化id
	OptStatus    string `json:"opt_status,omitempty"`    // 操作状态，详见[【附录-操作状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值：`"AD_CONVERT_OPT_STATUS_ENABLE"`,`"AD_CONVERT_OPT_STATUS_DISABLE"`
}

// 【新】查询计划可用转化id
type ToolsAdConvertQueryRequest struct {
	AdvertiserId         int    `json:"advertiser_id ,omitempty"`          //  广告主id
	LandingType          string `json:"landing_type ,omitempty"`           //  广告组推广目的, 详见[【附录-枚举值-推广目的类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528) 可选值: `APP`,`ARTICLE`,`AWEME`,`DPA`,`GOODS`,`LINK`,`SHOP`
	PromotionContent     string `json:"promotion_content ,omitempty"`      //  投放内容，根据不同推广目的对应的不同的投放内容， 详见[【附录-枚举值】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)，可选值: `AWEME_HOME_PAGE`,`DOUYIN`,`DOWNLOAD_URL`,`EXTERNAL_URL`,`GOODS_LINK`,`LIVE_ROOM`,`MICRO_APP`,`NORMAL`,`QUICK_APP_URL`,`SHOP`,`THIRD_PARTY`
	DeliveryRange        string `json:"delivery_range ,omitempty"`         //  广告投放范围，详见[【附录-枚举值-广告投放范围】](https://ad.oceanengine.com/openapi/doc/index.html?id=528) 可选值: `DEFAULT`,`UNION`,`UNIVERSAL`
	ExternalUrl          string `json:"external_url ,omitempty"`           // 落地页链接
	AppType              string `json:"app_type ,omitempty"`               //  应用下载类型, 详见[【附录-应用下载类型-枚举值】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)，可选值: `APP_ANDROID`,`APP_IOS`
	PackageName          string `json:"package_name ,omitempty"`           //  android应用包名
	ItunesUrl            string `json:"itunes_url ,omitempty"`             //  iOS应用下载链接
	AppSchema            string `json:"app_schema ,omitempty"`             //  小程序app_schema
	AdvancedCreativeType string `json:"advanced_creative_type ,omitempty"` //  附加创意类型, 详见[【附录-枚举值-附加创意类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	MarketingScene       string `json:"marketing_scene ,omitempty"`        //  游戏预约场景，附加创意类型为游戏预约时选取, 详见[【附录-枚举值-游戏预约场景】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)，可选值: `GAME_PROMOTION`,`GAME_SUBSCRIBE`,`NORMAL`
}

// 转化ID列表
type ToolsAdvConvertSelectRequest struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 广告主ID
	ConvertIds   []int  `json:"convert_ids,omitempty"`   // 需要查询的convert_ids，如不填写默认返回所有的转化ID
	OptStatus    string `json:"opt_status,omitempty"`    // 操作状态
	Page         int    `json:"page,omitempty"`          // 页数默认值:`1`
	PageSize     int    `json:"page_size,omitempty"`     // 页面大小默认值:`10`
}

// 查询转化详细信息
type ToolsAdConvertReadRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	ConvertId    int   `json:"convert_id,omitempty"`    // 转化id，其中较小数值convert_id为预定义转化，具体枚举可查看[【附录-预定义转化类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
}

// 转化ID推送
type ToolsAdConvertPushRequest struct {
	AdvertiserId        int64 `json:"advertiser_id,omitempty"`         // 广告主ID
	ConvertId           int   `json:"convert_id,omitempty"`            // 转化id
	TargetAdvertiserIds []int `json:"target_advertiser_ids,omitempty"` // 推送的广告主ID列表
}

// 查询深度优化方式
type ToolsAdConvertDeepbidReadRequest struct {
	AdvertiserId       int    `json:"advertiser_id ,omitempty"`       // 广告主id
	CampaignId         int    `json:"campaign_id ,omitempty"`         // 广告组id
	DeepExternalAction string `json:"deep_external_action,omitempty"` // 深度转化目标，详见[【附录-深度转化类型】](https://openapi.bytedance.net/openapi/doc/index.html?id=395)允许值：`AD_CONVERT_TYPE_PAY`、`AD_CONVERT_TYPE_NEXT_DAY_OPEN`、`AD_CONVERT_TYPE_GAME_ADDICTION`、`AD_CONVERT_TYPE_PURCHASE_ROI`、`AD_CONVERT_TYPE_LT_ROI`
	DeliveryRange      string `json:"delivery_range ,omitempty"`      // 投放范围,[【附录-枚举值-广告投放范围】](https://ad.oceanengine.com/openapi/doc/index.html?id=528#custom-anchor-广告投放范围)
	ConvertId          int    `json:"convert_id,omitempty"`           // 转化id，`convert_id`和`external_action`二选一
	ExternalAction     string `json:"external_action,omitempty"`      // 转化类型，`convert_id`和`external_action`二选一[【附录-枚举值-转化类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	FlowControlMode    string `json:"flow_control_mode,omitempty"`    // 竞价策略(投放方式),[【附录-计划投放速度类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528#custom-anchor-计划投放速度类型)
	SmartBidType       string `json:"smart_bid_type,omitempty"`       // 投放场景(出价方式)，[【附录-自动出价类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528#custom-anchor-自动出价类型)
}
