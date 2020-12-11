/* #
# Author: wangchunxin
# Created Time: 2020/11/23 10:18 上午
# File Name:
# Description:
# */
package api

import (
	"context"
	"encoding/json"
	"errors"
	"git.hsuanyuen.cn/oceanengine/marketing-api-go-sdk/config"
	. "git.hsuanyuen.cn/oceanengine/marketing-api-go-sdk/model"
)

type CampaignApiService Service

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 获取组列表
// @Date 6:06 下午 2020/12/1
// @Param
// @return

func (c *CampaignApiService) Get(ctx context.Context, data CampaignGetRequest) (CampaignGetReturn, error) {
	var campaignReturn CampaignGetReturn
	path := c.Cfg.BasePath + config.CAMPAIGN_GET_URI
	header := map[string]string{
		"Access-Token": ctx.Value("access_token").(string),
	}
	if data.AdvertiserId == 0 {
		return campaignReturn, errors.New("缺少账户id")
	}
	if data.Page == 0 {
		data.Page = 1
	}
	if data.PageSize == 0 {
		data.PageSize = 10
	}
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_GET, path, dataByte, header, &campaignReturn); err != nil {
		return campaignReturn, err
	}
	if campaignReturn.Code == 0 {
		return campaignReturn, nil
	} else {
		return campaignReturn, errors.New(campaignReturn.Message)
	}
}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 创建组
// @Date 6:07 下午 2020/12/1
// @Param
// @return

func (c *CampaignApiService) Add(ctx context.Context, data CampaignAddRequest) (CampaignAddReturn, error) {
	var campaignAddReturn CampaignAddReturn

	path := c.Cfg.BasePath + config.CAMPAIGN_CREATE_URI

	header := map[string]string{
		"Content-Type": "application/json",
		"Access-Token": ctx.Value("access_token").(string),
	}
	if data.AdvertiserId == 0 {
		return campaignAddReturn, errors.New("缺少账户id")
	}
	if data.CampaignName == "" {
		return campaignAddReturn, errors.New("组名称为空")
	}
	if data.BudgetMode == "" {
		return campaignAddReturn, errors.New("请填写广告组预算类型")
	}
	if data.BudgetMode == "BUDGET_MODE_DAY" && data.Budget < 300 {
		return campaignAddReturn, errors.New("请填写广告组预算类型")
	}

	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_POST, path, dataByte, header, &campaignAddReturn); err != nil {
		return campaignAddReturn, err
	}
	if campaignAddReturn.Code == 0 {
		return campaignAddReturn, nil
	}
	return campaignAddReturn, errors.New(campaignAddReturn.Message)
}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 修改组
// @Date 3:09 下午 2020/12/2
// @Param
// @return

func (c *CampaignApiService) Update(ctx context.Context, data CampaignUpdateRequest) (CampaignUpdateReturn, error) {
	var campaignUpdateReturn CampaignUpdateReturn
	header := map[string]string{
		"Content-Type": "application/json",
		"Access-Token": ctx.Value("access_token").(string),
	}
	path := c.Cfg.BasePath + config.CAMPAIGN_UPDATE_URI
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_POST, path, dataByte, header, &campaignUpdateReturn); err != nil {
		return campaignUpdateReturn, err
	}
	if campaignUpdateReturn.Code == 0 {
		return campaignUpdateReturn, nil
	}
	return campaignUpdateReturn, errors.New(campaignUpdateReturn.Message)
}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 更新状态
// @Date 3:18 下午 2020/12/2
// @Param
// @return

func (c *CampaignApiService) UpdateStatus(ctx context.Context, data CampaignUpdateStatusRequest) (CampaignUpdateStatusReturn, error) {
	var campaignUpdateStatusReturn CampaignUpdateStatusReturn
	header := map[string]string{
		"Content-Type": "application/json",
		"Access-Token": ctx.Value("access_token").(string),
	}
	path := c.Cfg.BasePath + config.CAMPAIGN_UPDATE_STATUS_URI
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_POST, path, dataByte, header, &campaignUpdateStatusReturn); err != nil {
		return campaignUpdateStatusReturn, err
	}
	if campaignUpdateStatusReturn.Code == 0 {
		return campaignUpdateStatusReturn, nil
	}
	return campaignUpdateStatusReturn, errors.New(campaignUpdateStatusReturn.Message)
}
