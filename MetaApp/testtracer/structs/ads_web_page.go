package structs

type WebPageResourceResp struct {
	Code int32            `json:"code"`
	Msg  string           `json:"msg"`
	Data *WebPageResource `json:"data"`
}

type WebPageResource struct {
	Icon             string `json:"icon"`
	Name             string `json:"name"`
	VideoUrl         string `json:"videoUrl"`
	ImageUrl         string `json:"imageUrl"`
	Intro            string `json:"intro"`
	DownloadLink     string `json:"downloadLink"`
	PkgName          string `json:"pkgName"`
	PkgMd5           string `json:"pkgMd5"`
	DeveloperName    string `json:"developerName"`
	AppPermissions   string `json:"appPermissions"`
	PrivacyAgreement string `json:"privacyAgreement"`
}

type GameRecWebPageResource struct {
	Aiid        int64  `json:"id,omitempty"`
	Pkg         string `json:"pkg,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}
