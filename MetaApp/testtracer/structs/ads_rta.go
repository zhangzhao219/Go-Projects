package structs

import "gitlab.appshahe.com/service-cloud-rec/rec-utils.git/cp_thrift/ads_rta_thrift/ads_rta_thrift"

type RtaRequestParam struct {
	RequestId  string
	DeviceInfo *DeviceParam
}

type RtaTask struct {
	RtaId           string
	RtaRequestParam *RtaRequestParam
	Pkg             string
	Uid             string
}

type RtaResult struct {
	PkgUidStr string
	RtaDetail *ads_rta_thrift.RtaDetail
}

type PkgUidRtaId struct {
	Pkg   string
	Uid   string
	RtaId string
}

type UidTaskIdConfig struct {
	UidTaskIdMap map[string]string `json:"uidTaskIdMap"`
}

type RtaConfig struct {
	Enable             bool                     `json:"enable"`
	Timeout            int                      `json:"timeout"`
	TopN               int                      `json:"topN"`
	RtaMap             map[string]*RIdPkgConfig `json:"rtaMap"`
	PackageUidRtaIdMap map[string]string        `json:"packageUidRtaIdMap"`
}

type RIdPkgConfig struct {
	Enable     bool                `json:"enable"`
	PackageMap map[string][]string `json:"packageMap"`
}
