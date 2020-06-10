package etcd

type Config struct {
	Hosts []string `mapstructure:"hosts"`
	Dir   string   `mapstructure:"dir"`
}
