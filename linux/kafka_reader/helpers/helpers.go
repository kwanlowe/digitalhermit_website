package helpers

import (
  "fmt"
  toml "github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Title string
        KafkaConfig kafkaConfig `toml:"kafka"`
}

type kafkaConfig struct {
       TopicName string
       BootstrapServers string
       AutoOffsetReset string
       EnableAutoCommit bool
       ConsumerTimeoutMs int
       GroupID string
}

func GetConfig(configFile string) *tomlConfig {
    config := new(tomlConfig)
    if _, err := toml.DecodeFile(configFile, &config); err != nil {
         fmt.Println(err)
         return config
    }
    return config
}

