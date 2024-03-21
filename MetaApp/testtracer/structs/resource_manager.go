package structs

import (
	"github.com/bits-and-blooms/bitset"
	"gitlab.appshahe.com/service-cloud-rec/sort_base.git/base_global/request"
)

type AdsResourceManager struct {
	*request.ResourceManager
	Param                       *RequestParam
	DspParam                    *DspRequestParam
	AdItems                     []*AdsItem
	AdItemsInfo                 map[string]*AdsItem
	PredictUseCascadeFinalScore bool
	Ctx                         map[string]string
	BitMap                      *bitset.BitSet // 需要初始化 todo
	PipelineParam               *DspPipelineParam
	AdMap                       map[string]*AdsItem
	FilterData                  *DspFilterData
	PredictErr                  bool
	PredictCutAll               bool
	InternalInstallSet          map[string]bool
	RetentionPkgAdItems         map[string][]*AdsItem
	FilterRecord                map[string]int
}

type DspRequestParam struct {
	Uuid               string
	OnlyId             string
	AndroidId          string
	Imei               string
	Oaid               string
	DeviceBrand        string
	DeviceManufacturer string
	DeviceModel        string
	SystemVersion      string
	ChannelId          string
	ApkChannelId       string
	AppVersionName     string
	AppVersionCode     int
	SelfPackageName    string
	IosAndroid         string
	PandoraGroup       string
	ReqId              string
	UnitId             string
	Offset             int
	Size               int
	RefreshStatus      int
	LastShowGame       string
	Ip                 string
	InstallPos         int
	NetType            string
	NewUser            bool
	CategoryId         string
	RsConfigArr        []string
	GameId             int
	GamePackageName    string
	Keyword            string
	InfoMap            string
	QueryParam         string
	AdrecallParam      string
}
