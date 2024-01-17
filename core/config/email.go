package config

type Email struct {
	Host     string            `yaml:"host"`
	Port     int               `yaml:"port"`
	Nickname string            `yaml:"nickname"`
	Username string            `yaml:"username"`
	Password string            `yaml:"password"`
	Subject  map[string]string `yaml:"subject"`
	Template map[string]string `yaml:"template"`
}
