/* #
# Author: wangchunxin
# Created Time: 2020/11/20 5:03 下午
# File Name:
# Description:
# */
package main

import (
	"fmt"
	"git.hsuanyuen.cn/oceanengine/marketing-api-go-sdk/common"
	. "git.hsuanyuen.cn/oceanengine/marketing-api-go-sdk/model"
)

func main() {
	accessToken := "07cda89b9dcda9ecf0b5015c3dc5808bc6873920"
	sdk := common.Init(accessToken)
	ctx := *sdk.Content
	data := PlanGetRequest{
		AdvertiserId:66424737208,
	}
	//data := CampaignAddRequest{
	//	AdvertiserId:        66424737208,
	//	CampaignName:        "SDK-test",
	//	Operation:           "",
	//	BudgetMode:          "BUDGET_MODE_INFINITE",
	//	Budget:              0,
	//	LandingType:         "LINK",
	//	DeliveryRelationNum: "",
	//}
	resdata, err := sdk.Plan().Get(ctx, data)
	fmt.Println("resdata:",resdata)
	fmt.Println("err:",err)
}
