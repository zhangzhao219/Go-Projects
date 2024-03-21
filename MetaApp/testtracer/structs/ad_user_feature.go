package structs

type RecentResult struct {
	ClickList      []string
	ClickTimeList  []string
	PlayList       []string
	PlayTimeList   []string
	PlayListT1     []string
	PlayTimeListT1 []string
	PlayIdx        map[string]int64
}

var (
	UserFeatureCacheMapping map[string]string = map[string]string{
		RecentClick30:      "ClickList",
		RecentClick10:      "ClickList",
		RecentPlay30:       "PlayList",
		RecentPlay10:       "PlayList",
		RecentPlay1D50:     "PlayListT1",
		RecentClick30Time:  "ClickTimeList",
		RecentClick10Time:  "ClickTimeList",
		RecentPlay30Time:   "PlayTimeList",
		RecentPlay10Time:   "PlayTimeList",
		RecentPlay1D50Time: "PlayTimeListT1",
		PlayIdx:            "PlayIdx",
	}
)

const (
	RecentClick30      = "recent_click_30"
	RecentClick10      = "recent_click_10"
	RecentPlay10       = "recent_play_10"
	RecentPlay30       = "recent_play_30"
	RecentPlay1D50     = "recent_play1d_50"
	RecentClick30Time  = "recent_click_30_time"
	RecentClick10Time  = "recent_click_10_time"
	RecentPlay30Time   = "recent_play_30_time"
	RecentPlay10Time   = "recent_play_10_time"
	RecentPlay1D50Time = "recent_play1d_50_time"
	PlayIdx            = "play_idx"
)
