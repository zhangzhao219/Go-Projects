package structs

type GameRecAdTransferResult2MQReqParam struct {
	ReqId           string `json:"reqid"`
	UnitId          string `json:"unitId"`
	Uuid            string `json:"uuid"`
	OnlyId          string `json:"onlyId"`
	GameIds         []int  `json:"gameIds"`
	ChannelPackage  string `json:"channelPackage"`
	AppVersionName  string `json:"appVersionName"`
	AppVersionCode  int    `json:"appVersionCode"`
	ChannelId       string `json:"channelId"`
	ApkChannelId    string `json:"apkChannelId"`
	SelfPackageName string `json:"selfPackageName"`
}

type GameRecAdTransferResult2MQResp struct {
	Code    int32                               `json:"code"`
	Message string                              `json:"message"`
	Data    *GameRecAdTransferResult2MQRespData `json:"data"`
}

type GameRecAdTransferResult2MQRespData struct {
	RequestId string `json:"requestId"`
}
