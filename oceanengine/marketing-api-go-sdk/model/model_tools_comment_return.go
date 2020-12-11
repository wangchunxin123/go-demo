/* #
# Author: wangchunxin
# Created Time: 2020/12/8 12:00 下午
# File Name:
# Description:
# */
package model

//工具--查询评论
type ToolsCommentGetReturn struct {
	CommonReturn
	Data struct {
		CommentsList []struct {
			Id          int64       `json:"id"`
			AppName     string      `json:"app_name"`
			UserInfo    interface{} `json:"user_info"`
			AdId        int64       `json:"ad_id"`
			AdName      string      `json:"ad_name"`
			Text        string      `json:"text"`
			LikeCount   int         `json:"like_count"`
			CreateTime  string      `json:"create_time"`
			Stick       int         `json:"stick"`
			ReplyCount  int         `json:"reply_count"`
			InBlackList int         `json:"in_black_list"`
		} `json:"comments_list"`
		PageInfo PageInfoReturn `json:"page_info"`
	}
}


