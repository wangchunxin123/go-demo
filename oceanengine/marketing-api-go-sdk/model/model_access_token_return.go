/* #
# Author: wangchunxin
# Created Time: 2020/11/21 2:24 下午
# File Name:
# Description:
# */
package model

type AccessTokenReturn struct {
	CommonReturn
	Data struct {
		AccessToken           string `json:"access_token"`             // token
		ExpiresIn             int    `json:"expires_in"`               // access_token剩余有效时间,单位(秒)
		RefreshToken          string `json:"refresh_token"`            // 刷新access_token,用于获取新的access_token和refresh_token，并且刷新过期时间
		RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"` // refresh_token剩余有效时间,单位(秒)
		AdvertiserID          int    `json:"advertiser_id"`            //（将废弃，当前支持代理商角色账号授权可一次授权多个账号，请使用advertiser_ids字段获取授权账号ID）登录用户对应的广告账户ID, 如果授权多个广告主默认为第一个
		AdvertiserIds         []int  `json:"advertiser_ids"`           //授权的账户id列表
	} `json:"data"`
}
