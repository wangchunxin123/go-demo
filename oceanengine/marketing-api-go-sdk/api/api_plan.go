/* #
# Author: wangchunxin
# Created Time: 2020/12/2 3:42 下午
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

type PlanApiService Service

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 获取计划列表
// @Date 4:44 下午 2020/12/2
// @Param
// @return

func (p *PlanApiService) Get(ctx context.Context, data PlanGetRequest) (*PlanGetReturn, error) {
	var planGetReturn *PlanGetReturn
	path := p.Cfg.BasePath + config.AD_GET_URI
	header := map[string]string{
		"Access-Token": ctx.Value("access_token").(string),
	}
	if data.AdvertiserId == 0 {
		return planGetReturn, errors.New("缺少账户id")
	}
	if data.Page == 0 {
		data.Page = 1
	}
	if data.PageSize == 0 {
		data.PageSize = 10
	}
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_GET, path, dataByte, header, &planGetReturn); err != nil {
		return planGetReturn, err
	}
	if planGetReturn.Code == 0 {
		return planGetReturn, nil
	} else {
		return planGetReturn, errors.New(planGetReturn.Message)
	}
}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 添加计划
// @Date 4:48 下午 2020/12/2
// @Param
// @return

func (p *PlanApiService) Add(ctx context.Context,  ) {

}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 修改计划
// @Date 4:48 下午 2020/12/2
// @Param 
// @return

func (p *PlanApiService) Update() {

}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 更新计划状态
// @Date 4:54 下午 2020/12/2
// @Param
// @return

func (p *PlanApiService) UpdateStatus(ctx context.Context, data PlanUpdateStatusReturn) (PlanUpdateStatusReturn, error) {
	var planUpdateStatusReturn PlanUpdateStatusReturn
	header := map[string]string{
		"Content-Type": "application/json",
		"Access-Token": ctx.Value("access_token").(string),
	}
	path := p.Cfg.BasePath + config.AD_UPDATE_STATUS_URI
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_POST, path, dataByte, header, &planUpdateStatusReturn); err != nil {
		return planUpdateStatusReturn, err
	}
	if planUpdateStatusReturn.Code == 0 {
		return planUpdateStatusReturn, nil
	}
	return planUpdateStatusReturn, errors.New(planUpdateStatusReturn.Message)
}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 更新计划预算
// @Date 6:08 下午 2020/12/2
// @Param 
// @return
 
func (p *PlanApiService) UpdateBudget(ctx context.Context, data PlanUpdateBudgetRequest) (PlanUpdateBudgetReturn, error) {
	var planUpdateBudgetReturn PlanUpdateBudgetReturn
	header := map[string]string{
		"Content-Type": "application/json",
		"Access-Token": ctx.Value("access_token").(string),
	}
	path := p.Cfg.BasePath + config.AD_UPDATE_BUDGET_URI
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_POST, path, dataByte, header, &planUpdateBudgetReturn); err != nil {
		return planUpdateBudgetReturn, err
	}
	if planUpdateBudgetReturn.Code == 0 {
		return planUpdateBudgetReturn, nil
	}
	return planUpdateBudgetReturn, errors.New(planUpdateBudgetReturn.Message)
}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 更新计划出价
// @Date 6:13 下午 2020/12/2
// @Param 
// @return 
 
func (p *PlanApiService) UpdateBid(ctx context.Context, data PlanUpdateBidRequest) (PlanUpdateBidReturn, error) {
	var planUpdateBidReturn PlanUpdateBidReturn
	header := map[string]string{
		"Content-Type": "application/json",
		"Access-Token": ctx.Value("access_token").(string),
	}
	path := p.Cfg.BasePath + config.AD_UPDATE_BID_URI
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_POST, path, dataByte, header, &planUpdateBidReturn); err != nil {
		return planUpdateBidReturn, err
	}
	if planUpdateBidReturn.Code == 0 {
		return planUpdateBidReturn, nil
	}
	return planUpdateBidReturn, errors.New(planUpdateBidReturn.Message)

}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 获取计划审核建议
// @Date 6:31 下午 2020/12/2
// @Param
// @return

func (p *PlanApiService) RejectReason(ctx context.Context, data PlanRejectReasonRequest) (PlanRejectReasonReturn, error){
	var planRejectReasonReturn PlanRejectReasonReturn
	header := map[string]string{
		"Access-Token": ctx.Value("access_token").(string),
	}
	path := p.Cfg.BasePath + config.AD_REJECT_REASON_URI
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_GET, path, dataByte, header, &planRejectReasonReturn); err != nil {
		return planRejectReasonReturn, err
	}
	if planRejectReasonReturn.Code == 0 {
		return planRejectReasonReturn, nil
	}
	return planRejectReasonReturn, errors.New(planRejectReasonReturn.Message)
}
