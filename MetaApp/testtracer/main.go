package main

import (
	"fmt"
	"testtracer/structs"
)

// func outputLog(name string, timenum int) {
// 	num := 0
// 	for {
// 		time.Sleep(time.Millisecond * time.Duration(timenum))
// 		fmt.Println(time.Now())
// 		rm := &structs.AdsResourceManager{
// 			AdItems: []*structs.AdsItem{
// 				&structs.AdsItem{
// 					Bid: 1,
// 				},
// 				&structs.AdsItem{
// 					Bid: 2,
// 				},
// 			},
// 			PipelineParam: &structs.DspPipelineParam{
// 				TracerLogInfo: true,
// 				TracerInfo:    make(map[string][]*structs.AdsItem, 0),
// 			},
// 		}

// 		items := rm.AdItems

// 		tracer_log.Log(rm.PipelineParam, name, items[0])
// 		fmt.Println(rm.PipelineParam.TracerInfo["100"][0])
// 		rm.AdItems = items[:1]
// 		rm.AdItems[0].Bid = 3
// 		fmt.Println(rm.PipelineParam.TracerInfo["100"][0])
// 		num += 1
// 		tracer_log.RecordTracer(rm)
// 	}
// }

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

func fixedGrouping2(iter []*structs.AdsItem, batchSize int) [][]*structs.AdsItem {
	if iter == nil || len(iter) == 0 {
		return make([][]*structs.AdsItem, 0)
	}
	result := make([][]*structs.AdsItem, 0, (len(iter)/batchSize)+1)
	length := len(iter)
	if length <= batchSize {
		result = append(result, iter)
		return result
	}
	temp := make([]*structs.AdsItem, 0, batchSize)
	for _, value := range iter {
		if len(temp) < batchSize {
			temp = append(temp, value)
		} else {
			result = append(result, temp)
			temp = make([]*structs.AdsItem, 0, batchSize)
			temp = append(temp, value)
		}
	}
	if len(temp) != 0 {
		result = append(result, temp)
	}
	return result
}

func main() {
	// go outputLog(strconv.FormatInt(int64(100), 10), 1000)
	adItems := []*structs.AdsItem{
		{Aiid: 0}, {Aiid: 1}, {Aiid: 2}, {Aiid: 3}, {Aiid: 4}, {Aiid: 5}, {Aiid: 6}, {Aiid: 7}, {Aiid: 8}, {Aiid: 9}, {Aiid: 10},
	}
	result := fixedGrouping2(adItems, 11)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Println(result[i][j])
		}
	}
	// time.Sleep(time.Hour)
	// parsejson()
}

// type TransferPipelineDetailLog struct {
// 	PipelineName   string         `json:"pipeline_name"`
// 	RequestId      string         `json:"request_id"`
// 	Oaid           string         `json:"oaid"`
// 	Imei           string         `json:"imei"`
// 	Uuid           string         `json:"uuid"`
// 	Onlyid         string         `json:"onlyid"`
// 	UnitId         string         `json:"unit_id"`
// 	LocalTimestamp int64          `json:"local_timestamp"`
// 	Kind           string         `json:"kind"`
// 	ServerName     string         `json:"server_name"`
// 	AdTracerInfos  []AdTracerInfo `json:"ad_tracer_infos"`
// }

// type AdTracerInfo struct {
// 	Aiid     string  `json:"aiid"`
// 	Uid      string  `json:"uid"`
// 	Apid     string  `json:"apid"`
// 	Bid      string  `json:"bid"`
// 	Ecpm     float64 `json:"ecpm"`
// 	Rankecpm float64 `json:"rankecpm"`
// }

// func parsejson() {
// 	fileHanle, _ := os.OpenFile("/mnt/d/testtracer/ads-sort-202303161645.log", os.O_RDONLY, 0666)

// 	defer fileHanle.Close()

// 	readBytes, _ := ioutil.ReadAll(fileHanle)

// 	results := strings.Split(string(readBytes), "\n")
// 	mp := make(map[string]int)
// 	for i := 0; i < len(results); i++ {
// 		strt := results[i]
// 		a := &TransferPipelineDetailLog{}
// 		json.Unmarshal([]byte(strt), a)
// 		_, ok := mp[a.PipelineName]
// 		if !ok {
// 			mp[a.PipelineName] = 0
// 		}
// 		tmp := a.AdTracerInfos
// 		for j := 0; j < len(tmp); j++ {
// 			if tmp[j].Aiid == "239868" {
// 				mp[a.PipelineName] += 1
// 			}
// 		}
// 	}
// 	fmt.Println(mp)
// }
