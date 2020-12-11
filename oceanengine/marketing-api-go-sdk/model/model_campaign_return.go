/* #
# Author: wangchunxin
# Created Time: 2020/11/23 10:19 上午
# File Name:
# Description:
# */
package model

//获取组的返回信息
type CampaignGetReturn struct {
	CommonReturn
	Data struct {
		List     CampaignGetListReturn `json:"list"`
		PageInfo PageInfoReturn        `json:"page_info"`
	} `json:"data"`
}

type CampaignGetListReturn struct {
	Status             string      `json:"status"`               // 广告组状态
	BudgetMode         string      `json:"budget_mode"`          // 广告预算目的
	ModifyTime         string      `json:"modify_time"`          // 广告组时间戳,用于更新时提交,服务端判断是否基于最新信息修改
	DeliveryRelatedNum string      `json:"delivery_related_num"` // 广告组商品数量
	CampaignModifyTime string      `json:"campaign_modify_time"` //
	DeliveryMode       string      `json:"delivery_mode"`        //投放类型
	CampaignID         string      `json:"campaign_id"`          // 组ID
	Budget             int         `json:"budget"`               // 预算
	AdvertiserID       int64       `json:"advertiser_id"`        // 账户ID
	LandingType        string      `json:"landing_type"`         // 广告组推广目的
	MarketingScene     string      `json:"marketing_scene"`
	CampaignType       string      `json:"campaign_type"`        // 广告组类型
	UniqueFk           interface{} `json:"unique_fk"`            // （废弃）第三方唯一键，传该值时保证接口重试的幂等性
	CampaignCreateTime string      `json:"campaign_create_time"` // 组创建时间
	ID                 string      `json:"id"`                   //广告组ID
	Name               string      `json:"name"`                 // 广告组名称
}

//添加返回
type CampaignAddReturn struct {
	CommonReturn
	Data struct{
		CampaignId int64 `json:"campaign_id"`
	}
}

//修改返回
type CampaignUpdateReturn struct {
	CommonReturn
	Data struct{
		CampaignId int64 `json:"campaign_id"`
	}
}

//修改状态返回
type CampaignUpdateStatusReturn struct {
	CommonReturn
	Data struct{
		CampaignIds []int64 `json:"campaign_ids"`
	}
}