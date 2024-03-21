package structs

import (
	"github.com/bits-and-blooms/bitset"
)

const (
	DeliveryTypeAppDownload = 0
	DeliveryTypeLandingPage = 1
)

type AdsItem struct {
	Aiid                     int
	Uid                      string
	MediaType                int
	Apid                     int
	Agid                     int
	MediaUrl                 string
	Width                    *int
	Height                   *int
	Icon                     string
	Title                    string
	Intro                    string
	Description              string
	GameId                   int64
	BigPic                   string
	GifPic                   string
	ImgUrls                  []string
	Duration                 int
	Type                     int
	DownloadType             int
	DownloadUrl              string
	DownloadPkg              string
	WebUrl                   string
	Deeplink                 string
	ThirdMonitor             string
	PayType                  int
	Budget                   int
	Status                   int
	ScheduleDay              string
	StartDay                 int64
	EndDay                   int64
	SpeedType                int
	Schedule                 int
	PlanDayBudget            int
	GroupDayBudget           int
	UserDayBudget            int
	RemainingBudget          int
	TodayRemainingBudgetTime int
	AreaInfo                 string
	Freq                     int
	NameKeywords             []string
	UserBalance              int
	InternalInstall          bool
	AllowJumpMarket          bool
	AdRange                  int
	AppSize                  int
	TradeId                  int
	TradeCode                string
	TradeName                string
	CreateWay                int
	AdIdeaKeywords           string
	AppInfo                  *AdvertiserAppInfo
	CategoryFirst            string
	CategorySecond           string
	CategoryThird            string
	VideoMd5                 string
	IdeaName                 string
	CrowdPackBOList          []AdPlanCrowdPack
	PkgInstallTarget         bool
	BobtailVersionList       []string
	CreateTime               int64
	Bid                      int
	InternalAtr              float64
	ExternalAtr              float64
	RealCtr                  float64
	AppActivateRate          float64
	GspBid                   int
	AreaCode                 string
	AdTagC                   *AdTagCode
	AiidStr                  string
	AiidInt                  int
	InternalInstallStr       string
	FormatMediaUrl           string
	BitSize                  int
	AreaCodeBitSet           *bitset.BitSet
	UseLpTemplate            bool
	LpImgUrl                 string
	AdAppMd5                 string
	OcpxBid                  int
	OcpxFlag                 bool
}

type AdvertiserAppInfo struct {
	AppName          string
	AppVersion       string
	DeveloperName    string
	PackageName      string
	PrivacyAgreement string
}

type AdPlanCrowdPack struct {
	ApId        int64
	CrowdPackId int64
	Name        string
	Num         int
	Type        int
	RedisKey    string
}

type ServExtras struct {
	UnitId    string
	RequestId string
	Aiid      string
	GameId    string
	Su        string
	Sb        string
	Sgb       string
	Sp        string
	Sr        string
	Ss        string
	Ro        string
	Rt        string
	Ct        string
	Ls        string
}
