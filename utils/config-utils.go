package utils

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Validable represents any
type Validable interface {
	Validate() error
}

// InitConfiguration initializes a given configuration by using viper framework
func InitConfiguration(configuration Validable, cfgName, prefix, cfgFile string, defaults map[string]string) {

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for k, v := range defaults {
		viper.SetDefault(k, v)
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(cfgName)
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file not found but the config could be still valid because
			// all its values were provided in another way, eg environment variables or defauls
			log.Fatal("Error getting config file: ", err)
		}
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	if err := configuration.Validate(); err != nil {
		log.Fatal("Invalid configuration file: ", err)
	}
}
