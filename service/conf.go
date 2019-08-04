package service

import "github.com/spf13/viper"

const (
	path string = "./server.yaml"
)

var cachedConf *viper.Viper = nil

func Conf() (*viper.Viper, error) {
	if cachedConf == nil {
		v := viper.New()
		v.SetConfigFile(path)

		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}
		cachedConf = v
	}
	return cachedConf, nil
}
