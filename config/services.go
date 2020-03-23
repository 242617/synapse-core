package config

type ServicesConfig struct {
	Vault *VaultConfig `yaml:"vault"`
}

type VaultConfig struct {
	Address string `yaml:"address"`
	Token   string `yaml:"token"`
	Path    string `yaml:"path"`
}
