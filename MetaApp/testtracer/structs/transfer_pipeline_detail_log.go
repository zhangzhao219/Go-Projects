package structs

import "encoding/json"

const (
	TAG = "log_transfer_pipeline_detail"
)

type TransferPipelineDetailLog struct {
	PipelineName   string          `json:"pipeline_name"`
	RequestId      string          `json:"request_id"`
	Oaid           string          `json:"oaid"`
	Imei           string          `json:"imei"`
	Uuid           string          `json:"uuid"`
	Onlyid         string          `json:"onlyid"`
	UnitId         string          `json:"unit_id"`
	LocalTimestamp int64           `json:"local_timestamp"`
	Kind           string          `json:"kind"`
	ServerName     string          `json:"server_name"`
	AdTracerInfos  []*AdTracerInfo `json:"ad_tracer_infos"`
}

type AdTracerInfo struct {
	Aiid     string  `json:"aiid"`
	Uid      string  `json:"uid"`
	Apid     string  `json:"apid"`
	Bid      string  `json:"bid"`
	Ecpm     float64 `json:"ecpm"`
	Rankecpm float64 `json:"rankecpm"`
}

func (resultInfo TransferResultInfo) pack() string {
	jb, _ := json.Marshal(resultInfo)
	js := string(jb)
	return TAG + ": " + js
}
