package env

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Values struct {
	DB
	Env   string `required:"true" split_words:"true"`
	Debug bool   `default:"false" split_words:"true"`
	Port  string `default:"8080" split_words:"true"`
}

type DB struct {
	PostgresHost     string `required:"true" split_words:"true"`
	PostgresPort     string `default:"3306" split_words:"true"`
	PostgresDatabase string `default:"go_graphql_grpc" split_words:"true"`
	PostgresUsername string `required:"true" split_words:"true"`
	PostgresPassword string `required:"true" split_words:"true"`
}

func NewValue() (*Values, error) {
	var v Values
	err := envconfig.Process("app", &v)
	if err != nil {
		s := fmt.Sprintf("need to set all env values %+v", v)
		return nil, errors.Wrap(err, s)
	}
	return &v, nil
}
