package config

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Redis struct {
	Address string `yaml:"address"`
	DB      int    `yaml:"db"`
}

type Parameters struct {
	N int `yaml:"N"`
	M int `yaml:"M"`
	K int `yaml:"K"`
}
