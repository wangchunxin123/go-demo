/* #
# Author: wangchunxin
# Created Time: 2020/11/20 5:18 下午
# File Name:
# Description:
# */
package model

type AuthGetUrlRequest struct {
	AppId        int64  `json:"app_id,omitempty"` // app_id
	State        string `json:"state,omitempty"`  // 自定义参数
	Scope        []int  `json:"scope,omitempty"`  // 权限范围
	RedirectUri  string `json:"redirect_uri,omitempty"` // 回调地址
	MaterialAuth int    `json:"material_auth,omitempty"` //
}
