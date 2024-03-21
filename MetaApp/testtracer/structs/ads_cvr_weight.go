package structs

type CvrWeight struct {
	Dimension      string
	DimensionValue string
	Priority       int
	CpcWeight      float64
	CpaWeight      float64
	CpiWeight      float64
	CpmWeight      float64
	MaxShow        int
	Operator       string
	OperateTime    string
	Remark         string
}

var CvrWeightDimension = []string{
	"apid", "sdkVersion", "default",
}

const (
	DefaultCvrWeightDimension        = "default"
	ApIdCvrWeightDimension           = "apid"
	SdkVersionCodeCvrWeightDimension = "sdkVersion"
)
