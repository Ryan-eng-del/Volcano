package config



var BaseMapConfInstance = &BaseMapConf{}
var NacoMapConfInstance = &NacosMapConf{}
var ServerConfInstance = &ServerConfig{}

type BaseMapConf struct {
  Base BaseConf `mapstructure:"base"`
}

type BaseConf struct {
	Addr string `mapstructure:"addr"`
	Port int `mapstructure:"port"`
}

type NacosMapConf struct {
  Base NacosConf `mapstructure:"base"`
}

type NacosConf struct {
	Host      string `mapstructure:"host"`
	Port      uint64    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
	LogDir  string `mapstructure:"log_dir"`
	CacheDir string `mapstructure:"cache_dir"`
	MaxAge int `mapstructure:"max_age"`
	LogLevel string `mapstructure:"log_level"`
	RotateTime string `mapstructure:"rotate_time"`
	Timeout uint64 `mapstructure:"time_out"`
	NotLoadCacheAtStart bool `mapstructure:"not_load_cache_at_start"`
	MaxBackUp int `mapstructure:"max_backup"`
}


type MysqlConfig struct{
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"db" json:"db"`
	User string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	MaxOpenConn int `mapstructure:"max_open_conn" json:"max_open_conn"`
	MaxIdleConn int `mapstructure:"max_idle_conn" json:"max_idle_conn"`
	MaxCoonLifeTime int `mapstructure:"max_conn_life_time" json:"max_conn_life_time"`
	TimeLocation string `mapstructure:"location" json:"location"`
}

type ConsulConfig struct{
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Timeout string `mapstructure:"timeout" json:"timeout"`
	Interval string `mapstructure:"interval" json:"interval"`
	DeregisterCriticalServiceAfter string `mapstructure:"deregister_critical_service_after" json:"deregister_critical_service_after"`	
	Tags []string `mapstructure:"tags" json:"tags"`

}

type ServerConfig struct{
	Name string `mapstructure:"name" json:"name"`
	MysqlInfo MysqlConfig `mapstructure:"mysql" json:"mysql"`
	ConsulInfo ConsulConfig `mapstructure:"consul" json:"consul"`
	ZapInfo ZapConf `mapstructure:"zap" json:"zap"`
}


type ZapConf struct {
	MaxSize int `mapstructure:"max_size" json:"max_size"`
	MaxAge int `mapstructure:"max_age" json:"max_age"`
	MaxBackups int `mapstructure:"max_backups" json:"max_backups"`
	DebugFileName string `mapstructure:"debug_file_name" json:"debug_file_name"`
	InfoFileName string `mapstructure:"info_file_name" json:"info_file_name"`
	ErrorFileName string `mapstructure:"error_file_name" json:"error_file_name"`
}