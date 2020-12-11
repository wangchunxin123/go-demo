/* #
# Author: wangchunxin
# Created Time: 2020/11/21 2:49 下午
# File Name:
# Description:
# */
package api

import (
	"git.hsuanyuen.cn/oceanengine/marketing-api-go-sdk/config"
)

type APIClient struct {
	Cfg            *config.Config
	AccessTokenApi *AccessTokenApiService
	CampaignApi    *CampaignApiService
	PlanApi        *PlanApiService
	common         Service
}

type Service struct {
	Cfg *config.Config
}

func NewAPIClient(cfg *config.Config) *APIClient {
	c := &APIClient{}
	c.Cfg = cfg
	c.common.Cfg = cfg

	c.CampaignApi = (*CampaignApiService)(&c.common)
	c.PlanApi = (*PlanApiService)(&c.common)
	c.AccessTokenApi = &AccessTokenApiService{}
	//c.CampaignApi =&CampaignApiService{}
	return c
}
