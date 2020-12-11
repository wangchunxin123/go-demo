/* #
# Author: wangchunxin
# Created Time: 2020/12/10 2:47 下午
# File Name:
# Description:
# */
package model

// 创建转化ID
type ToolsAdConvertCreateReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		Id              int    `json:"id"`                // 转化id
		OptStatus       string `json:"opt_status"`        // 操作状态，详见[【附录-操作状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
		Status          string `json:"status"`            // 转化状态，详见[【附录-转化状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
		ConvertDataType string `json:"convert_data_type"` // 转化统计方式，广告平台统计该转化目标是否发生的方式
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 修改转化ID
type ToolsAdConvertUpdateReturn struct {
	Code    int         `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string      `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    interface{} `json:"data"`    // json返回值
}

// 更新转化状态
type ToolsAdConvertUpdateStatusReturn struct {
	Code    int         `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string      `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    interface{} `json:"data"`    // json返回值
}

// 转化ID列表
type ToolsAdvConvertSelectReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		PageInfo struct {
			CurrentPage int `json:"current_page"` // 当前页数
			TotalPage   int `json:"total_page"`   // 总页数
		} `json:"page_info"` // 分页相关信息
		AdConvertList []struct {
			Id                int    `json:"id"`                  // 转化id，其中较小数值convert_id为预定义转化，具体枚举可查看[【附录-预定义转化类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
			Name              string `json:"name"`                // 转化名称
			OptStatus         string `json:"opt_status"`          // 操作状态，详见[【附录-操作状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
			ConvertSourceType string `json:"convert_source_type"` // 转化来源，详见[【附录-转化来源】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
			Status            string `json:"status"`              // 转化状态，详见[【附录-转化状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
			ConvertType       string `json:"convert_type"`        // 转化类型，详见[【附录-转化类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
			ActionTrackUrl    string `json:"action_track_url"`    // 转化监测链接
		} `json:"ad_convert_list"` // 转化的数据list
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 查询转化详细信息
type ToolsAdConvertReadReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		AdvertiserId               int64  `json:"advertiser_id"`                  // 广告主id
		Id                         int64    `json:"id"`                             // 转化id
		Name                       string `json:"name"`                           // 转化名称
		AppType                    string `json:"app_type"`                       // 应用类型允许值: `"APP_ANDROID"`, `"APP_IOS"`
		PackageName                string `json:"package_name"`                   // 包名
		DownloadUrl                string `json:"download_url"`                   // 下载地址
		OptStatus                  string `json:"opt_status"`                     // 操作状态，详见[【附录-操作状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
		ConvertSourceType          string `json:"convert_source_type"`            // 转化来源，详见[【附录-转化来源】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
		Status                     string `json:"status"`                         // 转化状态，详见[【附录-转化状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
		ConvertType                string `json:"convert_type"`                   // 转化类型，详见[【附录-转化类型】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
		ActionTrackUrl             string `json:"action_track_url"`               // 点击监测链接
		DisplayTrackUrl            string `json:"display_track_url"`              // 展示监测链接
		VideoPlayEffectiveTrackUrl string `json:"video_play_effective_track_url"` // 视频有效播放监测链接
		VideoPlayDoneTrackUrl      string `json:"video_play_done_track_url "`     // 视频播放完毕监测链接
		VideoPlayTrackUrl          string `json:"video_play_track_url "`          // 视频播放监测链接
		ConvertActivateCallbackUrl string `json:"convert_activate_callback_url"`  // 激活回传地址
		ConvertSecretKey           string `json:"convert_secret_key"`             // 加密密钥
		AppId                      string `json:"app_id"`                         // APP ID
		ExternalUrl                string `json:"external_url"`                   // 落地页链接
		ConvertTrackParams         string `json:"convert_track_params"`           // 监测参数
		ConvertBaseCode            string `json:"convert_base_code"`              // 转化基础代码
		ConvertJsCode              string `json:"convert_js_code"`                // 转化代码（JS方式）
		ConvertHtmlCode            string `json:"convert_html_code"`              // 转化代码（HTML方式）
		ConvertXpathUrl            string `json:"convert_xpath_url"`              // 转化页面
		ConvertXpathValue          string `json:"convert_xpath_value"`            // 转化路径
		OpenUrl                    string `json:"open_url"`                       // 直达链接
		CreateTime                 string `json:"create_time"`                    // 创建时间
		ModifyTime                 string `json:"modify_time"`                    // 更新时间
		IgnoreParams               interface{}   `json:"ignore_params"`                  // 转化类型下匹配规则字段
		ConvertDataType            string `json:"convert_data_type"`              // 转化统计方式
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 转化ID推送
type ToolsAdConvertPushReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct{
		Data    struct {
			Id int `json:"id"` // 转化id
		} `json:"data"`
	}   `json:"data"`    // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 查询深度优化方式
type ToolsAdConvertDeepbidReadReturn struct {
	Code      int    `json:"code"`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message   string `json:"message"`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	RequestId string `json:"request_id"` // 请求日志id
	Data      struct {
		SuccessList []string `json:"success_list"` // 可用的深度转化方式列表[【附录：枚举值-深度优化方式】](https://ad.oceanengine.com/openapi/doc/index.html?id=528)
	} `json:"data"` // 返回数据
}
