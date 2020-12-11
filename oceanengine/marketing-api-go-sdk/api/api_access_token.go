/* #
# Author: wangchunxin
# Created Time: 2020/11/21 2:31 下午
# File Name:
# Description:
# */
package api

import (
	"context"
	"fmt"
	"git.hsuanyuen.cn/oceanengine/marketing-api-go-sdk/model"
)

type AccessTokenApiService Service

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 获取token
// @Date 2:48 下午 2020/11/21
// @Param
// @return

func (a AccessTokenApiService) Get(ctx context.Context, params *model.AccessTokenRequest) {
	fmt.Println("111")
}
