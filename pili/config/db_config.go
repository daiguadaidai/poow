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
	Username          string `toml:"username"`
	Password          string `toml:"password"`
	Database          string `toml:"database"`
	Charset           string `toml:"charset"`
	Host              string `toml:"host"`
	Timeout           int    `toml:"timeout"`
	Port              int    `toml:"port"`
	MaxOpenConns      int    `toml:"max_open_conns"`
	MaxIdelConns      int    `toml:"max_idel_conns"`
	AllowOldPasswords int    `toml:"allow_old_passwords"`
	AutoCommit        bool   `toml:"auto_commit"`
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
		Charset:           _charset,
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
		this.Charset,
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

// 补充默认值
func (this *DBConfig) SupDefault() {
	if len(this.Username) == 0 {
		this.Username = DB_USERNAME
	}
	if len(this.Password) == 0 {
		this.Password = DB_PASSWORD
	}
	if len(this.Charset) == 0 {
		this.Charset = DB_CHARSET
	}
	if len(this.Host) == 0 {
		this.Host = DB_HOST
	}
	if this.Port < 1 {
		this.Port = DB_PORT
	}
	if this.Timeout < 0 {
		this.Timeout = DB_TIMEOUT
	}
	if this.MaxOpenConns < 1 {
		this.MaxOpenConns = DB_MAX_OPEN_CONNS
	}
	if this.MaxIdelConns < 1 {
		this.MaxIdelConns = DB_MAX_IDEL_CONNS
	}
	this.AllowOldPasswords = 1
	this.AutoCommit = DB_AUTO_COMMIT
}
