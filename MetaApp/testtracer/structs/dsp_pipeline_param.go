package structs

import (
	"math/rand"
)

type DspPipelineParam struct {
	TodayFormat                           string
	AbConfigResult                        *ABConfigResult //ab试验参数
	FreqConfig                            *FreqControlNum //全局缓存
	AdPlanConfigMap                       map[int]*AdPlanConfig
	UnitIdParam                           *SdkUserUnitIdParam
	NewAdSet                              map[int]bool
	ContextFeature                        *ContextFeature        //存feature collector
	UserFeature                           *UserFeature           //放feature collector
	BiddingFactorMap                      map[int]*BiddingFactor // 放全局缓存
	UuidHasBusinessTag                    bool
	UuidHasRecTag                         bool
	UuidHasCrowdPackTag                   bool
	Scene                                 string
	TracerLogInfo                         bool
	TopKNewAdSet                          map[int]bool
	UserFids                              []int64                 //透传参数可删
	ItemFidsMap                           map[int][]int64         //之前的透传参数可删
	BiddingInfoMap                        map[string]*BiddingInfo //sort_base item存储 可以删
	RandSource                            *rand.Rand
	LandingPageAiidMap                    map[string]bool
	LandingPageAdForceInsertItems         []*AdsItem
	LandingPageWeChatAdForceInsertAiidMap map[string]bool
	TracerInfo                            map[string][]*AdsItem
	RemoteAB                              *RemoteABConfig
	RtaConfig                             *RtaConfig
	UidProductIdConfig                    *UidTaskIdConfig
	ForceInsertAdItemsMap                 map[string]map[string][]*AdsItem
	DownloadNotInstalledAdIdMap           map[string]bool
}

type RemoteABConfig struct {
	CpmStatsCutOffFilter     *CpmConfig                `json:"cpm_stats_cut_off_filter"`
	ExportNumFilter          *ExportConfig             `json:"export_num_filter"`
	PreProcessor             *PreProcessorConfig       `json:"pre_processor"`
	LandingPageFilter        *LandingPageFilterConfig  `json:"landing_page_filter"`
	NewAdFilter              *NewAdFilterConfig        `json:"new_ad_filter"`
	PostRankingConfig        *PostRankingConfig        `json:"post_ranking_step"`
	DspRecallConfig          *DspRecallConfig          `json:"dsp_recall_config"`
	LandingPageRankingConfig *LandingPageRankingConfig `json:"landing_page_ranking_config"`
	ForceInsertConfig        *ForceInsertConfig        `json:"force_insert_config"`
}

type CpmConfig struct {
	Cpm   int `json:"cpm,omitempty"`
	Max   int `json:"max,omitempty"`
	Show  int `json:"show,omitempty"`
	TopK1 int `json:"topK1,omitempty"`
	TopK2 int `json:"topK2,omitempty"`
}

type ExportConfig struct {
	ExportItemLimit              int     `json:"exportItemLimit"`
	ExportNewAd                  float64 `json:"exportNewAd"`
	DefaultExportSwitch          bool    `json:"defaultExportSwitch"`
	OnlineFeatureLogLimit        int     `json:"onlineFeatureLogLimit"`
	LandingPageForcedInsertLimit int     `json:"landingPageForcedInsertLimit"`
}

type LandingPageRankingConfig struct {
	IsPredict bool `json:"isPredict"`
}

type DspRecallConfig struct {
	DspRecallTopK int `json:"dsp_recall_topK"`
}

type ForceInsertConfig struct {
	Activated            *ForceInsert `json:"activated"`
	DownloadNotInstalled *ForceInsert `json:"download_not_installed"`
}

type ForceInsert struct {
	PkgMap map[string]bool `json:"uidPkgMap"`
	Limit  int64           `json:"limit"`
}

type PostRankingConfig struct {
	CtrRandom *CtrRandom `json:"ctrRandom"`
	ATRRandom *AtrRandom `json:"atrRandom"`
}

type CtrRandom struct {
	RandomCtrBase  float64 `json:"randomCtrBase"`
	RandomCtrDelta float64 `json:"randomCtrDelta"`
}

type AtrRandom struct {
	RandomAtrBase  float64 `json:"randomAtrBase"`
	RandomAtrDelta float64 `json:"randomAtrDelta"`
}

