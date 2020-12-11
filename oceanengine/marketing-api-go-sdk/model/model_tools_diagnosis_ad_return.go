/* #
# Author: wangchunxin
# Created Time: 2020/12/7 6:20 下午
# File Name:
# Description:
# */
package model

type ToolsDiagnosisAdGetReturn struct {
	CommonReturn
	Data struct {
		List []struct {
			Severity      float64 `json:"severity"`
			FutureStarTag bool    `json:"future_star_tag"`
			SceneTag      string  `json:"scene_tag"`
			Suggest       string  `json:"suggest"`
			TargetQuality string  `json:"target_quality"`
			RitResult     struct {
				Severity     float64 `json:"severity"`
				Conclusion   string  `json:"conclusion"`
				FunnelData   string  `json:"funnel_data"`
				Attributions string  `json:"attributions"`
			} `json:"rit_result"`
			TimeRange    string        `json:"time_range"`
			SuggestValue float64       `json:"suggest_value"`
			Reason       string        `json:"reason"`
			ColdStart    float64       `json:"cold_start"`
			ModifyOption string        `json:"modify_option"`
			OriginValue  float64       `json:"origin_value"`
			Attributions []interface{} `json:"attributions"`
			FunnelData   struct {
				RateList []float64 `json:"rate_list"`
				BidList  []float64 `json:"bid_list"`
				CostList []float64 `json:"cost_list"`
			} `json:"funnel_data"`
			Issue      string `json:"issue"`
			ID         int64  `json:"id"`
			Conclusion string `json:"conclusion"`
		} `json:"list"`
	} `json:"data"`
}


//获取计划诊断预估变化趋势
type ToolsDiagnosisAdCurveReturn struct {
	CommonReturn
	Data struct {
		TimeLine       []int    `json:"time_line"`
		CurveRank      []string `json:"curve_rank"`
		BenchCurveRank []string `json:"bench_curve_rank"`
	} `json:"data"`
}
