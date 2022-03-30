package config

import "sync"

var once sync.Once
var config *Config

type Config struct {
	App App
	Job Job
}

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{
			App: NewApp(),
			Job: NewJob(),
		}
	})

	return config
}
