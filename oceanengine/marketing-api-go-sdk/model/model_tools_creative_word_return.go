/* #
# Author: wangchunxin
# Created Time: 2020/12/10 2:14 下午
# File Name:
# Description:
# */
package model

// 创建动态创意词包
type ToolsCreativeWordCreateReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		CreativeWordId int `json:"creative_word_id"` // 创意词包ID
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 查询动态创意词包
type ToolsCreativeWordSelectReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		CreativeWord []struct {
			CreativeWordId int      `json:"creative_word_id"` // 创意词包ID
			Name           string   `json:"name"`             // 创意词包名称
			DefaultWord    string   `json:"default_word"`     // 默认词
			Words          []string `json:"words"`            // 替换词
			ContentType    string   `json:"content_type"`     // 创意词包类型，详见[【附件-创意词包类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
			MaxWordLen     int      `json:"max_word_len"`     // 替换词最大长度
			Status         string   `json:"status"`           // 创意词包状态，详见[【附件-创意词包状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
			Rate           float64  `json:"rate"`             // 创意词包人群覆盖率取值范围: `0-1`
		} `json:"creative_word"` // 创意词包列表
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 更新动态创意词包
type ToolsCreativeWordUpdateReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		CreativeWordId int `json:"creative_word_id"` // 创意词包ID
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 删除动态创意词包
type ToolsCreativeWordDeleteReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		CreativeWordId int `json:"creative_word_id"` // 创意词包ID
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}
