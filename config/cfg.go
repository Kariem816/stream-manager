package config

type TConfig struct {
	Env ENV // tags are not required for this special variables

	SharedSecretKey string `env:"SHARED_SECRET_KEY;b;"`
}
