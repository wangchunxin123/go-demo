/* #
# Author: wangchunxin
# Created Time: 2020/12/2 3:56 下午
# File Name:
# Description:
# */
package model

//获取计划返回
type PlanGetReturn struct {
	CommonReturn
	Data PlanGetDataReturn `json:"data"`
}
type PlanGetDataAudienceReturn struct {
	Geolocation               []interface{} `json:"geolocation"`
	InterestCategories        []int64       `json:"interest_categories"`
	InterestActionMode        string        `json:"interest_action_mode"`
	UserType                  []interface{} `json:"user_type"`
	ExcludeCustomActions      []interface{} `json:"exclude_custom_actions"`
	FilterAwemeFansCount      int           `json:"filter_aweme_fans_count"`
	FilterOwnAwemeFans        int           `json:"filter_own_aweme_fans"`
	FlowPackage               []int64       `json:"flow_package"`
	DeviceType                interface{}   `json:"device_type"`
	OwnAwemeNumber            interface{}   `json:"own_aweme_number"`
	InterestTags              []interface{} `json:"interest_tags"`
	AndroidOsv                string        `json:"android_osv"`
	City                      []interface{} `json:"city"`
	District                  string        `json:"district"`
	AutoExtendTargets         interface{}   `json:"auto_extend_targets"`
	AwemeFanBehaviors         []interface{} `json:"aweme_fan_behaviors"`
	BusinessIds               []interface{} `json:"business_ids"`
	IncludeCustomActions      []interface{} `json:"include_custom_actions"`
	Platform                  []string      `json:"platform"`
	ActivateType              []interface{} `json:"activate_type"`
	IosOsv                    string        `json:"ios_osv"`
	AdTag                     []interface{} `json:"ad_tag"`
	AwemeFansNumbers          []interface{} `json:"aweme_fans_numbers"`
	Ac                        []interface{} `json:"ac"`
	AppCategory               []interface{} `json:"app_category"`
	AutoExtendEnabled         int           `json:"auto_extend_enabled"`
	RetargetingTagsExclude    []interface{} `json:"retargeting_tags_exclude"`
	ArticleCategory           []interface{} `json:"article_category"`
	FilterAwemeAbnormalActive interface{}   `json:"filter_aweme_abnormal_active"`
	RetargetingTags           []interface{} `json:"retargeting_tags"`
	LocationType              string        `json:"location_type"`
	SuperiorPopularityType    string        `json:"superior_popularity_type"`
	AppBehaviorTarget         string        `json:"app_behavior_target"`
	LaunchPrice               []interface{} `json:"launch_price"`
	RetargetingTagsInclude    []interface{} `json:"retargeting_tags_include"`
	Age                       []interface{} `json:"age"`
	RetargetingType           string        `json:"retargeting_type"`
	InterestWords             []interface{} `json:"interest_words"`
	AwemeFanCategories        []interface{} `json:"aweme_fan_categories"`
	ExcludeFlowPackage        []interface{} `json:"exclude_flow_package"`
	Carrier                   []interface{} `json:"carrier"`
	DpaLocalAudience          interface{}   `json:"dpa_local_audience"`
	Gender                    string        `json:"gender"`
	Action                    interface{}   `json:"action"`
	DeviceBrand               []interface{} `json:"device_brand"`
	AwemeFanAccounts          []interface{} `json:"aweme_fan_accounts"`
	AppIds                    []interface{} `json:"app_ids"`
}
type PlanGetDataListReturn struct {
	BudgetMode              string                    `json:"budget_mode"`
	OpenURL                 string                    `json:"open_url"`
	DpaProvince             int                       `json:"dpa_province"`
	AdModifyTime            string                    `json:"ad_modify_time"`
	AppDesc                 string                    `json:"app_desc"`
	UnionVideoType          string                    `json:"union_video_type"`
	UniqueFk                string                    `json:"unique_fk"`
	AuditRejectReason       string                    `json:"audit_reject_reason"`
	CampaignID              int64                     `json:"campaign_id"`
	InventoryType           []string                  `json:"inventory_type"`
	DpaOpenURLType          string                    `json:"dpa_open_url_type"`
	ExternalURLParams       string                    `json:"external_url_params"`
	AdvertiserID            int64                     `json:"advertiser_id"`
	Pricing                 string                    `json:"pricing"`
	GamePackageThumbnailIds []string                  `json:"game_package_thumbnail_ids"`
	PromotionType           string                    `json:"promotion_type"`
	AdCreateTime            string                    `json:"ad_create_time"`
	DpaCategories           []int                     `json:"dpa_categories"`
	ScheduleType            string                    `json:"schedule_type"`
	OpenURLParams           string                    `json:"open_url_params"`
	CpaBid                  float64                   `json:"cpa_bid"`
	ConvertID               int                       `json:"convert_id"`
	AdjustCpa               int                       `json:"adjust_cpa"`
	DeepCpabid              float64                   `json:"deep_cpabid"`
	AppType                 string                    `json:"app_type"`
	FlowControlMode         string                    `json:"flow_control_mode"`
	AwemeAccount            string                    `json:"aweme_account"`
	StoreType               string                    `json:"store_type"`
	SubscribeURL            string                    `json:"subscribe_url"`
	GamePackageDesc         string                    `json:"game_package_desc"`
	DpaOpenURLField         string                    `json:"dpa_open_url_field"`
	DownloadURL             string                    `json:"download_url"`
	ID                      int64                     `json:"id"`
	DpaProducts             []int                     `json:"dpa_products"`
	CategoryType            string                    `json:"category_type"`
	SmartBidType            string                    `json:"smart_bid_type"`
	DpaRecommendType        int                       `json:"dpa_recommend_type"`
	AdvancedCreativeType    string                    `json:"advanced_creative_type"`
	DpaExternalURLField     string                    `json:"dpa_external_url_field"`
	DpaExternalUrls         []string                  `json:"dpa_external_urls"`
	StartTime               string                    `json:"start_time"`
	OptStatus               string                    `json:"opt_status"`
	EndTime                 string                    `json:"end_time"`
	Status                  string                    `json:"status"`
	FormIndex               int                       `json:"form_index"`
	StoreproUnit            string                    `json:"storepro_unit"`
	FormID                  int                       `json:"form_id"`
	AdGroupID               interface{}               `json:"ad_group_id"`
	Bid                     float64                   `json:"bid"`
	ConvertedTimeDuration   string                    `json:"converted_time_duration"`
	DpaCity                 int                       `json:"dpa_city"`
	LearningPhase           string                    `json:"learning_phase"`
	QuickAppURL             string                    `json:"quick_app_url"`
	AudiencePackageID       int64                     `json:"audience_package_id"`
	DeepBidType             string                    `json:"deep_bid_type"`
	DpaAdtype               []string                  `json:"dpa_adtype"`
	AppIntroduction         string                    `json:"app_introduction"`
	AdID                    int64                     `json:"ad_id"`
	Audience                PlanGetDataAudienceReturn `json:"audience"`
	AppThumbnails           []string                  `json:"app_thumbnails"`
	ExternalActions         []string                  `json:"external_actions"`
	RoiGoal                 float64                   `json:"roi_goal"`
	DpaOpenUrls             []string                  `json:"dpa_open_urls"`
	HideIfConverted         string                    `json:"hide_if_converted"`
	Name                    string                    `json:"name"`
	Package                 string                    `json:"package"`
	ScheduleTime            string                    `json:"schedule_time"`
	DpaProductTarget        interface{}               `json:"dpa_product_target"`
	GamePackageBatchID      int64                     `json:"game_package_batch_id"`
	Budget                  float64                   `json:"budget"`
	HideIfExists            int                       `json:"hide_if_exists"`
	DeliveryRange           string                    `json:"delivery_range"`
	IntelligentFlowSwitch   interface{}               `json:"intelligent_flow_switch"`
	DownloadMode            string                    `json:"download_mode"`
	DpaLbs                  int                       `json:"dpa_lbs"`
	ModifyTime              string                    `json:"modify_time"`
	AdvertiserStoreIds      []int64                   `json:"advertiser_store_ids"`
	LubanRoiGoal            float64                   `json:"luban_roi_goal"`
	ProductPlatformID       int64                     `json:"product_platform_id"`
	StoreproPackID          int64                     `json:"storepro_pack_id"`
	DownloadType            string                    `json:"download_type"`
	ProductID               int64                     `json:"product_id"`
	ParamsType              string                    `json:"params_type"`
	ExternalURL             string                    `json:"external_url"`
}
type PlanGetDataReturn struct {
	PageInfo PageInfoReturn          `json:"page_info"`
	List     []PlanGetDataListReturn `json:"list"`
}

