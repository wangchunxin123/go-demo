/* #
# Author: wangchunxin
# Created Time: 2020/12/10 2:01 下午
# File Name:
# Description:
# */
package model

// 获取兴趣关键词词包
type ToolsInterestTagsSelectReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		List []struct {
			Id    int    `json:"id"`   // 兴趣词包id
			Name  string `json:"name"` // 兴趣词包名称
			Words []struct {
				WordId int    `json:"word_id"` // 关键词id
				Word   string `json:"word"`    // 关键词名称
			} `json:"words"` // 兴趣词包关键词
		} `json:"list"` // 兴趣词包列表
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 创建兴趣词包
type ToolsInterestTagsCreateReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		Id int `json:"id"` // 词包id
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 更新兴趣词包
type ToolsInterestTagsUpdateReturn struct {
	Code      int         `json:"code"`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message   string      `json:"message"`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data      interface{} `json:"data"`       // json返回值
	RequestId string      `json:"request_id"` // 请求日志id
}

// 删除兴趣词包
type ToolsInterestTagsDeleteReturn struct {
	Code      int         `json:"code"`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message   string      `json:"message"`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data      interface{} `json:"data"`       // json返回值
	RequestId string      `json:"request_id"` // 请求日志id
}

// 兴趣关键词ID转词
type ToolsInterestTagsId2WordReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		IdWordDict interface{} `json:"id_word_dict"` // 兴趣关键词内容
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 兴趣关键词词转ID
type ToolsInterestTagsWord2IdReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		WordIdDict interface{} `json:"word_id_dict"` // 兴趣关键词内容
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}
