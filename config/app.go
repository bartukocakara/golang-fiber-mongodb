package config

import (
	"go-delivery-food/exception"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}

func Version() (urlPrefix string) {
	prefix := "api"
	version := "v1"
	urlPrefix = prefix+"/"+version

	return urlPrefix
}