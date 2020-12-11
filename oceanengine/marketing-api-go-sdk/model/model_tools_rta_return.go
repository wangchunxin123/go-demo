/* #
# Author: wangchunxin
# Created Time: 2020/12/10 3:36 下午
# File Name:
# Description:
# */
package model

// 绑定RTA策略
type ToolsRtaBindReturn struct {
	Code      int         `json:"code "`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message   string      `json:"message "`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data      interface{} `json:"data "`       // 返回数据，空json
	RequestId string      `json:"request_id "` // 请求的日志id，唯一标识一个请求
}

// 解绑RTA策略
type ToolsRtaUnbindReturn struct {
	Code      int         `json:"code "`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message   string      `json:"message "`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data      interface{} `json:"data "`       // 返回数据，空json
	RequestId string      `json:"request_id "` // 请求的日志id，唯一标识一个请求
}

// 获取RTA策略数据
type ToolsRtaGetInfoReturn struct {
	Code    int    `json:"code "`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message "` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		InterfaceInfo struct {
			Status        int    `json:"status "`         // 接口地址状态1：生效 0：失效
			DeliveryRange string `json:"delivery_range "` //  适用流量范围: `LOCAL_ONLY`: 站内`UNION_ONLY`: 穿山甲 `UNIVERSAL_DELIVERY`: 全部
			LocalQps      int    `json:"local_qps "`      // 站内QPS
			UnionQps      int    `json:"union_qps "`      // 穿山甲QPS
			Url           string `json:"url "`            // 接口地址
		} `json:"interface_info "` // RTA配置数据
		RtaInfo struct {
			RtaId  int    `json:"rta_id "` // RTA策略ID
			Remark string `json:"remark "` // 备注，即RTA策略描述
		} `json:"rta_info "` // RTA策略信息
	} `json:"data "` // 返回数据

	RequestId string `json:"request_id "` // 请求的日志id，唯一标识一个请求
}

// 获取可用的RTA策略
type ToolsRtaGetReturn struct {
	Code    int    `json:"code "`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message "` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		List []struct {
			RtaId  int    `json:"rta_id "` // RTA策略ID
			Remark string `json:"remark "` // 备注，即RTA策略描述
		} `json:"list "` // 可用的RTA策略列表
	} `json:"data "`                      // 返回数据
	RequestId string `json:"request_id "` // 请求的日志id，唯一标识一个请求
}
