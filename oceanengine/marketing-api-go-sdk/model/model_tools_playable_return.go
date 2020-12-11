/* #
# Author: wangchunxin
# Created Time: 2020/12/10 11:50 上午
# File Name:
# Description:
# */
package model

// 试玩素材上传
type ToolsPlayableUploadReturn struct {
	Code    int    `json:"code,omitempty"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message,omitempty"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data struct{
		PlayableId int64 `json:"playable_id"`
	} `json:"data"`
	RequestId string `json:"request_id,omitempty"` // 请求日志id
}

// 试玩素材上传校验结果
type ToolsPlayableValidateReturn struct {
	Code    int    `json:"code,omitempty"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message,omitempty"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		PlayableId          int    `json:"playable_id,omitempty"`          // 试玩素材ID
		PlayableUrl         string `json:"playable_url,omitempty"`         // 试玩素材url
		PlayableOrientation string `json:"playable_orientation,omitempty"` // 试玩素材方向，详情见[【附录-试玩素材方向】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
		Status              string `json:"status,omitempty"`               // 试玩素材状态，详情见[【附录-试玩素材状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
		FailMessage         string `json:"fail_message,omitempty"`         // 验证失败原因
	} `json:"data,omitempty"`                      // json返回值
	RequestId string `json:"request_id,omitempty"` // 请求日志id
}

// 保存试玩素材
type ToolsPlayableSaveReturn struct {
	Code    int    `json:"code,omitempty"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message,omitempty"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		PlayableId          int    `json:"playable_id,omitempty"`          // 试玩素材ID
		PlayableUrl         string `json:"playable_url,omitempty"`         // 试玩素材url
		PlayableName        string `json:"playable_name,omitempty"`        // 试玩素材名
		PlayableOrientation string `json:"playable_orientation,omitempty"` // 试玩素材方向，详情见[【附录-试玩素材方向】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
		Status              string `json:"status,omitempty"`               // 试玩素材状态，详情见[【附录-试玩素材状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	} `json:"data,omitempty"`                      // json返回值
	RequestId string `json:"request_id,omitempty"` // 请求日志id
}

// 获取试玩素材列表
type ToolsPlayableListGetReturn struct {
	Code    int    `json:"code,omitempty"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message,omitempty"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		List interface{} `json:"list,omitempty"` // 试玩列表
	} `json:"data,omitempty"`                                                  // json返回值
	PlayableId          int            `json:"playable_id,omitempty"`          // 试玩素材ID
	PlayableUrl         string         `json:"playable_url,omitempty"`         // 试玩素材url
	PlayableName        string         `json:"playable_name,omitempty"`        // 试玩素材名
	PlayableOrientation string         `json:"playable_orientation,omitempty"` // 试玩素材方向，详情见[【附录-试玩素材方向】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	Status              string         `json:"status,omitempty"`               // 试玩素材状态，详情见[【附录-试玩素材状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	PageInfo            PageInfoReturn `json:"page_info"`
}

// 获取试玩素材可推送列表
type ToolsPlayableGrantableAdvListReturn struct {
	Code    int    `json:"code,omitempty"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message,omitempty"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		List []struct {
			AdvertiserId   int64  `json:"advertiser_id,omitempty"`   // 广告主ID
			AdvertiserName string `json:"advertiser_name,omitempty"` // 广告主名
		} `json:"list,omitempty"` // 可推送的广告主列表
	} `json:"data,omitempty"`                      // json返回值
	RequestId string `json:"request_id,omitempty"` // 请求日志id
}

// 推送试玩素材
type ToolsPlayableGrantReturn struct {
	Code    int    `json:"code,omitempty"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message,omitempty"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		List []struct {
			GrantedId int `json:"granted_id,omitempty"` // 推送目标（广告主ID）
			TaskId    int `json:"task_id,omitempty"`    // 推送任务id，可通过【查询推送结果】接口获取任务推送结果
		} `json:"list,omitempty"` // 推送任务列表
	} `json:"data,omitempty"`                      // json返回值
	RequestId string `json:"request_id,omitempty"` // 请求日志id
}

// 获取试玩素材推送结果
type ToolsPlayableGrantResultReturn struct {
	Code    int    `json:"code,omitempty"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message,omitempty"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		List []struct {
			TaskId         int    `json:"task_id,omitempty"`          // 任务id
			PlayableId     int    `json:"playable_id,omitempty"`      // 被推送的试玩素材id
			PlayableUrl    string `json:"playable_url,omitempty"`     // 被推送的试玩素材url
			GrantedId      int    `json:"granted_id,omitempty"`       // 推送目标（广告主id）
			Status         string `json:"status,omitempty"`           // 任务推送结果枚举值：`RUNNING`（推送中）,`SUCCESS`（推送成功）,`FAILED`（推送失败）
			NewPlayableId  int    `json:"new_playable_id,omitempty"`  // 推送成功后新生成的试玩素材ID
			NewPlayableUrl string `json:"new_playable_url,omitempty"` // 推送成功后新生成的试玩素材url
			CreateTime     string `json:"create_time,omitempty"`      // 推送任务创建的时间，格式：2020-06-03 16:08:47
			RequestId      string `json:"request_id,omitempty"`       // 请求日志id
		} `json:"list,omitempty"` // 推送结果信息列表
	} `json:"data,omitempty"` // json返回值

}
