package main

import (
	"fmt"
	"time"
)

// type CalculateItem interface {
// 	RecallResult() RecallResult
// }

// type RecallResult interface {
// 	RId() string
// }

// type ResourceManager[T CalculateItem] struct {
// 	items []T
// }

// type Item struct {
// 	WithRecallResult
// }

// type WithRecallResult struct {
// 	RecallId       string          `json:"id"`
// 	RecallPriority int             `json:"priority"`
// 	RecallName     string          `json:"recall_name"`
// 	RecallNameSet  map[string]bool `json:"recall_name_set"`
// 	RecallScore    float64         `json:"recall_score"`
// 	RecallRanking  int             `json:"recall_ranking"`
// }

// func (i *WithRecallResult) RecallResult() RecallResult {
// 	return i
// }

// func (i *WithRecallResult) RId() string {
// 	return i.RecallId
// }

// func (r *ResourceManager[T]) Items() []T {
// 	return r.items
// }

// type AdsItem struct {
// 	*Item
// 	AiidStr string
// }

type A struct {
	*B
}

type B struct {
	fid []int64
}

func main() {
	// rm := &ResourceManager[*AdsItem]{
	// 	[]*AdsItem{
	// 		&AdsItem{
	// 			AiidStr: "fdsfds",
	// 		},
	// 	},
	// }
	// items := rm.Items()
	// for _, elem := range items {
	// 	elem.Item = &Item{}
	// 	elem.RecallId = elem.AiidStr
	// }
	// fmt.Println(*rm)
	// fmt.Println(items[0].Item)
	// name2Fids := map[string][]int64{}

	// for i := 0; i < 5; i++ {
	// 	name2Fids["a"] = append(name2Fids["a"], 12)
	// 	name2Fids["b"] = append(name2Fids["b"], 13)
	// }
	// fmt.Println(name2Fids)

	// rm1 := make([]*A, 1)
	// rm2 := make([]*A, 1)
	// rm1[0] = &A{}
	// // rm2[0] = &A{
	// // 	B: &B{
	// // 		fid: []int64{},
	// // 	},
	// // }
	// // rm1[0].fid = []int64{4, 5, 6, 7}
	// copy(rm2, rm1)
	// for index, item := range rm1 {
	// 	tempItem := *item
	// 	rm2[index] = &tempItem
	// }
	// // fmt.Println(rm1, rm2)
	// fmt.Println(rm1[0].fid)
	// fmt.Println(searchRange([]int{0, 10, 20, 30}, 40))
	fmt.Println(time.Duration(12) * time.Hour)

}

func searchRange(nums []int, target int) int {
	var mid int
	low, high := 0, len(nums)-1
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return high
}
