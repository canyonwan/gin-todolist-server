package setting

import (
	"canyonwan.top/gin_todolist_server/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitMysql() (err error) {
	c := model.MysqlConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetInt("mysql.port"),
		User:     viper.GetString("mysql.user"),
		Password: viper.GetInt("mysql.password"),
		Database: viper.GetString("mysql.database"),
	}
	dsn := fmt.Sprintf("%s:%v@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
	DB, err = gorm.Open("mysql", dsn)
	//defer CloseDB()
	if err != nil {
		panic(fmt.Errorf("连接数据库失败: %s \n", err))
		return
	}
	pingErr := DB.DB().Ping()
	if pingErr != nil {
		return pingErr
	}
	DB.AutoMigrate(&model.TodoItem{})
	return
}

func CloseDB() {
	err := DB.Close()
	if err != nil {
		return
	}
}

func GetDatabaseInstance() *gorm.DB {
	return DB
}
