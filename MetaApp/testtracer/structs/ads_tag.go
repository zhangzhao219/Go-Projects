package structs

import (
	"github.com/bits-and-blooms/bitset"
)

type Void struct{}

var Value Void

const UserProfileTags = "userProfileTags"

var UserSexCodeMap = map[string]string{
	"1": "010qww00",
	"2": "010qwq00",
}

var UserAgeCodeMap = map[string]string{
	"1": "010qqw00",
	"2": "010qqe00",
	"3": "010qqt00",
	"4": "010qqy00",
}

var DeviceSysVersionTagMap = map[string]string{
	"4":  "android4",
	"5":  "android5",
	"6":  "android6",
	"7":  "android7",
	"8":  "android8",
	"9":  "android9",
	"10": "android10",
}

var NetWorkTagMap = map[string]int{
	"unknown": 0,
	"wifi":    1,
	"2g":      2,
	"3g":      3,
	"4g":      4,
	"5g":      5,
}

var DeviceBrandTagMap = map[string]Void{
	"vivo":    Value,
	"oppo":    Value,
	"huawei":  Value,
	"honor":   Value,
	"xiaomi":  Value,
	"gionee":  Value,
	"samsung": Value,
	"meizu":   Value,
	"letv":    Value,
	"meitu":   Value,
	"zte":     Value,
	"nubia":   Value,
	"hisense": Value,
	"oneplus": Value,
}

type AdTagCode struct {
	NormalTag   []string
	TradeTag    []string
	GameTag     map[string][]int16
	TtTag       []string
	BusinessTag map[string][]int16
}

type UserProfileTagIndex struct {
	Pid            int64
	NotLimitTagSet *bitset.BitSet
	SubTagSet      map[string]*bitset.BitSet
}

type UserProfileTagVo struct {
	Id       int64
	TagCode  string
	TagName  string
	ParentId *int64
	UserNum  int64
	Type     int32
}

type AdTradeVO struct {
	Id                  int
	TradeCode           string
	TradeName           string
	GetuiCode           string
	ParentId            int
	UserNum             int64
	WeightLowUserNum    int64
	WeightMediumUserNum int64
	WeightHighUserNum   int64
}

type GameTagVO struct {
	Id               int
	TradeCode        string
	TradeName        string
	ParentId         int
	TagLevelUserNums []*TagLevelUserNum
}

type TagLevelUserNum struct {
	Level   int16
	UserNum int64
}

type AdsAllTagsRespData struct {
	UserProfileTags   []*UserProfileTagVo
	AdTradeTrees      []*AdTradeVO
	GameInterestTags  []*GameTagVO
	BusinessTags      []*AdTradeVO
	AllUserNum        int64
	AllGameTagUserNum int64
}

type AdsAllTagsResp struct {
	Code int                 `json:"return_code"`
	Msg  string              `json:"return_msg"`
	Data *AdsAllTagsRespData `json:"data"`
}

func (index *UserProfileTagIndex) MatchTag(tag string) bitset.BitSet {
	noLimit := index.NotLimitTagSet.Clone()
	if bs, ok := index.SubTagSet[tag]; ok {
		noLimit.InPlaceUnion(bs)
	}
	return *noLimit
}
