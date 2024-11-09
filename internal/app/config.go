package app

import (
	"fmt"
	"github.com/spf13/viper"
	"myAccount/internal/utils"
	"path/filepath"
)

type Config struct {
	APP struct {
		Name string `yaml:"name"`
	}
	DB struct {
		Type string `yaml:"type"`
		File string `yaml:"file"`
	}
	WEB struct {
		Port string `yaml:"port"`
	}
}

// GetConfig 根据路径解析配置
func GetConfig(configFilePath string) (*Config, error) {
	if configFilePath == "" {
		return nil, fmt.Errorf("config file path is empty")
	}
	if !utils.IsExist(configFilePath) {
		abs, _ := filepath.Abs(configFilePath)
		return nil, fmt.Errorf("config file not exist: %s", abs)
	}
	config := &Config{}
	path, name, ext := utils.SplitFilePath(configFilePath)
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	if "" == ext {
		viper.SetConfigType(ext)
	}
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}
	return config, nil
}
