package structs

type TransferResultInfo struct {
	RequestId        string                 `json:"requestId"`
	DspId            string                 `json:"dspId"`
	AdStatsInfo      *AdStatsInfo           `json:"adStatsInfo,omitempty"`
	AdStatsInfos     []*AdStatsInfo         `json:"adStatsInfos,omitempty"`
	Bid              int                    `json:"bid"`
	GspBid           int                    `json:"gspBid"`
	PayType          int                    `json:"payType"`
	TransferType     string                 `json:"transferType"`
	PassThroughParam *StatsPassThroughParam `json:"passThroughParam,omitempty"`
	TransferParam    *TransferParam         `json:"transferParam,omitempty"`
}

type AdStatsInfo struct {
	Aiid                            int      `json:"aiid"`
	Uid                             string   `json:"uid"`
	MediaType                       int      `json:"mediaType"`
	Apid                            int      `json:"apid"`
	Agid                            int      `json:"agid"`
	MediaUrl                        string   `json:"mediaUrl"`
	Icon                            string   `json:"icon"`
	Title                           string   `json:"title"`
	Intro                           string   `json:"intro"`
	Description                     string   `json:"description"`
	GameId                          int      `json:"gameId"`
	BigPic                          string   `json:"bigPic"`
	GifPic                          string   `json:"gifPic"`
	ImgUrls                         []string `json:"imgUrls,omitempty"`
	Duration                        int      `json:"duration"`
	Type                            int      `json:"type"`
	DownloadType                    int      `json:"downloadType"`
	DownloadUrl                     string   `json:"downloadUrl"`
	DownloadPkg                     string   `json:"downloadPkg"`
	WebUrl                          string   `json:"webUrl"`
	Bid                             int      `json:"bid"`
	PayType                         int      `json:"payType"`
	Cpm                             int      `json:"cpm"`
	RankCpm                         int      `json:"rankCpm"`
	PreRankCpm                      int      `json:"preRankCpm"`
	Ctr                             float64  `json:"ctr"`
	Itr                             float64  `json:"itr"`
	Cvr                             float64  `json:"cvr"`
	Atr                             float64  `json:"atr"`
	InternalAtr                     float64  `json:"internalAtr"`
	ExternalAtr                     float64  `json:"externalAtr"`
	InternalInstall                 bool     `json:"internalInstall"`
	GspBid                          int      `json:"gspBid"`
	OcpxBid                         int      `json:"ocpxBid"`
	IdeaName                        string   `json:"ideaName"`
	PreRankScore                    float64  `json:"preRankScore"`
	DownloadNotInstalledForceInsert bool     `json:"downloadNotInstalledForceInsert"`
}

type StatsPassThroughParam struct {
	AbConfigResult           *ABConfigResult        `json:"abConfigResult,omitempty"`
	UuidHasBusinessTag       bool                   `json:"uuidHasBusinessTag"`
	UuidHasRecTag            bool                   `json:"uuidHasRecTag"`
	UuidHasCrowdPackTag      bool                   `json:"uuidHasCrowdPackTag"`
	BiddingFactorMap         map[int]*BiddingFactor `json:"biddingFactorMap,omitempty"`
	OnlineFeatureLogLimit    int                    `json:"onlineFeatureLogLimit"`
	ContextWeekday           string                 `json:"contextWeekday"`
	ContextHour              string                 `json:"contextHour"`
	ContextTs                string                 `json:"contextTs"`
	UserFeature              *UserFeature           `json:"userFeature,omitempty"`
	IsPreciseMatch           bool                   `json:"isPreciseMatch"`
	DetailGamePackageName    string                 `json:"detailGamePackageName"`
	DetailsPageGameId        string                 `json:"detailsPageGameId"`
	SearchAdPlacement        string                 `json:"searchAdPlacement"`
	AnalyzesResultList       []string               `json:"analyzesResultList,omitempty"`
	AnalyzesResultFilterList []string               `json:"analyzesResultFilterList,omitempty"`
	BiddingType              string                 `json:"biddingType"`
	Province                 string                 `json:"province"`
	AreaCode                 string                 `json:"areaCode"`
}

type TransferParam struct {
	UnitId                int         `json:"unitId"`
	Libragroup            string      `json:"libragroup"`
	Uuid                  string      `json:"uuid"`
	RequestId             string      `json:"requestId"`
	ApkChannelId          string      `json:"apkChannelId"`
	DetailsPageGameId     string      `json:"detailsPageGameId"`
	DetailGamePackageName string      `json:"detailGamePackageName"`
	Keyword               string      `json:"keyword"`
	DeviceInfo            *DeviceInfo `json:"deviceInfo,omitempty"`
	AppInfo               *AppInfo    `json:"appInfo,omitempty"`
	NetInfo               *NetInfo    `json:"netInfo,omitempty"`
}

type DeviceInfo struct {
	OnlyId             string `json:"onlyId"`
	Imei               string `json:"imei"`
	Oaid               string `json:"oaid"`
	DeviceManufacturer string `json:"deviceManufacturer"`
	DeviceBrand        string `json:"deviceBrand"`
	DeviceModel        string `json:"deviceModel"`
	DeviceSysVersion   string `json:"deviceSysVersion"`
	DeviceOs           string `json:"deviceOs"`
}

type AppInfo struct {
	AppPackage     string `json:"appPackage"`
	AppVersionCode int    `json:"appVersionCode"`
	AppVersionName string `json:"appVersionName"`
	AppChannel     string `json:"appChannel"`
	AppName        string `json:"appName"`
	SdkVersionName string `json:"sdkVersionName"`
	SdkVersionCode int    `json:"sdkVersionCode"`
}

type LocationInfo struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type NetInfo struct {
	NetworkState int    `json:"networkState"`
	IpAddress    string `json:"ipAddress"`
}
