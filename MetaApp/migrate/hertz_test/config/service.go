package config

type ServerConfig struct {
	ServiceConf ServiceConfig `json:"service_conf"`
	ListenConf  ListenConfig  `json:"listen_conf"`
}

type ServiceConfig struct {
	HTTPServiceName      string `json:"http_service_name"`
	RPCServiceName       string `json:"rpc_service_name"`
	ServiceProtectionQPS int    `json:"service_protection_qps"`
}

type ListenConfig struct {
	HttpAddr    string `json:"http_addr"`
	HttpPort    int    `json:"http_port"`
	RpcPort     int    `json:"rpc_port"`
	EnableAdmin bool   `json:"enable_admin"`
	AdminAddr   string `json:"admin_addr"`
	AdminPort   int    `json:"admin_port"`
}
