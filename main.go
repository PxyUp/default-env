package default_env

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type configMap struct {
	Variables map[string]string `yaml:"env_var"`
}

type DefaultEnvInstance struct {
	Get func(string) string
}

var once sync.Once
var cfg configMap
var instance DefaultEnvInstance

func get(env string) string {
	if len(os.Getenv(env)) == 0 {
		if cfg.Variables[env] == "" {
			log.Printf("Variables '%v' not set in file", env)
		}
		return cfg.Variables[env]
	}
	return os.Getenv(env)
}

func initInstance(path string) DefaultEnvInstance {

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err '%v' ", err)
	}
	err = yaml.Unmarshal(yamlFile, &cfg)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	inst := DefaultEnvInstance{
		Get: func(env string) string {
			return get(env)
		},
	}
	return inst
}

func defaultEnvInstanceFactory(path string) DefaultEnvInstance {
	return initInstance(path)
}

func GetInstance(path string) DefaultEnvInstance {
	once.Do(func() {
		instance = defaultEnvInstanceFactory(path)
	})
	return instance
}
