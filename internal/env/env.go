package env

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Values struct {
	DB
	Redis
	Env        string `default:"local" split_words:"true"`
	Debug      bool   `default:"true" split_words:"true"`
	ServerPort string `required:"true" split_words:"true"`
}

type DB struct {
	PostgresHost     string `default:"postgres" split_words:"true"`
	PostgresPort     string `default:"5432" split_words:"true"`
	PostgresDatabase string `default:"go_pg" split_words:"true"`
	PostgresUser     string `required:"true" split_words:"true"`
	PostgresPassword string `required:"true" split_words:"true"`
}

type Redis struct {
	RedisHost string `default:"localhost" split_words:"true"`
	RedisPort string `default:"6379" split_words:"true"`
}

func NewValue() (*Values, error) {
	var v Values
	err := envconfig.Process("star", &v)
	if err != nil {
		s := fmt.Sprintf("need to set all env values %+v", v)
		return nil, errors.Wrap(err, s)
	}
	return &v, nil
}
