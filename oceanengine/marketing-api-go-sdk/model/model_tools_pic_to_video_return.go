/* #
# Author: wangchunxin
# Created Time: 2020/12/10 2:06 下午
# File Name:
# Description:
# */
package model

// 创建图片转视频任务
type ToolsPicToVideoTaskCreateReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		VideoId int `json:"video_id"` // 视频id，需要根据video_id调用任务状态接口，当状态为SUCCESS时，可以调用查询视频信息接口预览
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 获取音乐列表
type ToolsPicToVideoMusicGetReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		Status string      `json:"status"` // 音乐数据
		Musics interface{} `json:"musics"` // 音乐数据
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 获取任务状态
type ToolsPicToVideoStatusGetReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		Status string `json:"status"` // 异步任务状态,详见[【附录-异步任务状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}
