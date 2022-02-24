package model

// MysqlConfig 数据库配置
type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password int    `json:"password"`
	Database string `json:"database"`
}
