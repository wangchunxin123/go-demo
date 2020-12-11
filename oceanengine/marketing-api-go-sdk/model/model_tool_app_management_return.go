/* #
# Author: wangchunxin
# Created Time: 2020/12/10 3:30 下午
# File Name:
# Description:
# */
package model

// 查询游戏信息
type ToolsAppManagementBookingGetReturn struct {
	Code    int    `json:"code "`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message "` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		List []struct {
			PackageId            string `json:"package_id "`             // 应用ID
			PackageName          string `json:"package_name "`           // 包名
			AppCloudId           int    `json:"app_cloud_id "`           // app id
			AppName              string `json:"app_name "`               // 应用名
			Version              string `json:"version "`                // 版本号
			DownloadUrl          string `json:"download_url "`           // 下载地址
			IconUrl              string `json:"icon_url "`               // icon地址
			PublishTime          string `json:"publish_time "`           // 发布时间，格式：%Y-%m-%d %H:%M:%S
			ScheduledPublishTime string `json:"scheduled_publish_time "` // 预约发布时间，格式：%Y-%m-%d %H:%M:%S
			UpdateTime           string `json:"update_time "`            // 更新时间，格式：%Y-%m-%d %H:%M:%S
			CreateTime           string `json:"create_time "`            // 创建时间，格式：%Y-%m-%d %H:%M:%S
		} `json:"list "`                            // 应用列表
		PageInfo PageInfoReturn `json:"page_info "` // 分页信息
	} `json:"data "`                      // 返回数据
	RequestId string `json:"request_id "` // 请求的日志id，唯一标识一个请求
}

// 查询应用信息
type ToolsAppManagementAppGetReturn struct {
	Code    int    `json:"code "`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message "` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		List []struct {
			PackageId   string `json:"package_id "`   // 应用包ID
			PackageName string `json:"package_name "` // 包名
			AppCloudId  int    `json:"app_cloud_id "` // app id
			AppName     string `json:"app_name "`     // 应用名
			Version     string `json:"version "`      // 版本号
			DownloadUrl string `json:"download_url "` // 下载地址
			IconUrl     string `json:"icon_url "`     // icon地址
			PublishTime string `json:"publish_time "` // 发布时间，格式：%Y-%m-%d %H:%M:%S
			UpdateTime  string `json:"update_time "`  // 更新时间，格式：%Y-%m-%d %H:%M:%S
			CreateTime  string `json:"create_time "`  // 创建时间，格式：%Y-%m-%d %H:%M:%S
		} `json:"list "`                            // 应用列表
		PageInfo PageInfoReturn `json:"page_info "` // 分页信息
	} `json:"data "`                      // 返回数据
	RequestId string `json:"request_id "` // 请求的日志id，唯一标识一个请求
}

// 查询应用预约记录
type ToolsAppManagementBookingRecordsGetReturn struct {
	Code    int    `json:"code "`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message "` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		List []struct {
			AdId       int    `json:"ad_id "`       // 广告计划ID
			CreativeId int    `json:"creative_id "` // 广告创意ID
			OrderId    string `json:"order_id "`    // 预约ID
			PackageId  string `json:"package_id "`  // 应用ID
			ReqId      string `json:"req_id "`      // 请求ID
			CreateTime string `json:"create_time "` // 创建时间，格式：%Y-%m-%d %H:%M:%S
		} `json:"list "`                            // 预约记录数据
		PageInfo PageInfoReturn `json:"page_info "` // 分页信息
	} `json:"data "`                      // 返回数据
	RequestId string `json:"request_id "` // 请求的日志id，唯一标识一个请求
}

// 提交解析应用包任务
type ToolsDownloadPackageParseReturn struct {
	Code      int    `json:"code"`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message   string `json:"message"`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	RequestId string `json:"request_id"` // 请求日志id
	Data      struct {
		EventId string `json:"event_id"` // 事件ID，用于查询解析包状态接口，有效期为24小时
	} `json:"data"` // 返回数据
}

// 查询包解析状态
type ToolsDownloadPackageGetReturn struct {
	Code      int    `json:"code"`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message   string `json:"message"`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	RequestId string `json:"request_id"` // 请求日志id
	Data      struct {
		PackageStatus string `json:"package_status"` // 包解析状态，显示当前包解析状态，详见[【附件-包解析状态】](https://ad.oceanengine.com/openapi/doc/index.html?id=528#custom-anchor-包解析状态)
	} `json:"data"` // 返回数据
}

// 查询应用分包列表
type ToolsAppManagementExtendPackageListReturn struct {
	Code      int    `json:"code"`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message   string `json:"message"`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	RequestId string `json:"request_id"` // 请求日志id
	Data      struct {
		List []struct {
			Reason      string `json:"reason"`       // 分包失败原因
			Status      string `json:"status"`       // 状态
			UpdateTime  string `json:"update_time"`  // 更新时间
			VersionName string `json:"version_name"` // 版本号
			ChannelId   string `json:"channel_id"`   // 渠道号
			DownloadUrl string `json:"download_url"` // 下载链接
			PackageId   string `json:"package_id"`   // 应用包ID
		} `json:"list"`                            // 应用分包列表
		PageInfo PageInfoReturn `json:"page_info"` // 页面信息
	} `json:"data"`                       // 返回数据
	TotalNumber int `json:"total_number"` // 总数
}