type PreProcessorConfig struct {
	TracerLogSwitch    bool    `json:"tracerLogSwitch"`
	TracerLogInfoRatio float64 `json:"tracerLogInfoRatio"`
}

type NewAdFilterConfig struct {
	NewAdMaxNum             int `json:"newAdMaxNum"`
	NewAdLimitNum           int `json:"newAdLimitNum"`
	NewAdNewAccountLimitNum int `json:"newAdNewAccountLimitNum"`
	NewAdOldAccountLimitNum int `json:"newAdOldAccountLimitNum"`
}

type LandingPageFilterConfig struct {
	LandingPageAdLimitNum   int               `json:"landingPageAdLimitNum"`
	ProtectRtaAd            bool              `json:"protectRtaAd"`
	ProtectRtaAdMap         map[string]string `json:"protectRtaAdMap"`
	ProtectForceInsertAdMap map[string]string `json:"protectForceInsertAdMap"`
}

type ABConfigJson struct {
	Key            string `json:"key"`
	Vid            string `json:"vid"`
	IsNewMemberDay bool   `json:"is_new_member_day"`
	HasJoin        int    `json:"has_join"`
	Val            string `json:"val"`
}

type ABConfigResult struct {
	HitExperimentMap       map[string]string        `json:"hitExperimentMap,omitempty"`
	HitExperimentConfigMap map[string]*ABConfigJson `json:"hitExperimentConfigMap,omitempty"`
	AllVid                 []string                 `json:"allVid,omitempty"`
	NewVid                 []string                 `json:"newVid,omitempty"`
	SwitchAllVid           []string                 `json:"switchAllVid,omitempty"`
	SwitchNewVid           []string                 `json:"switchNewVid,omitempty"`
	JoinVid                []string                 `json:"joinVid,omitempty"`
}

type FreqControlNum struct {
	UserFreqNum    int
	AdFreqNum      int
	AdFreqNumIn30m int
}

type AdPlanConfig struct {
	Apid                int
	AppType             int
	DeliveryType        int
	BusinessType        int
	TtOrientation       int
	BusinessOrientation int
	CrowdPkgOrientation int
}

type SdkUserUnitIdParam struct {
	Id                 int
	AppId              int
	UnitId             string
	NewAdSupportFactor float64
	Ecpm               int
	Order              int
	Status             int
	AppType            int
	DeliveryType       int
	BusinessType       int
	StrategyType       int
	BiddingType        int
	DefaultTransfer    bool
}

type ContextFeature struct {
	ContextAppChannel        string
	ContextBrand             string
	ContextSysVersion        string
	ContextNet               string
	AreaCode                 string
	ContextLocation          string
	ContextAppVersion        string
	ContextSdkVersion        string
	ContextUnitId            string
	ContextInstallNum        string
	ContextWeekday           string
	ContextHour              string
	ContextTs                string
	ContextFreeSpace         string
	ContextAppId             string
	ContextAppPackage        string
	Onlyid                   string
	ContextOs                string
	ContextOsVersion         string
	DetailGamePackageName    string
	DetailsPageGameId        string
	IsPreciseMatch           bool
	SearchAdPlacement        string
	AnalyzesResultList       []string
	AnalyzesResultFilterList []string
	FilteredKeyword          string
	FilteredKeywordMd5       string
	AppCategories            string
	AdxAppId                 string
	AdWidth                  int
	AdHeight                 int
	ScreenDensityDpi         string
	ScreenWidth              string
	TotalMem                 string
	DeviceBrand              string
	FreeSpace                string
	InstalledAppPkgSet       map[string]bool
}

