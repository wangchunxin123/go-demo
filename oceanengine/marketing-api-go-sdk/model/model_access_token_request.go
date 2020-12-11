/* #
# Author: wangchunxin
# Created Time: 2020/11/20 5:13 下午
# File Name:
# Description:
# */
package model

type AccessTokenRequest struct {
	AppId     int64  `json:"app_id"`     // 开发者申请的应用APP_ID
	Secret    string `json:"secret"`     // 开发者应用的私钥Secret
	GrantType string `json:"grant_type"` // 授权类型。允许值: "auth_code"
	AuthCode  string `json:"auth_code"`  // 授权码
}
