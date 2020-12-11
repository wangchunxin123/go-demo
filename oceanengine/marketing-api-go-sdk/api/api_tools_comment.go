/* #
# Author: wangchunxin
# Created Time: 2020/12/10 10:18 上午
# File Name:
# Description:
# */
package api

import (
	"context"
	. "git.hsuanyuen.cn/oceanengine/marketing-api-go-sdk/model"
)

type ToolsComment Service

// @Author wangchunxin@hsuanyuen.com
// @Description //TODO 查询评论
// @Date 10:18 上午 2020/12/10
// @Param
// @return
func (t *ToolsComment)Get(ctx context.Context, data ToolsCommentGetRequest) {
	//var toolsDiagnosisAdGetReturn *ToolsDiagnosisAdGetReturn
	//path := t.Cfg.BasePath + config.TOOLS_COMMENT_GET_URI
	//header := map[string]string{
	//	"Access-Token": ctx.Value("access_token").(string),
	//}
	//dataByte, _ := json.Marshal(data)
	//
}
