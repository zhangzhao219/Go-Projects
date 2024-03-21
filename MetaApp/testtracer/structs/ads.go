package structs

var DeliveryTypeScene = map[string]string{
	"10_4_0": HomeIcon,
	"10_4_1": HotSearch,
	"10_4_2": DetailsPage,
	"10_4_3": Search,
}

const (
	RedisBatchSize = 1000
	Scene          = "%d_%d"
)

const (
	HomeIcon    = "home_icon"
	HotSearch   = "hot_search"
	Search      = "search"
	DetailsPage = "details_page"
)

const (
	//首页场景appType_deliveryType_businessType
	ATypeDTypeBTypeHomeIcon = "10_4_0"
)

const FormatDate = "2006-01-02"
