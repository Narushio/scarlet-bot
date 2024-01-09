package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppID uint64 `required:"true" split_words:"true"`
	Token string `required:"true" split_words:"true"`
}

var FromEnv = &Config{}

func init() {
	envconfig.MustProcess("", FromEnv)
}
