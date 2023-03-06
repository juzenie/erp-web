package config

type DbInfo struct {
	Host     string `mapstructure:"host" json:"host"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Dbname   string `mapstructure:"dbname" json:"dbname"`
	Port     int    `mapstructure:"port" json:"port"`
}
type Postgres struct {
	Postgres DbInfo `mapstructure:"postgres" json:"postgres"`
}

type Mysql struct {
	Mysql DbInfo `mapstructure:"mysql" json:"mysql"`
}
