/* #
# Author: wangchunxin
# Created Time: 2020/12/10 2:17 下午
# File Name:
# Description:
# */
package model

// 获取穿山甲流量包
type ToolsUnionFlowPackageGetReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		PageInfo PageInfoReturn `json:"page_info"` // 分页信息
		List     []struct {
			FlowPackageId int    `json:"flow_package_id"` // 流量包ID
			Name          string `json:"name"`            // 流量包名称
			Status        string `json:"status"`          // 流量包状态，`"FLOW_PACKAGE_ENABLE"`：已启用，`"FLOW_PACKAGE_DISABLE"`：已废弃
			Rit           []int  `json:"rit"`             // 流量位ID数组
		} `json:"list"` // 流量包列表
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 创建穿山甲流量包
type ToolsUnionFlowPackageCreateReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		FlowPackageId int `json:"flow_package_id"` // 流量包ID
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 修改穿山甲流量包
type ToolsUnionFlowPackageUpdateReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		FlowPackageId int `json:"flow_package_id"` // 流量包ID
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 删除穿山甲流量包
type ToolsUnionFlowPackageDeleteReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		FlowPackageId int `json:"flow_package_id"` // 流量包ID，从[【获取穿山甲流量包】](https://ad.oceanengine.com/openapi/doc/index.html?id=384)获取
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}

// 查看rit数据
type ToolsUnionFlowPackageReportReturn struct {
	Code    int    `json:"code"`    // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Message string `json:"message"` // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/openapi/doc/index.html?id=529)
	Data    struct {
		PageInfo PageInfoReturn `json:"page_info"` // 分页信息
		List     []struct {
			Rit                int     `json:"rit"`                  // 广告位
			ActivePayRate      float64 `json:"active_pay_rate"`      // 付费率
			ActiveRegisterRate float64 `json:"active_register_rate"` // 注册率Form int `json:"form"`                               // 表单提交数
			ConversionCost     float64 `json:"conversion_cost"`      // 转化成本
			ActiveCost         float64 `json:"active_cost"`          // 激活成本
			GameAddictionCost  float64 `json:"game_addiction_cost"`  // 关键行为成本
			NextDayOpenCost    float64 `json:"next_day_open_cost"`   // 次留成本
			StatCost           float64 `json:"stat_cost"`            // 消耗，单位：元
			ActivePayCost      float64 `json:"active_pay_cost"`      // 付费成本
			NextDayOpenRate    float64 `json:"next_day_open_rate"`   // 次留率
			ActivePay          int     `json:"active_pay"`           // 付费数
			GameAddictionRate  float64 `json:"game_addiction_rate"`  // 关键行为率
			GameAddiction      int     `json:"game_addiction"`       // 关键行为数
			Active             int     `json:"active"`               // 激活数
			ActiveRegister     int     `json:"active_register"`      // 注册数
			ConvertCnt         int     `json:"convert_cnt"`          // 转化量
			ActiveRegisterCost int     `json:"active_register_cost"` // 注册成本
			NextDayOpen        int     `json:"next_day_open"`        // 次留数
			FormCost           int     `json:"form_cost"`            // 表单
			LoanCompletion     int     `json:"loan_completion"`      // 完件数
			PreLoanCredit      int     `json:"pre_loan_credit"`      // 预授信数
			LoanCredit         int     `json:"loan_credit"`          // 授信数
			LoanCompletionCost float64 `json:"loan_completion_cost"` // 完件成本
			PreLoanCreditCost  float64 `json:"pre_loan_credit_cost"` // 预授信成本
			LoanCreditCost     float64 `json:"loan_credit_cost"`     // 授信成本
		} `json:"list"` // 数据列表
	} `json:"data"`                      // json返回值
	RequestId string `json:"request_id"` // 请求日志id
}
