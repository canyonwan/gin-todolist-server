package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper() (err error) {
	viper.SetConfigFile("./config.toml") // 指定配置文件路径
	viper.SetConfigName("config")        // 指定配置文件名
	viper.SetConfigType("toml")          // 指定配置文件类型
	viper.AddConfigPath("./config/")     // 搜索路径

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("viper读取配置文件失败: %s \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("viper配置文件已经修改了!")
	})
	return
}
