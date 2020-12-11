/* #
# Author: wangchunxin
# Created Time: 2020/12/8 2:27 下午
# File Name:
# Description:
# */
package model

// 创建图片转视频任务
type ToolsPicToVideoTaskCreateRequest struct {
	AdvertiserId  int64       `json:"advertiser_id,omitempty"`   // 操作的广告主id
	TemplateType  string      `json:"template_type,omitempty"`   // 模版类型
	MusicName     string      `json:"music_name,omitempty"`      // 音乐名称
	MapIdResource interface{} `json:"map_id_resource,omitempty"` // 资源信息
	CallbackUrl   string      `json:"callback_url,omitempty"`    // 回调地址，详见视频任务回调接口
	OtherInfo     interface{} `json:"other_info,omitempty"`      // 其他信息
}

// 获取音乐列表
type ToolsPicToVideoMusicGetRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 操作的广告主id
}

// 获取任务状态
type ToolsPicToVideoStatusGetRequest struct {
	AdvertiserId int64  `json:"advertiser_id,omitempty"` // 操作的广告主id
	VideoId      string `json:"video_id,omitempty"`      // 视频id
}
