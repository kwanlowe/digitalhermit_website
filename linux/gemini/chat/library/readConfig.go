package readConfig

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/spf13/viper"
)

func LoadConfig(path,appname string) (*viper.Viper, error) {

	homeDir, err := os.UserHomeDir()

	// Create a new Viper instance
	v := viper.New()
	v.SetConfigName(path)
	v.SetConfigType("toml")
	v.AddConfigPath(".") 
	v.AddConfigPath(homeDir)
	v.AddConfigPath(filepath.Join(homeDir, ".config", appname))

	// Read the configuration file
	err = v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No configuration file found. Using defaults or environment variables.")
		} else {
			// Config file was found but another error occurred
			log.Fatalf("Error reading config file: %s", err)
		}
	} else {
		fmt.Println("Successfully read configuration file:", v.ConfigFileUsed())
	}
	return v, err
}

