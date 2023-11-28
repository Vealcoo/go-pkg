package cnfhelper

import "github.com/spf13/viper"

func SetConfig(path, name, fileType string) *viper.Viper {
	cnf := viper.New()
	cnf.AddConfigPath(path)
	cnf.SetConfigName(name)
	cnf.SetConfigType(fileType)
	cnf.AutomaticEnv()

	err := cnf.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return cnf
}
