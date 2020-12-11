/* #
# Author: wangchunxin
# Created Time: 2020/12/10 3:28 下午
# File Name:
# Description:
# */
package model

// 启动一键起量
type ToolsAdRaiseSetReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		Status string `json:"status"` // 一键起量状态,详见[【附录-一键起量状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"ENABLE_RAISE"`, `"DISABLE_RAISE"`, `"RAISING"`, `"FINISH_RAISE"`
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 获取起量预估值
type ToolsAdRaiseEstimateGetReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息
	Data    struct {
		EstimateNum int `json:"estimate_num"` //  预估展示量
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 获取当前起量状态
type ToolsAdRaiseStatusGetReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息
	Data    struct {
		Status string `json:"status"` // 一键起量状态,详见[【附录-一键起量状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)允许值: `"ENABLE_RAISE"`, `"DISABLE_RAISE"`, `"RAISING"`, `"FINISH_RAISE"`
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 获取一键起量的后验数据
type ToolsAdRaiseResultGetReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息
	Data    struct {
		StartTime string  `json:"start_time"` // 一键起量开始时间
		EndTime   string  `json:"end_time"`   // 一键起量结束时间
		Cost      int     `json:"cost"`       //  一键起量阶段产生消耗
		Show      int     `json:"show"`       //  一键起量阶段产生展示
		Click     int     `json:"click"`      //  一键起量阶段产生点击数
		Convert   int     `json:"convert"`    //  一键起量阶段产生转换数
		Ctr       float64 `json:"ctr"`        //  CTR 一键起量期间点击率
		Cvr       float64 `json:"cvr"`        //  CVR 一键起量期间转化率
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}
