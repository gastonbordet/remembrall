package util

import "github.com/spf13/viper"

type Config struct {
	DB_USER    string `mapstructure:"DB_USER"`
	DB_PWD     string `mapstructure:DB_PWD`
	DB_ADDRESS string `mapstructure:DB_ADDRESS`
	DB_PORT    string `mapstructure:DB_PORT`
	DB_NAME    string `mapstructure:DB_NAME`
}

func LoadConfig(path string, filename string, fileType string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType(fileType)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	return config, nil
}
