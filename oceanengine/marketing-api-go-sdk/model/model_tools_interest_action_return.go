/* #
# Author: wangchunxin
# Created Time: 2020/12/10 12:01 下午
# File Name:
# Description:
# */
package model

// 行为类目查询
type ToolsInterestActionActionCategoryReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		Num      string `json:"num"` // 数量
		Children struct {
			Num      string `json:"num"` // 数量
			Children struct {
				Num      string `json:"num"` // 数量
				Children struct {
					Num  string      `json:"num"`  // 数量
					Id   interface{} `json:"id"`   // 行为类目id
					Name string      `json:"name"` // 行为类目
				} `json:"children"` // 行为子类目
			} `json:"children"` // 行为子类目
		} `json:"children"` // 行为子类目
	} `json:"data"` // json返回值

	RequestId string `json:"request_id"` // 请求日志id
}

// 行为关键词查询
type ToolsInterestActionActionKeywordReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		List []struct {
			Id   int64  `json:"id"`   // 关键词id
			Name string `json:"name"` // 关键词名称
			Num  int    `json:"num"`  // 关键词数目
		} `json:"list"` // 词包列表
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 兴趣类目查询
type ToolsInterestActionInterestCategoryReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    []struct {
		Num      string `json:"num"` // 数量
		Children []struct {
			Num      string `json:"num"` // 数量
			Children []struct {
				Num      string `json:"num"` // 数量
				Children []struct {
					Num  string      `json:"num"`  // 数量
					Id   interface{} `json:"id"`   // 行为类目id
					Name string      `json:"name"` // 行为类目
				} `json:"children"` // 行为子类目
			} `json:"children"` // 行为子类目
		} `json:"children"` // 行为子类目
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 兴趣关键词查询
type ToolsInterestActionInterestKeywordReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		List []struct {
			Id   interface{} `json:"id"`   // 关键词id
			Name string      `json:"name"` // 关键词名称
			Num  interface{} `json:"num"`  // 关键词数目
		} `json:"list"` // 词包列表
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 兴趣行为类目关键词id转词
type ToolsInterestActionId2WordReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		Categories []struct {
			Id   int    `json:"id"`   // 类目ID
			Name string `json:"name"` // 类目名称
			Num  int    `json:"num"`  // 覆盖人数
		} `json:"categories"` // 类目列表
		Keywords []struct {
			Id   int    `json:"id"`   // 关键词ID
			Name string `json:"name"` // 关键词名称
			Num  int    `json:"num"`  // 覆盖人数
		}
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 获取行为兴趣推荐关键词
type ToolsInterestActionKeywordSuggestReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		Keywords []struct{
			Id        int    `json:"id"`         // 关键词ID
			Name      string `json:"name"`       // 关键词名称
			Num       int    `json:"num"`        // 覆盖人数
		} `json:"keywords"` // 关键词列表
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}
