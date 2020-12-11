/* #
# Author: wangchunxin
# Created Time: 2020/11/26 1:53 下午
# File Name:
# Description:
# */
package config

type Config struct {
	BaseUrl      string
	ApiVersion   string
	BasePath     string
	CommonParams CommonParams
}

type CommonParams struct {
	Page     int
	PageSize int
}

func NewConfig() *Config {
	return &Config{
		BaseUrl:    BASE_URL,
		ApiVersion: VERSION,
		BasePath:   BASE_URL + VERSION,
		CommonParams:CommonParams{
			Page:     1,
			PageSize: 10,
		},
	}
}