//更新计划状态的返回结果
type PlanUpdateStatusReturn struct {
	CommonReturn
	Data struct {
		AdIds []int64 `json:"ad_ids"`
	} `json:"data"`
}

//更新计划预算的返回结果
type PlanUpdateBudgetReturn struct {
	CommonReturn
	Data struct {
		AdIds []int64 `json:"ad_ids"`
	}
}

//更新计划预算的返回结果
type PlanUpdateBidReturn struct {
	CommonReturn
	Data struct {
		AdIds []int64 `json:"ad_ids"`
	}
}

//获取计划审核建议的返回数据的数据结构
type PlanRejectReasonReturn struct {
	CommonReturn
	Data AdRejectReasonResList `json:"data"`
}

type AdRejectReasonResList struct {
	List []AdRejectReasonResListInfo `json:"list"`
}
type AdRejectReasonResListInfo struct {
	AdReject       AdRejectReasonResListInfoPlan             `json:"ad_reject"`
	CreativeReject []AdRejectReasonMsgResDataCreativeReject  `json:"creative_reject"`
	IsProcedualAd  int                                       `json:"is_procedual_ad"`
	MaterialReject []AdRejectReasonResListInfoMaterialReject `json:"material_reject"`
}

type AdRejectReasonResListInfoPlan struct {
	AdId       int64             `json:"ad_id"`
	RejectData map[string]string `json:"reject_data"`
}

//创意审核建议creative_reject
type AdRejectReasonMsgResDataCreativeReject struct {
	CreativeId     int64                                        `json:"creative_id"`
	RejectData     []AdRejectReasonMsgResDataCreativeRejectData `json:"reject_data"`
	MaterialReject []AdRejectReasonResListInfoMaterialReject    `json:"material_reject"`
}
type AdRejectReasonMsgResDataCreativeRejectData struct {
	RejectItem   string `json:"reject_item"`
	RejectReason string `json:"reject_reason"`
}

//程序化素材审核意见
type AdRejectReasonResListInfoMaterialReject struct {
	MaterialType int      `json:"material_type"`
	Title        string   `json:"title"`
	ImageId      []string `json:"image_id"`
	VideoId      string   `json:"video_id"`
	RejectReason string   `json:"reject_reason"`
	AdvertiserId int64    `json:"advertiser_id"`
	PlanId       int64    `json:"plan_id"`
	CreativeId   int64    `json:"creative_id"`
}
