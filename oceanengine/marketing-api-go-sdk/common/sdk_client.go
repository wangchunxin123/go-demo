/* #
# Author: wangchunxin
# Created Time: 2020/11/21 3:09 下午
# File Name:
# Description:
# */
package common

import (
	"context"
	"git.hsuanyuen.cn/oceanengine/marketing-api-go-sdk/api"
	"git.hsuanyuen.cn/oceanengine/marketing-api-go-sdk/config"
)

type SDKClient struct {
	Cfg        *config.Config
	Client     *api.APIClient
	Content    *context.Context
}

func Init(accessToken string) *SDKClient {
	cfg := config.NewConfig()
	client := api.NewAPIClient(cfg)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "access_token", accessToken)

	return &SDKClient{
		Cfg:        cfg,
		Client:     client,
		Content:    &ctx,
	}
}

// access-token
func (s *SDKClient) AccessToken() *api.AccessTokenApiService {
	return s.Client.AccessTokenApi
}


// 组
func (s *SDKClient) Campaign() *api.CampaignApiService {
	return  s.Client.CampaignApi
}

// 计划
func (s *SDKClient) Plan() *api.PlanApiService {
	return  s.Client.PlanApi
}
