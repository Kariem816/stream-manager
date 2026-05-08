package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type configVar struct {
	Key         string
	DevOnly     bool
	ProdOnly    bool
	Transform   bool
	Transformer func(string, string) (any, error)
}

var transformers = map[string]func(string, string) (any, error){
	"toInt": func(name, val string) (any, error) {
		v, err := strconv.Atoi(val)
		if err != nil {
			return v, fmt.Errorf("ENV: field `%s` couldn't be converted to int. provided value: %s", name, val)
		}
		return v, err
	},
}

func setFieldByName(obj any, fieldName string, value any) error {
	v := reflect.ValueOf(obj)

	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected a pointer to a struct")
	}

	v = v.Elem()
	field := v.FieldByName(fieldName)

	if !field.IsValid() {
		return fmt.Errorf("no such field: %s in struct", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("cannot set field: %s", fieldName)
	}

	field.Set(reflect.ValueOf(value).Convert(field.Type()))

	return nil
}

func readConfig(f reflect.StructField) (configVar, error) {
	fName := f.Name
	tag, ok := f.Tag.Lookup("env")
	if !ok {
		return configVar{}, fmt.Errorf("ENV: field `%s` is not configured", fName)
	}

	pieces := strings.Split(tag, ";")
	if len(pieces) != 3 {
		return configVar{}, fmt.Errorf("ENV: field `%s` config should match the format `<Key>;<d|p|b>;[transformer]`", fName)
	}

	cv := configVar{}

	key := strings.TrimSpace(pieces[0])
	if len(key) == 0 {
		return configVar{}, fmt.Errorf("ENV: field `%s` key is not optional", fName)
	}
	cv.Key = key

	scope := strings.TrimSpace(pieces[1])
	if len(scope) != 1 {
		return configVar{}, fmt.Errorf("ENV: field `%s` scope should either be d|p|b (dev, prod, both)", fName)
	}
	switch scope {
	case "d":
		cv.DevOnly = true
	case "p":
		cv.ProdOnly = true
	case "b":
	default:
		return configVar{}, fmt.Errorf("ENV: field `%s` scope should either be d|p|b (dev, prod, both)", fName)
	}

	t := strings.TrimSpace(pieces[2])
	if len(t) == 0 {
		return cv, nil
	}

	cv.Transform = true

	ter, ok := transformers[t]
	if !ok {
		return configVar{}, fmt.Errorf("ENV: field `%s` transformer (%s) is not registered", fName, t)
	}

	cv.Transformer = ter

	return cv, nil
}

func validateConfig(env ENV, name string, cv configVar) (any, bool, error) {
	val, ok := os.LookupEnv(cv.Key)
	if !ok {
		if (env == ENV_DEV && !cv.ProdOnly) || (env == ENV_PROD && !cv.DevOnly) {
			return nil, false, fmt.Errorf("ENV: `%s` variable is required", name)
		}
		return "", false, nil
	}

	if cv.Transform {
		vot, err := cv.Transformer(name, val)
		return vot, true, err
	}

	return val, true, nil
}

var Cfg TConfig

func Init() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	var env ENV
	v, ok := os.LookupEnv("ENV")
	if !ok {
		return errors.New("ENV: `ENV` variable is required")
	}
	err = env.Scan(v)
	if err != nil {
		return err
	}

	typ := reflect.TypeOf(Cfg)
	for i := range typ.NumField() {
		f := typ.Field(i)
		name := f.Name

		if name == "Env" {
			continue
		}

		cv, err := readConfig(f)
		if err != nil {
			return err
		}

		val, ok, err := validateConfig(ENV_DEV, f.Name, cv)
		if err != nil {
			return err
		}

		if ok { // some variables might not be needed in some environments
			err = setFieldByName(&Cfg, name, val)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
