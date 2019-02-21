package config

import "github.com/BurntSushi/toml"

const (
	CONFIG_FILE_PATH = "./pili.toml"
)

type Config struct {
	SC  *ServerConfig `toml:"server"`
	DBC *DBConfig     `toml:"database"`
	LC  *LogConfig    `toml:"log"`
}

var cfg Config

func NewConfig(fPath string) (*Config, error) {
	if _, err := toml.DecodeFile(fPath, &cfg); err != nil {
		return nil, err
	}

	cfg.SC.SupDefault()  // 补充服务配置文件默认值
	cfg.DBC.SupDefault() // 补充数据库配置文件默认值
	cfg.LC.SupDefault()  // 补充日志配置文件默认值

	return &cfg, nil
}
