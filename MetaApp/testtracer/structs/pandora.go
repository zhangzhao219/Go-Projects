package structs

type PandoraABConfigRequest struct {
	OnlyId          string `json:"onlyId"`
	ChannelId       string `json:"channelId"`
	SelfPackageName string `json:"selfPackageName"`
	AppVersionCode  int    `json:"appVersionCode"`
}

type PandoraABConfigResponse struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    map[string]*ABConfigJson `json:"data"`
}
