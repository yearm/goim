package http

type Config struct {
	Addr      string `mapstructure:"addr"`
	PProfAddr string `mapstructure:"pprofAddr"`
}
