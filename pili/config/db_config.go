package config

import (
	"fmt"
	"strings"
)

const (
	DB_HOST           = "127.0.0.1"
	DB_PORT           = 3306
	DB_USERNAME       = "root"
	DB_PASSWORD       = "root"
	DB_SCHEMA         = ""
	DB_AUTO_COMMIT    = true
	DB_MAX_OPEN_CONNS = 100
	DB_MAX_IDEL_CONNS = 100
	DB_CHARSET        = "utf8mb4"
	DB_TIMEOUT        = 10
)

var dbConfig *DBConfig

type DBConfig struct {
	Username          string
	Password          string
	Database          string
	CharSet           string
	Host              string
	Timeout           int
	Port              int
	MaxOpenConns      int
	MaxIdelConns      int
	AllowOldPasswords int
	AutoCommit        bool
}

/* 新建一个数据库执行器
Params:
    _host: ip
    _port: 端口
    _username: 链接数据库用户名
    _password: 链接数据库密码
    _database: 要操作的数据库
*/
func NewDBConfig(
	_host string,
	_port int,
	_username string,
	_password string,
	_database string,
	_charset string,
	_autoCommit bool,
	_timeout int,
	_maxOpenConns int,
	_maxIdelConns int,
) *DBConfig {
	dbConfig := &DBConfig{
		Username:          _username,
		Password:          _password,
		Host:              _host,
		Port:              _port,
		Database:          _database,
		CharSet:           _charset,
		MaxOpenConns:      _maxOpenConns,
		MaxIdelConns:      _maxIdelConns,
		Timeout:           _timeout,
		AllowOldPasswords: 1,
		AutoCommit:        _autoCommit,
	}

	return dbConfig
}

func (this *DBConfig) GetDataSource() string {
	dataSource := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=%v&allowOldPasswords=%v&timeout=%vs&autocommit=%v&parseTime=True&loc=Local",
		this.Username,
		this.Password,
		this.Host,
		this.Port,
		this.Database,
		this.CharSet,
		this.AllowOldPasswords,
		this.Timeout,
		this.AutoCommit,
	)

	return dataSource
}

func (this *DBConfig) Check() error {
	if strings.TrimSpace(this.Database) == "" {
		return fmt.Errorf("数据库不能为空")
	}

	return nil
}

// 设置 DBConfig
func SetDBConfig(_dbConfig *DBConfig) {
	dbConfig = _dbConfig
}

func GetDBConfig() *DBConfig {
	return dbConfig
}
