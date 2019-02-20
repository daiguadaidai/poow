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

var config Config

func NewConfig(fPath string) (*Config, error) {
	if _, err := toml.DecodeFile(fPath, &config); err != nil {
		return nil, err
	}

	config.SC.SupDefault()  // 补充服务配置文件默认值
	config.DBC.SupDefault() // 补充数据库配置文件默认值
	config.LC.SupDefault()  // 补充日志配置文件默认值

	return &config, nil
}
