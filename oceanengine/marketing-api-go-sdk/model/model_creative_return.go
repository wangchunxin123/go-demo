/* #
# Author: wangchunxin
# Created Time: 2020/12/3 9:54 上午
# File Name:
# Description:
# */
package model

//获取创意列表
type CreativeGetReturn struct {
	CommonReturn
	Data CreativeGetDataReturn `json:"data"`
}

type CreativeGetDataReturn struct {
	PageInfoReturn
	List []CreativeGetDataListReturn `json:"list"`
}

type CreativeGetDataListReturn struct {
	Status             string        `json:"status"`
	CreativeWordIds    []int64       `json:"creative_word_ids"`
	VideoID            string        `json:"video_id"`
	AdID               int64         `json:"ad_id"`
	CreativeCreateTime string        `json:"creative_create_time"`
	Title              string        `json:"title"`
	CreativeModifyTime string        `json:"creative_modify_time"`
	CreativeID         int64         `json:"creative_id"`
	AuditRejectReason  interface{}   `json:"audit_reject_reason"`
	AdvertiserID       int64         `json:"advertiser_id"`
	ImageID            string        `json:"image_id"`
	Metarials          []interface{} `json:"metarials"`
	ThirdPartyID       string        `json:"third_party_id"`
	ImageMode          string        `json:"image_mode"`
	ImageIds           []string      `json:"image_ids"`
	BidwordList        interface{}   `json:"bidword_list"`
	OptStatus          string        `json:"opt_status"`
}

type CreativeGetDataListMaterialReturn struct {
	Title   string `json:"title"`
	ImageId string `json:"image_id"`
	VideoId string `json:"video_id"`
}

//获取创意详情
type CreativeReadReturn struct {
	CommonReturn
	Data CreativeReadDataReturn `json:"data"`
}
type CreativeReadCreativesReturn struct {
	Title           string   `json:"title"`
	CreativeID      int64    `json:"creative_id"`
	DerivePosterCid int      `json:"derive_poster_cid"`
	CreativeWordIds []int    `json:"creative_word_ids"`
	ImageIds        []string `json:"image_ids"`
	DpaTemplate     int      `json:"dpa_template"`
	ImageMode       string   `json:"image_mode"`
}
type CreativeReadDataReturn struct {
	ModifyTime                 string                        `json:"modify_time"`
	AppName                    string                        `json:"app_name"`
	GenerateDerivedAd          int                           `json:"generate_derived_ad"`
	TrackURL                   string                        `json:"track_url"`
	AbstractList               []interface{}                 `json:"abstract_list"`
	AdKeywords                 []string                      `json:"ad_keywords"`
	WebURL                     string                        `json:"web_url"`
	ThirdIndustryID            int                           `json:"third_industry_id"`
	IsSmartTitle               int                           `json:"is_smart_title"`
	InventoryType              []string                      `json:"inventory_type"`
	AdvancedCreativeType       string                        `json:"advanced_creative_type"`
	AdCategory                 int                           `json:"ad_category"`
	SceneInventory             interface{}                   `json:"scene_inventory"`
	DynamicCreativeSwitch      []interface{}                 `json:"dynamic_creative_switch"`
	AdID                       int64                         `json:"ad_id"`
	ImageList                  []interface{}                 `json:"image_list"`
	Creatives                  []CreativeReadCreativesReturn `json:"creatives"`
	AdvertiserID               int64                         `json:"advertiser_id"`
	SubLinkIDList              []interface{}                 `json:"sub_link_id_list"`
	ActionText                 string                        `json:"action_text"`
	TrackURLSendType           string                        `json:"track_url_send_type"`
	CloseVideoDetail           int                           `json:"close_video_detail"`
	ActionTrackURL             string                        `json:"action_track_url"`
	IsCommentDisable           int                           `json:"is_comment_disable"`
	CreativeDisplayMode        string                        `json:"creative_display_mode"`
	VideoPlayDoneTrackURL      string                        `json:"video_play_done_track_url"`
	TitleList                  []interface{}                 `json:"title_list"`
	VideoPlayEffectiveTrackURL string                        `json:"video_play_effective_track_url"`
	SmartInventory             int                           `json:"smart_inventory"`
	VideoPlayTrackURL          string                        `json:"video_play_track_url"`
}

//更改创意的状态返回值
type CreativeUpdateStatusReturn struct {
	CommonReturn
	Data struct {
		CreativeIds []int64 `json:"creative_ids"`
	}
}

//创意素材信息返回值
type CreativeMaterialReadReturn struct {
	CommonReturn
	Data []struct {
		Status       string `json:"status"`
		AdID         int64  `json:"ad_id"`
		Title        string `json:"title"`
		AdvertiserID int64  `json:"advertiser_id"`
		ImageMode    string `json:"image_mode"`
		ImageInfo    []struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"image_info"`
		ID        int64  `json:"id"`
		OptStatus string `json:"opt_status"`
	} `json:"data"`
}

//创意审核信息的返回值
type CreativeRejectReasonReturn struct {
	CommonReturn
	Data struct {
		List []struct {
			CreativeID     int64 `json:"creative_id"`
			MaterialReject []struct {
				ImageID      []string `json:"image_id"`
				VideoID      string   `json:"video_id"`
				MaterialType int      `json:"material_type"`
				RejectReason string   `json:"reject_reason"`
				Title        string   `json:"title,omitempty"`
			} `json:"material_reject"`
			RejectData []struct {
				RejectItem   string `json:"reject_item"`
				RejectReason string `json:"reject_reason"`
			} `json:"reject_data"`
		} `json:"list"`
	} `json:"data"`
}
