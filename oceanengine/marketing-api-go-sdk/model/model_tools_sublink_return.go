/* #
# Author: wangchunxin
# Created Time: 2020/12/10 3:39 下午
# File Name:
# Description:
# */
package model

// 获取子链列表
type ToolsSublinkListReturn struct {
	Code      int    `json:"code"`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/athena/docs/manage.html?plat_id=1)
	Message   string `json:"message"`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/athena/docs/manage.html?plat_id=1)
	RequestId string `json:"request_id"` // 请求日志id
	Data      struct {
		List []struct {
			Tag         string `json:"tag"`           // 子链标签，仅标签长子链有值
			AuditStatus string `json:"audit_status"`  // 子链状态: `AUDIT` 待审核, `ACCEPT` 审核通过, `REJECT` 审核不通过
			BindAdCount int64  `json:"bind_ad_count"` // 关联计划数
			Id          int64  `json:"id"`            // 子链id
			Info        string `json:"info"`          // 子链内容
			RejectInfo  string `json:"reject_info"`   // 审核拒绝理由
			SubLinkType string `json:"sub_link_type"` // 子链类型: `SHORT` 短子链, `LONG` 长子链, `TAG_LONG` 标签长子链
		} `json:"list"` // 子链列表
		PageInfo PageInfoReturn `json:"page_info"`
	} `json:"data"` // 返回数据
}

// 创建子链
type ToolsSublinkCreateReturn struct {
	Code      int    `json:"code"`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/athena/docs/manage.html?plat_id=1)
	Message   string `json:"message"`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/athena/docs/manage.html?plat_id=1)
	RequestId string `json:"request_id"` // 请求日志id
	Data      struct {
		SubLinks []struct {
			Tag          string `json:"tag"`           // 子链标签
			AdvertiserId int64  `json:"advertiser_id"` // 广告主id
			AuditStatus  string `json:"audit_status"`  // 子链状态: `AUDIT` 待审核, `ACCEPT` 审核通过, `REJECT` 审核不通过
			Id           int64  `json:"id"`            // 子链id
			Info         string `json:"info"`          // 子链内容
			SubLinkType  string `json:"sub_link_type"` // 子链类型: `SHORT` 短子链, `LONG` 长子链, `TAG_LONG` 标签长子链
		} `json:"sub_links"` // 新建的子链信息
	} `json:"data"` // 返回数据
}

// 编辑子链
type ToolsSublinkUpdateReturn struct {
	Code      int    `json:"code"`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/athena/docs/manage.html?plat_id=1)
	Message   string `json:"message"`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/athena/docs/manage.html?plat_id=1)
	RequestId string `json:"request_id"` // 请求日志id
	Data      struct {
		SubLinks []struct {
			Id           int64  `json:"id"`            // 子链id
			Info         string `json:"info"`          // 子链内容
			SubLinkType  string `json:"sub_link_type"` // 子链类型: `SHORT` 短子链, `LONG` 长子链, `TAG_LONG` 标签长子链
			Tag          string `json:"tag"`           // 子链标签
			AdvertiserId int64  `json:"advertiser_id"` // 广告主id
			AuditStatus  string `json:"audit_status"`  // 子链状态: `AUDIT` 待审核, `ACCEPT` 审核通过, `REJECT` 审核不通过

		} `json:"sub_links"` // 编辑后的子链信息
	} `json:"data"` // 返回数据
}

// 删除子链
type ToolsSublinkDeleteReturn struct {
	Code      int    `json:"code"`       // 返回码,详见[【附录-返回码】](https://ad.oceanengine.com/athena/docs/manage.html?plat_id=1)
	Message   string `json:"message"`    // 返回信息,详见[【附录-返回码】](https://ad.oceanengine.com/athena/docs/manage.html?plat_id=1)
	RequestId string `json:"request_id"` // 请求日志id
	Data      struct {
		FailSubLinkIds []struct {
			FailMsg string `json:"fail_msg"` // 失败原因
			Id      int64  `json:"id"`       // 子链id
		} `json:"fail_sub_link_ids"`                      // 删除失败的子链id列表
		SuccSubLinkIds []int64 `json:"succ_sub_link_ids"` // 删除成功的子链id列表
	} `json:"data"` // 返回数据
}
