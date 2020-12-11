/* #
# Author: wangchunxin
# Created Time: 2020/11/21 2:36 下午
# File Name:
# Description:
# */
package model

type CommonReturn struct {
	Code      int    `json:"code"`       // 返回码
	Message   string `json:"message"`    // 返回信息
	RequestId string `json:"request_id"` // 请求日志ID
}

//PageInfo
type PageInfoReturn struct {
	TotalNumber int `json:"total_number"`
	Page        int `json:"page"`
	PageSize    int `json:"page_size"`
	TotalPage   int `json:"total_page"`
}