type UserFeature struct {
	Exists                        bool            `json:"exists"`
	Uuid                          string          `json:"uuid"`
	UserSex                       string          `json:"userSex"`
	UserAge                       string          `json:"userAge"`
	UserRegisterTime              int64           `json:"userRegisterTime"`
	UserActiveDayLastMonth        string          `json:"userActiveDayLastMonth"`
	UserPlayGameNumLastMonth      string          `json:"userPlayGameNumLastMonth"`
	UserLaunchTime                string          `json:"userLaunchTime"`
	UserLast3DayCtr               float64         `json:"userLast3DayCtr"`
	UserLast3DayCvr               float64         `json:"userLast3DayCvr"`
	UserPayAmountLastDay          int             `json:"userPayAmountLastDay"`
	UserPayAmountLastWeek         int             `json:"userPayAmountLastWeek"`
	UserPayAmountLastMonth        int             `json:"userPayAmountLastMonth"`
	UserPayCountLastDay           int             `json:"userPayCountLastDay"`
	UserPayCountLastWeek          int             `json:"userPayCountLastWeek"`
	UserPayCountLastMonth         int             `json:"userPayCountLastMonth"`
	UserPayAmountAll              int             `json:"userPayAmountAll"`
	UserPayCountAll               int             `json:"userPayCountAll"`
	UserPayGameNum                int             `json:"userPayGameNum"`
	UserShowAdLast3Day            string          `json:"userShowAdLast3Day"`
	UserClickAdLast3Day           string          `json:"userClickAdLast3Day"`
	UserShowAdNumLast3Day         string          `json:"userShowAdNumLast3Day"`
	UserClickAdNumLast3Day        string          `json:"userClickAdNumLast3Day"`
	UserInstallAdLast3Day         string          `json:"userInstallAdLast3Day"`
	UserPlayGameTagLastDay        string          `json:"userPlayGameTagLastDay"`
	UserPlayGameTagLastWeek       string          `json:"userPlayGameTagLastWeek"`
	UserClickGameTagLastDay       string          `json:"userClickGameTagLastDay"`
	UserClickGameTagLastWeek      string          `json:"userClickGameTagLastWeek"`
	UserPlayGameCategoryLastDay   string          `json:"userPlayGameCategoryLastDay"`
	UserPlayGameCategoryLastWeek  string          `json:"userPlayGameCategoryLastWeek"`
	UserClickGameCategoryLastDay  string          `json:"userClickGameCategoryLastDay"`
	UserClickGameCategoryLastWeek string          `json:"userClickGameCategoryLastWeek"`
	UserFavoriteGameList          string          `json:"userFavoriteGameList"`
	UserPayGameList               string          `json:"userPayGameList"`
	UserTTTagSet                  map[string]bool `json:"userTtTagSet,omitempty"`
	UserTTInstalledTagSet         map[string]bool `json:"userTtInstalledTagSet,omitempty"`
	UserAdTagMatchNum             int             `json:"userAdTagMatchNum"`
	UuidHashCode20w               int             `json:"uuidHashCode20W"`
	UserClassifyLevel2Tags        string          `json:"userClassifyLevel2Tags"`
	RecUserInterestList           string          `json:"recUserInterestList"`
	RecUserTagList                string          `json:"recUserTagList"`
	RecUserPlayPkg1d              string          `json:"recUserPlayPkg1D"`
	RecUserPlayPkg7d              string          `json:"recUserPlayPkg7D"`
	RecUserPlayPkg30d             string          `json:"recUserPlayPkg30D"`
	UserMonCtrSammo               string          `json:"userMonCtrSammo"`
	UserWeekCtrSammo              string          `json:"userWeekCtrSammo"`
	UserClickNum3Sammo            string          `json:"userClickNum3Sammo"`
	UserShowNum3Sammo             string          `json:"userShowNum3Sammo"`
	UserClickAd3Sammo             string          `json:"userClickAd3Sammo"`
	UserShowAd3Sammo              string          `json:"userShowAd3Sammo"`
	UserSearchClickList           string          `json:"userSearchClickList"`
	UserCategories                string          `json:"userCategories"`
}

type BiddingFactor struct {
	NewAd               bool    `json:"newAd"`
	NewAdSupportFactor  float64 `json:"newAdSupportFactor"`
	NewUid              bool    `json:"newUid"`
	NewUidSupportFactor float64 `json:"newUidSupportFactor"`
	CvrWeight           float64 `json:"cvrWeight"`
	ColdStartExp        string  `json:"coldStartExp"`
	ColdStart           bool    `json:"coldStart"`
	Alpha               float64 `json:"alpha"`
	Beta                float64 `json:"beta"`
}

type BiddingInfo struct {
	Cvr          float64
	Cpm          int
	RankCpm      int
	PreRankCpm   int
	Ctr          float64
	Itr          float64
	Atr          float64
	PreRankScore float64
}
