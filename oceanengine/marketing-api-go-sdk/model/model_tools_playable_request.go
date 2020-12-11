/* #
# Author: wangchunxin
# Created Time: 2020/12/8 12:03 下午
# File Name:
# Description:
# */
package model

import "mime/multipart"

// 试玩素材上传
type ToolsPlayableUploadRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	PlayablePackage *multipart.File `json:"playable_package,omitempty"` // 试玩素材文件【格式说明】：1.格式为zip，大小不超过3M2.已接入穿山甲JS-SDK，并且调用方式为window.openAppStore();3.
	// 主html文件需命名为index，且位于一级目录中4.试玩广告播放方向字段应存储于config.json文件中，位于一级目录中，取值为0,1,25.文件名称仅支持大小写字母、数字、减号和下划线6.素材中不允许使用 mraid.js 格式7.素材不允许通过外部网络加载动态素材8.素材中不允许包含JS 重定向9.素材不允许发出 HTTP 请求10.素材中不允许存在内容为空的文件
}

// 试玩素材上传校验结果
type ToolsPlayableValidateRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	PlayableId int `json:"playable_id,omitempty"` // 试玩素材ID
}

// 保存试玩素材
type ToolsPlayableSaveRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	PlayableId int `json:"playable_id,omitempty"` // 试玩素材ID
	PlayableName string `json:"playable_name,omitempty"` // 试玩素材名称
}

// 获取试玩素材列表
type ToolsPlayableListGetRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	PlayableUrl string `json:"playable_url,omitempty"` // 试玩素材url，过滤条件
	Status string `json:"status,omitempty"` // 试玩素材的状态，详情见[【附录-试玩素材状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值：`"VALIDATING"` ,`"VALIDATE_FAIL"`,`"VALIDATE_SUCCESS"`,`"AUDIT_FAI"`,`"AUDIT_SUCCESS"`
	Page int `json:"page,omitempty"` // 页数默认值: `1`
	PageSize int `json:"page_size,omitempty"` // 页面大小默认值: `10`
}

// 获取试玩素材可推送列表
type ToolsPlayableGrantableAdvListRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
}

// 推送试玩素材
type ToolsPlayableGrantRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	PlayableId int `json:"playable_id,omitempty"` // 需要推送的试玩素材ID，与playable_url有且仅有一个
	PlayableUrl string `json:"playable_url,omitempty"` // 需要推送的试玩素材url，与playable_id有且仅有一个
	GrantedIds []int `json:"granted_ids,omitempty"` // 推送的目标（广告主id）列表，一次最多同时推送给50个广告主
}

// 获取试玩素材推送结果
type ToolsPlayableGrantResultRequest struct {
	AdvertiserId int64 `json:"advertiser_id,omitempty"` // 广告主ID
	TaskIds []int `json:"task_ids,omitempty"` // 推送任务id，可根据推送的task_id进行检索
	StartTime string `json:"start_time,omitempty"` // 根据任务创建时间进行过滤，过滤的起始时间，格式YYYY-mm-dd
	EndTime string `json:"end_time,omitempty"` // 根据任务的创建时间进行过滤，过滤的截止时间，格式YYYY-mm-dd
	Page int `json:"page,omitempty"` // 页码，默认1
	PageSize int `json:"page_size,omitempty"` // 页面大小，默认20，最大200
}

