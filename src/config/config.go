package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Config = Schema{}

type Schema struct {
	Server struct {
		Port string
	}
	Database struct {
		Type     string
		Url      string
		Port     string
		Username string
		Password string
	}
	Security struct {
		Token struct {
			Auth  string
			Slate string
		}
	}
	Email struct {
		Smtp     string
		Host     string
		Account  string
		Password string
	}
	Redis struct {
		Url      string
		DB       int
		Password string
	}
}

func InitConfig(filePath string) Schema {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(content, &Config) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	return Config
}
