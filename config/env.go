package config

import "errors"

type ENV int

const (
	ENV_DEV ENV = iota
	ENV_PROD
	ENV_UNSET
)

func (e ENV) String() string {
	return [...]string{"dev", "prod", "unset"}[e]
}

func (e *ENV) Scan(value string) error {
	if value == "" {
		*e = ENV_UNSET
		return nil
	}

	switch value {
	case "dev":
		*e = ENV_DEV
	case "prod":
		*e = ENV_PROD
	default:
		return errors.New("ENV: `ENV` variable should be either dev|prod")
	}
	return nil
}
