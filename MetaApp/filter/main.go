package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/RoaringBitmap/roaring"
)

const (
	filter1 = "filter1"
	filter2 = "filter2"
	filter3 = "filter3"
	filter4 = "filter4"

	filter5 = "filter5"
	filter6 = "filter6"
	filter7 = "filter7"
)

type AdsItem struct {
	id uint32
	A  int
	B  int
}

type forceContentFunc struct {
	filter     func(item *AdsItem) bool
	filterName string
}

type forceContentMultiFunc struct {
	filter     func(item []*AdsItem) []*AdsItem
	filterName string
}

type forceContentMultiBitMapFunc struct {
	filter     func(item []*AdsItem) *roaring.Bitmap
	filterName string
}

type FilterMultiFunc func(itemList []*AdsItem) []*AdsItem

var contForceFilterMap = map[string]forceContentFunc{

	filter1: {filter: Filter1Single, filterName: filter1},
	filter2: {filter: Filter2Single, filterName: filter2},
	filter3: {filter: Filter3Single, filterName: filter3},
	filter4: {filter: Filter4Single, filterName: filter4},

	filter5: {filter: Filter5Single, filterName: filter5},
	filter6: {filter: Filter6Single, filterName: filter6},
	filter7: {filter: Filter7Single, filterName: filter7},
}

// contForceMultiFilterMap 多个ad遍历filter map
var contForceMultiFilterMap = map[string]forceContentMultiFunc{
	filter5: {filter: Filter5Multi, filterName: filter5},
	filter6: {filter: Filter6Multi, filterName: filter6},
	filter7: {filter: Filter7Multi, filterName: filter7},

	filter1: {filter: Filter1Multi, filterName: filter1},
	filter2: {filter: Filter2Multi, filterName: filter2},
	filter3: {filter: Filter3Multi, filterName: filter3},
	filter4: {filter: Filter4Multi, filterName: filter4},
}

// contForceMultiFilterMap 多个ad遍历filter map
var contForceBitMap = map[string]forceContentMultiBitMapFunc{
	filter5: {filter: Filter5BitMap, filterName: filter5},
	filter6: {filter: Filter6BitMap, filterName: filter6},
	filter7: {filter: Filter7BitMap, filterName: filter7},

	filter1: {filter: Filter1BitMap, filterName: filter1},
	filter2: {filter: Filter2BitMap, filterName: filter2},
	filter3: {filter: Filter3BitMap, filterName: filter3},
	filter4: {filter: Filter4BitMap, filterName: filter4},
}

func main() {
	items := prepareFilterData()

	// var contForceMultiFilterInfos []forceContentMultiFunc
	// for _, filterMultiName := range []string{"filter1", "filter2", "filter3", "filter4", "filter5", "filter6", "filter7"} {
	// 	if filterMulti, registered := contForceMultiFilterMap[filterMultiName]; registered {
	// 		contForceMultiFilterInfos = append(contForceMultiFilterInfos, filterMulti)
	// 	}
	// }

	var contForceMultiFilterInfos []forceContentMultiBitMapFunc
	for _, filterMultiName := range []string{"filter1", "filter2", "filter3", "filter4", "filter5", "filter6", "filter7"} {
		if filterMulti, registered := contForceBitMap[filterMultiName]; registered {
			contForceMultiFilterInfos = append(contForceMultiFilterInfos, filterMulti)
		}
	}

	// starttime2 := time.Now().UnixNano()
	// items = forceContentMultiFilter(items, contForceMultiFilterInfos)
	// fmt.Println("forceContentMultiFilter time ", time.Now().UnixNano()-starttime2)

	// var contForceFilterInfos []forceContentFunc
	// for _, filterName := range []string{"filter1", "filter2", "filter3", "filter4", "filter5", "filter6", "filter7"} {
	// 	if filter, registered := contForceFilterMap[filterName]; registered {
	// 		contForceFilterInfos = append(contForceFilterInfos, filter)
	// 	}
	// }

	starttime3 := time.Now().UnixNano()
	bitMap := roaring.BitmapOf()
	fmt.Println("bitmap time ", time.Now().UnixNano()-starttime3)
	for _, ads := range items {
		bitMap.Add(ads.id)
	}
	fmt.Println("bitmap time ", time.Now().UnixNano()-starttime3)
	fmt.Println(bitMap.GetCardinality())
	bitMapfilterlist := bitMapFilter(items, contForceMultiFilterInfos)
	for _, result := range bitMapfilterlist {
		fmt.Println(result.GetCardinality())
		bitMap.And(result)
	}
	fmt.Println(bitMap.GetCardinality())

	fmt.Println("bitmap time ", time.Now().UnixNano()-starttime3)
	adsItems := make([]*AdsItem, 0, len(items))
	for _, ad := range items {
		if bitMap.Contains(ad.id) {
			adsItems = append(adsItems, ad)
		}
	}
	items = adsItems

	fmt.Println("bitmap time ", time.Now().UnixNano()-starttime3)

	// starttime := time.Now().UnixNano()
	// items = forceContentFilter(items, contForceFilterInfos)
	// fmt.Println("forceContentFilter time ", time.Now().UnixNano()-starttime)
	fmt.Println(bitMap.GetCardinality())
	// for _, item := range items {
	// 	if item.A == 1 || item.A == 2 || item.A == 3 || item.A == 4 {
	// 		fmt.Println("alert")
	// 	}
	// 	if item.B == 5 || item.B == 6 || item.B == 7 {
	// 		fmt.Println("alert")
	// 	}
	// }
}

func bitMapFilter(items []*AdsItem, forceContentMultiFilters []forceContentMultiBitMapFunc) []*roaring.Bitmap {
	bitmapList := make([]*roaring.Bitmap, 0)
	for _, filterIns := range forceContentMultiFilters {
		bitmapList = append(bitmapList, filterIns.filter(items))
	}
	return bitmapList
}

func prepareFilterData() []*AdsItem {
	adsitem := make([]*AdsItem, 0)
	for i := 100000; i < 100000+200000; i++ {
		temp := &AdsItem{
			id: uint32(i),
			A:  rand.Intn(10),
			B:  rand.Intn(10),
		}
		adsitem = append(adsitem, temp)
	}
	return adsitem
}
