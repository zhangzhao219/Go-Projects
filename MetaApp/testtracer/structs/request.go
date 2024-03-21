package structs

import (
	"time"

	"gitlab.appshahe.com/service-cloud-rec/rec-utils.git/component/ab_config"
)

type RequestParam struct {
	Uuid             string                       `json:"uuid"`
	TopK             int                          `json:"top_k"`
	Ip               string                       `json:"ip"`
	RefreshStatus    *int                         `json:"refreshStatus"`
	ReqId            string                       `json:"reqId"`
	IsDebug          bool                         `json:"debug"`
	NewUser          bool                         `json:"new_user"`
	Pkg              string                       `json:"pkg"` //用户点击此游戏进入内流
	ChannelId        string                       `json:"channel_id"`
	SocialName       string                       `json:"self_package_name"`
	SourcePackage    string                       `json:"super_game_package"` // 该参数有可能为空（重构版无法获取该参数）
	DeviceName       string                       `json:"device_name"`
	CategoryId       string                       `json:"category_id"`
	Keyword          string                       `json:"keyword"` // 搜索查询字段
	Page             int                          `json:"page"`
	Size             int                          `json:"size"`
	Offset           int                          `json:"offset"` // 分页参数：服务端偏移量
	DeviceBrand      string                       `json:"deviceBrand"`
	AppVersionCode   int                          `json:"appVersionCode"`
	EdgeRecSessionId string                       `json:"edgeRecSessionId"`
	RsConfigArr      []string                     `json:"rs_config_arr"`
	AdModelType      string                       `json:"ad_model_type"`  // for ad
	AppId            string                       `json:"appid"`          // for adlm
	ExtraRsConf      ab_config.Config             `json:"extra_rs_conf"`  // for debug
	OnlyId           string                       `json:"only_id"`        //for lock area
	DeviceModel      string                       `json:"device_model"`   //for lock area
	SystemVersion    string                       `json:"system_version"` //for lock area
	ExtraItemFeature map[string]map[string]string `json:"extra_item_feature"`
	PreRanked        string                       `json:"preRanked"`      // todo 某一个合适的版本需要删掉这个参数，因为对接问题，广告传过来一个不需要的参数，严格模式必须接
	TimeStamp        time.Time                    `json:"timestamp"`      // for total cost time
	AdSourceList     []string                     `json:"ad_source_list"` // for ad
}
type RespParam struct {
	Code     int     `json:"code"`
	Msg      string  `json:"msg"`
	List     []*Item `json:"list"`
	UserFids []int64 `json:"user_fids"`
}

type Item struct {
	Id       string  `json:"id"`
	Score    float64 `json:"score"`
	ItemFids []int64 `json:"item_fids"`
}

type DebugResp struct {
	Code     int                    `json:"code"`
	Msg      string                 `json:"msg"`
	List     []*DebugItem           `json:"list"`
	Features map[string]interface{} `json:"features"`
	Fids     []int64                `json:"fids"`
}

type DebugItem struct {
	Item
	Idx        int     `json:"idx"`
	RecallName string  `json:"recall_name"`
	Expr       string  `json:"expr"`
	Fids       []int64 `json:"fids"`
}
