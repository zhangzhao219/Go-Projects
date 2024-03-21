package tracer_log

import (
	"encoding/json"
	"strconv"
	"testtracer/structs"
	"time"
)

const logBatchSize = 1

func Log(pipelineParam *structs.DspPipelineParam, filterName string, item *structs.AdsItem) {
	if !pipelineParam.TracerLogInfo {
		return
	}
	tracer := pipelineParam.TracerInfo
	items, ok := tracer[filterName]
	if !ok {
		items = []*structs.AdsItem{item}
		tracer[filterName] = items
		return
	}
	items = append(items, item)
	tracer[filterName] = items
}

func ItemsLog(pipelineParam *structs.DspPipelineParam, filterName string, items []*structs.AdsItem) {
	if !pipelineParam.TracerLogInfo {
		return
	}
	tracer := pipelineParam.TracerInfo
	res := make([]*structs.AdsItem, len(items))
	copy(res, items)
	tracer[filterName] = res
}

func RecordTracer(rm *structs.AdsResourceManager) {
	pipelineParam := rm.PipelineParam
	if pipelineParam.TracerLogInfo {
		tracerInfos := pipelineParam.TracerInfo
		for k, v := range tracerInfos {
			// 将adItems 等分成大小为 500 的切片
			averageAdsItems := fixedGrouping(v, logBatchSize)
			for _, adItems := range averageAdsItems {
				bs, _ := json.Marshal(convertToTracerLog(rm, k, adItems))
				TracerLog.Info(string(bs))
			}
		}
	}
}

func convertToTracerLog(rm *structs.AdsResourceManager, filterName string, items []*structs.AdsItem) structs.TransferPipelineDetailLog {
	// dspParam := rm.DspParam
	adTracerInfos := make([]*structs.AdTracerInfo, 0)
	for _, item := range items {
		info := structs.AdTracerInfo{
			Aiid:     item.AiidStr,
			Uid:      item.Uid,
			Apid:     strconv.Itoa(item.Apid),
			Bid:      strconv.Itoa(item.Bid),
			Ecpm:     0,
			Rankecpm: 0,
		}
		adTracerInfos = append(adTracerInfos, &info)
	}

	return structs.TransferPipelineDetailLog{
		PipelineName: filterName,
		// RequestId:      dspParam.ReqId,
		// Oaid:           dspParam.Oaid,
		// Imei:           dspParam.Imei,
		// Uuid:           dspParam.Uuid,
		// Onlyid:         dspParam.OnlyId,
		// UnitId:         dspParam.UnitId,
		LocalTimestamp: time.Now().UnixNano() / 1e6,
		Kind:           filterName,
		ServerName:     "serverNameZhangZhao",
		AdTracerInfos:  adTracerInfos,
	}
}

/*
*
将切片拆分成固定大小
param:
iter 待拆分数组
batchSize 拆分后数组定长
return:
[][]*structs.AdsItem
*/
func fixedGrouping(iter []*structs.AdsItem, batchSize int) [][]*structs.AdsItem {
	result := make([][]*structs.AdsItem, 0)
	if iter == nil || len(iter) == 0 {
		return result
	}
	length := len(iter)
	if length <= batchSize {
		result = append(result, iter)
		return result
	}
	temp := make([]*structs.AdsItem, 0)
	for _, value := range iter {
		if len(temp) < batchSize {
			temp = append(temp, value)
		} else {
			result = append(result, temp)
			temp = temp[:0]
			temp = append(temp, value)
		}
	}
	if len(temp) != 0 {
		result = append(result, temp)
	}
	return result
}
