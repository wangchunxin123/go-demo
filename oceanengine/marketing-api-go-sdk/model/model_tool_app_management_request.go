/* #
# Author: wangchunxin
# Created Time: 2020/12/8 2:46 下午
# File Name:
# Description:
# */
package model

// 查询游戏信息
type ToolsAppManagementBookingGetRequest struct {
	AdvertiserId int    `json:"advertiser_id ,omitempty"` // 广告主ID
	SearchKey    string `json:"search_key ,omitempty"`    // 搜索关键字appid或者应用名，可以为空，可以传中文长度不超过50
	SearchType   string `json:"search_type ,omitempty"`   // 搜索类型: `ALL`:查询全部，包括创建和被共享的应用（默认）`CREATE_ONLY`:只查询广告主创建的应用`SHARED_ONLY`:只查询被共享的应用
	Status       string `json:"status ,omitempty"`        //  应用状态:`ALL`:所有状态`AUDIT_DOING`:审核中`AUDIT_REJECTED`:审核失败`BOOKING`:预约中`ENABLE`:已发布（默认）`PAST_DUE`:已逾期
	Page         int    `json:"page ,omitempty"`          // 页码，默认值为1
	PageSize     int    `json:"page_size ,omitempty"`     // 页面大小，默认值为10，最大不超过200
	CreateTime   struct {
		StartTime string `json:"start_time ,omitempty"` // 创建起始时间，默认为空，格式：%Y-%m-%d
		EndTime   string `json:"end_time ,omitempty"`   // 创建结束时间，默认为当天，格式：%Y-%m-%d
	} `json:"create_time ,omitempty"` // 按创建时间查询的时间范围
	ScheduledPublishTime struct {
		StartTime string `json:"start_time ,omitempty"` // 预约发布起始时间，默认为空，格式：%Y-%m-%d
		EndTime   string `json:"end_time ,omitempty"`   // 预约发布结束时间，默认为当天，格式：%Y-%m-%d
	} `json:"scheduled_publish_time ,omitempty"` // 按预约发布时间查询的时间范围
}

// 查询应用信息
type ToolsAppManagementAppGetRequest struct {
	AdvertiserId int    `json:"advertiser_id ,omitempty"` // 广告主ID
	SearchKey    string `json:"search_key ,omitempty"`    // 搜索关键字appid或者应用名，可以为空，可以传中文长度不超过50
	SearchType   string `json:"search_type ,omitempty"`   // 搜索类型: `ALL`:查询全部，包括创建和被共享的应用（默认）`CREATE_ONLY`:只查询广告主创建的应用`SHARED_ONLY`:只查询被共享的应用
	Status       string `json:"status ,omitempty"`        //  应用状态:`ALL`:所有状态`AUDIT_DOING`:审核中`AUDIT_REJECTED`:审核失败`AUDIT_ACCEPTED`:审核成功，即待发布`ENABLE`:已发布（默认）
	Page         int    `json:"page ,omitempty"`          // 页码，默认值为1
	PageSize     int    `json:"page_size ,omitempty"`     // 页面大小，默认值为10，最大不超过200
	CreateTime   struct {
		StartTime string `json:"start_time ,omitempty"` // 创建起始时间，默认为空，格式：%Y-%m-%d
		EndTime   string `json:"end_time ,omitempty"`   // 创建结束时间，默认为当天，格式：%Y-%m-%d
	} `json:"create_time ,omitempty"` // 按创建时间查询的时间范围
	PublishTime struct {
		StartTime string `json:"start_time ,omitempty"` // 发布起始时间，默认为空，格式：%Y-%m-%d
		EndTime   string `json:"end_time ,omitempty"`   // 发布结束时间，默认为当天，格式：%Y-%m-%d
	} `json:"publish_time ,omitempty"` // 按发布时间查询的时间范围
}

// 查询应用预约记录
type ToolsAppManagementBookingRecordsGetRequest struct {
	AdvertiserId int    `json:"advertiser_id ,omitempty"` // 广告主ID
	PackageId    string `json:"package_id ,omitempty"`    // 应用id如果是Anroid应用，package_id从【查询应用信息】接口获取，接口预计下周上线；如果是IOS应用，package_id为App Store中的ItunesID；
	HostType     string `json:"host_type ,omitempty"`     //  应用是否寄存在应用管理中: `HOST_IN`:寄存`HOST_OUT`:未寄存如果是Anroid应用，必须先上传至应用管理中，请选择寄存；如果是IOS应用,请选择未寄存;
	Page         int    `json:"page ,omitempty"`          // 页码，默认值为1
	PageSize     int    `json:"page_size ,omitempty"`     // 页面大小，默认值为10，最大不超过200
	CreateTime   struct {
		StartTime string `json:"start_time ,omitempty"` // 创建起始时间，默认为空，格式：%Y-%m-%d
		EndTime   string `json:"end_time ,omitempty"`   // 创建结束时间，默认为当天，格式：%Y-%m-%d
	} `json:"create_time ,omitempty"` // 按创建时间查询的时间范围
}

// 提交解析应用包任务
type ToolsDownloadPackageParseRequest struct {
	AdvertiserId int    `json:"advertiser_id ,omitempty"` // 广告主ID
	DownloadUrl  string `json:"download_url ,omitempty"`  // 下载链接
}

// 查询包解析状态
type ToolsDownloadPackageGetRequest struct {
	AdvertiserId string `json:"advertiser_id ,omitempty"` // 广告主ID
	EventId      string `json:"event_id ,omitempty"`      // 事件ID,由[【提交解析应用包任务】](https://ad.oceanengine.com/openapi/doc/index.html?id=1681127683816455)获取
}

// 查询应用分包列表
type ToolsAppManagementExtendPackageListRequest struct {
	PackageId    string `json:"package_id ,omitempty"`    // 应用包ID，获取方法见接口文档[【查询应用信息】](https://ad.oceanengine.com/openapi/doc/index.html?id=1670889966726148)
	AdvertiserId int    `json:"advertiser_id ,omitempty"` // 广告主ID
	Filtering    struct {
		Status string `json:"status,omitempty"` // 状态允许值：`"ALL"`：全部`"NOT_UPDATE"`： 未更新`"CREATING"`：创建中`"UPDATING"`：更新中`"PUBLISHED"`：已发布`"CREATION_FAILED"`：创建失败`"UPDATE_FAILED"`：更新失败默认值：`"ALL"`
	} `json:"filtering,omitempty"`            // 过滤条件
	Page     int `json:"page,omitempty"`      // 页数默认值：`1`
	PageSize int `json:"page_size,omitempty"` // 页面大小默认值：`10`，范围：`1～1000`
}
