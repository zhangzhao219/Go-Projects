package structs

type DeviceParam struct {
	OnlyId             string
	Imei               string
	Oaid               string
	Idfa               string
	AndroidId          string
	ImeiMd5            string
	OaidMd5            string
	IdfaMd5            string
	AndroidIdMd5       string
	DeviceName         string
	DeviceManufacturer string
	DeviceBrand        string
	DeviceProduct      string
	DeviceModel        string
	DeviceSys          string
	DeviceSysVersion   string
	DeviceOs           string
	BootMark           string
	UpdateMark         string
	ScreenWidth        int
	ScreenHeight       int
	ScreenDensity      float32
	ScreenDensityDpi   int
	FreeSpace          int64
	TotalSpace         int64
	AvailMem           int64
	TotalMem           int64
	InstalledApps      []*InstalledApp
	RunningApp         []*RunningApp
}

type InstalledApp struct {
	InstalledAppPackageName string
	InstalledAppVersionName string
	InstalledAppVersionCode string
}

type RunningApp struct {
	RunningAppPackageName string
	RunningAppVersionName string
	RunningAppVersionCode string
}

var DeviceRangeSysVersionMap = map[string]string{
	"4.4":  "4.4,4",
	"5.0":  "5.0,5",
	"6.0":  "6.0,6",
	"7.0":  "7.0,7",
	"8.0":  "8.0,8",
	"9.0":  "9.0,9",
	"10.0": "10.0,10",
	"11.0": "11.0,11",
	"12.0": "12.0,12",
	"未知":   "unknown",
	"全部":   "ALL",
}

var DeviceWhiteListMap = map[string]bool{
	"4":  true,
	"5":  true,
	"6":  true,
	"7":  true,
	"8":  true,
	"9":  true,
	"10": true,
	"11": true,
	"12": true,
	"13": true,
}
