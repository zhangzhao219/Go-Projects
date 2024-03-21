package structs

import "github.com/bits-and-blooms/bitset"

type DspFilterData struct {
	BadCaseIdSet      map[int]bool
	DeadIdSet         map[int]bool
	PtrMap            map[int]float64
	UserAdShowMap     map[string]int
	FilledAdIn30Min   map[string]int
	CpmCutOffAdIds    map[string]bool
	TodayStartTime    int64
	TodayEndTime      int64
	AdCostStatMap     map[int]*AdCostStatDO
	UserBehavior      *UserBehaviorData
	UidCostMap        map[string]float64
	GroupCostMap      map[string]float64
	PlanCostMap       map[string]float64
	AreaCodeInt       uint
	CrowdApIdSet      map[int]bool
	CrowdPackOriRes   map[string]bool
	NormalTagMatchAds *bitset.BitSet
	NewUidSet         map[string]bool
}

type UserBehaviorData struct {
	UserRecentPlayGamePkgs   map[string]bool
	UserRecentPlayGamePkgsT1 map[string]bool
	UserInstalledAppList     map[string]bool
}
