package structs

const (
	UnitIdIndex                = "adp:transfer:unit_id_index_new:%d_%d_%d" // {appType}_{deliveryType}_{businessType}
	UnitIdConfig               = "adp:platform:new_unit_id_config"
	CurrentDistributableAd     = "adp:platform:new_adBean_hash:%s" // yyyy-MM-dd
	ProfileAdIdeaTag           = "adp:platform:new_profile_ad_idea_tag"
	AdCostStatDate             = "new_ad_cost_stat_%dd"
	AdCostStatsHistory         = "new_ad_cost_stat_%sd_%s_%s"
	AdCostStatsRT              = "new_ad_cost_stat_rt_%s_%s"
	SmoothPacingPtr            = "adp:transfer:smooth_pace:ptr_new:%s"
	AdRunningPoolByScene       = "adp:ad_running_pool_by_scene:%d_%d"
	UidRunningPoolByScene      = "adp:new_account_by_scene:%d_%d"
	AdPlanConfigByScene        = "adp:platform:new_ad_plan_config:%d_%d_%d"
	BadCaseAdIcon              = "adp:transfer:sentinelAd:badcase_icon:%s"
	BadCaseAdHS                = "adp:transfer:sentinelAd:badcase_hot_search:%s"
	DeadAdIcon                 = "adp:transfer:sentinelAd:deadad_icon:%s"
	DeadAdHS                   = "adp:transfer:sentinelAd:deadad_hot_search:%s"
	DeadAdDP                   = "adp:transfer:sentinelAd:deadad_detail_page:%s"
	FreqThreshold              = "adp:transfer:new_freq_threshold"
	AdVersionConfig            = "adp:u:inst_lst:version"
	OnlyIdInstallAppList       = "adp:u:inst_lst:%s:%s"     //{onlyId}:{version}
	StatUidCost                = "statistics_uid_cost_%s"   //广告主对应消耗
	StatGroupCost              = "statistics_group_cost_%s" //广告组对应消耗
	StatPlanCost               = "statistics_plan_cost_%s"  //广告计划对应消耗
	OnlyIdSHow                 = "adp:showed:%s"            //{onlyId} onlyId 是否展示
	AdsCvrWeight               = "adp:transfer:ctr_model:new_cvr_weight"
	RandomCtrBase              = "adp:transfer:random_ctr_base"
	RandomCtrDelta             = "adp:transfer:random_ctr_delta"
	AdGlobalLocation           = "ad_global_location"
	AdLocation                 = "ad:location:%s"
	CurrentAdSupport           = "adp:platform:ad_support_hash:%s"
	AdOcpxConfigVersion        = "ads_sort:new_old_pkg_ocpx_hash_version"
	AdOcpxConfigKey            = "ads_sort:new_old_pkg_ocpx_hash_%s"
	AdAtrStatsHistory          = "ad_%s_atr_stat_%sd_%s"       // {type} {N} {version}
	AdDataStatsRealTime        = "ad_%s_data_stat_realtime_%s" // {type} {version}
	AdDataStatsRealTimeVersion = "ad_data_stat_realtime_version"
	GameRecTransferAdConfig    = "ads_sort:gamerec_transfer_ad_config"
)
