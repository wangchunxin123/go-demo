/* #
# Author: wangchunxin
# Created Time: 2020/12/3 9:38 上午
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

type CreativeApiService Service

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 获取创意列表
// @Date 11:08 上午 2020/12/3
// @Param
// @return

func (c *CreativeApiService) Get(ctx context.Context, data CreativeGetRequest) (*CreativeGetReturn, error) {
	var creativeGetReturn *CreativeGetReturn

	path := c.Cfg.BasePath + config.CREATIVE_GET_URI
	header := map[string]string{
		"Access-Token": ctx.Value("access_token").(string),
	}
	if data.AdvertiserId == 0 {
		return creativeGetReturn, errors.New("缺少账户id")
	}
	if data.Page == 0 {
		data.Page = 1
	}
	if data.PageSize == 0 {
		data.PageSize = 10
	}
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_GET, path, dataByte, header, &creativeGetReturn); err != nil {
		return creativeGetReturn, err
	}
	if creativeGetReturn.Code == 0 {
		return creativeGetReturn, nil
	} else {
		return creativeGetReturn, errors.New(creativeGetReturn.Message)
	}
}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 创意详情
// @Date 11:29 上午 2020/12/3
// @Param
// @return

func (c *CreativeApiService) Read(ctx context.Context, data CreativeReadRequest) (*CreativeReadReturn, error) {
	var creativeReadReturn *CreativeReadReturn

	path := c.Cfg.BasePath + config.CREATIVE_READ_V2_URI
	header := map[string]string{
		"Access-Token": ctx.Value("access_token").(string),
	}
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_GET, path, dataByte, header, &creativeReadReturn); err != nil {
		return creativeReadReturn, err
	}
	if creativeReadReturn.Code == 0 {
		return creativeReadReturn, nil
	} else {
		return creativeReadReturn, errors.New(creativeReadReturn.Message)
	}

}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 更新创意状态
// @Date 1:53 下午 2020/12/3
// @Param
// @return

func (c *CreativeApiService) UpdateStatus(ctx context.Context, data CreativeUpdateStatusRequest) (*CreativeUpdateStatusReturn, error) {
	var creativeUpdateStatusReturn *CreativeUpdateStatusReturn
	path := c.Cfg.BasePath + config.CREATIVE_UPDATE_STATUS_URI
	header := map[string]string{
		"Content-Type": "application/json",
		"Access-Token": ctx.Value("access_token").(string),
	}
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_POST, path, dataByte, header, &creativeUpdateStatusReturn); err != nil {
		return creativeUpdateStatusReturn, err
	}
	if creativeUpdateStatusReturn.Code == 0 {
		return creativeUpdateStatusReturn, nil
	} else {
		return creativeUpdateStatusReturn, errors.New(creativeUpdateStatusReturn.Message)
	}
}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 创意素材信息
// @Date 2:14 下午 2020/12/3
// @Param
// @return

func (c *CreativeApiService) MaterialRead(ctx context.Context, data CreativeMaterialReadRequest) (*CreativeMaterialReadReturn, error) {
	var creativeMaterialReadReturn *CreativeMaterialReadReturn
	path := c.Cfg.BasePath + config.CREATIVE_MATERIAL_READ_URI
	header := map[string]string{
		"Access-Token": ctx.Value("access_token").(string),
	}
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_GET, path, dataByte, header, &creativeMaterialReadReturn); err != nil {
		return creativeMaterialReadReturn, err
	}
	if creativeMaterialReadReturn.Code == 0 {
		return creativeMaterialReadReturn, nil
	} else {
		return creativeMaterialReadReturn, errors.New(creativeMaterialReadReturn.Message)
	}
}

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 获取创意审核建议
// @Date 2:32 下午 2020/12/3
// @Param
// @return

func (c *CreativeApiService) RejectReason(ctx context.Context, data CreativeRejectReasonRequest) (*CreativeRejectReasonReturn ,error) {
	var creativeRejectReasonReturn *CreativeRejectReasonReturn
	path := c.Cfg.BasePath + config.CREATIVE_REJECT_REASON_URI
	header := map[string]string{
		"Access-Token": ctx.Value("access_token").(string),
	}
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_GET, path, dataByte, header, &creativeRejectReasonReturn); err != nil {
		return creativeRejectReasonReturn, err
	}
	if creativeRejectReasonReturn.Code == 0 {
		return creativeRejectReasonReturn, nil
	} else {
		return creativeRejectReasonReturn, errors.New(creativeRejectReasonReturn.Message)
	}
}
