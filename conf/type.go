package conf

type LoggerConfig struct {
	LogLevel     uint32   `mapstructure:"logLevel"`
	HookLevel    uint32   `mapstructure:"hookLevel"`
	ElasticHosts []string `mapstructure:"elasticHosts"`
	Index        string   `mapstructure:"index"`
}

type Addrs struct {
	Addr      string `mapstructure:"addr"`
	PProfAddr string `mapstructure:"pprofAddr"`
	RpcPort   int    `mapstructure:"rpcPort"`
}
