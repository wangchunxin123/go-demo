/* #
# Author: wangchunxin
# Created Time: 2020/12/8 5:13 下午
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

type ToolsDiagnosisAd Service

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 获取计划诊断详情
// @Date 6:02 下午 2020/12/8
// @Param 
// @return
 
func (t *ToolsDiagnosisAd) Get(ctx context.Context, data ToolsDiagnosisAdGetRequest) (*ToolsDiagnosisAdGetReturn, error) {
	var toolsDiagnosisAdGetReturn *ToolsDiagnosisAdGetReturn
	path := t.Cfg.BasePath + config.TOOLS_DIAGNOSIS_AD_GET_URI
	header := map[string]string{
		"Access-Token": ctx.Value("access_token").(string),
	}
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_GET, path, dataByte, header, &toolsDiagnosisAdGetReturn); err != nil {
		return toolsDiagnosisAdGetReturn, err
	}
	if toolsDiagnosisAdGetReturn.Code == 0 {
		return toolsDiagnosisAdGetReturn, nil
	} else {
		return toolsDiagnosisAdGetReturn, errors.New(toolsDiagnosisAdGetReturn.Message)
	}
}


// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 获取计划诊断预估变化趋势
// @Date 6:03 下午 2020/12/8
// @Param 
// @return
 
func (t *ToolsDiagnosisAd) Curve(ctx context.Context, data ToolsDiagnosisAdCurveRequest) (*ToolsDiagnosisAdCurveReturn, error) {
	var toolsDiagnosisAdCurveReturn *ToolsDiagnosisAdCurveReturn
	path := t.Cfg.BasePath + config.TOOLS_DIAGNOSIS_AD_CURVE_URI
	header := map[string]string{
		"Access-Token": ctx.Value("access_token").(string),
	}
	dataByte, _ := json.Marshal(data)
	if err := ApiRequest(MOTHOD_GET, path, dataByte, header, &toolsDiagnosisAdCurveReturn); err != nil {
		return toolsDiagnosisAdCurveReturn, err
	}
	if toolsDiagnosisAdCurveReturn.Code == 0 {
		return toolsDiagnosisAdCurveReturn, nil
	} else {
		return toolsDiagnosisAdCurveReturn, errors.New(toolsDiagnosisAdCurveReturn.Message)
	}
}