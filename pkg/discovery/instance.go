package discovery

type Instance struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
	Tags map[string]string `json:"tags"`
}