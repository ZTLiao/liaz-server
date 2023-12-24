package config

type Security struct {
	Encrypt    bool     `yaml:"encrypt"`
	SignKey    string   `yaml:"signKey"`
	PublicKey  string   `yaml:"publicKey"`
	PrivateKey string   `yaml:"privateKey"`
	Excludes   []string `yaml:"excludes"`
	Authorizes []string `yaml:"authorizes"`
}
