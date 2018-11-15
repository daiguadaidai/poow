package dao

import "github.com/daiguadaidai/poow/pili/config"

func InitDBConfig() {
	dbConfig := &config.DBConfig{
		Host:         "10.10.10.21",
		Port:         3307,
		Username:     "HH",
		Password:     "oracle12",
		Database:     "poow",
		CharSet:      "utf8mb4",
		AutoCommit:   true,
		MaxOpenConns: 100,
		MaxIdelConns: 100,
		Timeout:      10,
	}

	config.SetDBConfig(dbConfig)
}
