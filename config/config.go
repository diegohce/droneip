package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	Values = configVars(map[string]string{})
)

type configVars map[string]string

func NewValues() configVars {
	return configVars(map[string]string{})
}

func (c configVars) FromEnv() {
	c.FromEnvWithPrefix("")
}

func (c configVars) FromEnvWithPrefix(pref string, must ...string) error {
	for _, v := range os.Environ() {
		if strings.HasPrefix(v, pref) {
			kv := strings.SplitN(v, "=", 2)
			key := kv[0][len(pref):]
			c[key] = kv[1]
		}
	}

	for _, m := range must {
		if v, exists := c[m]; !exists || len(v) == 0 {
			return fmt.Errorf("config: key %s is mandatory", m)
		}
	}

	return nil
}

func (c configVars) Get(name string, def ...string) string {
	default_value := ""
	if len(def) > 0 {
		default_value = def[0]
	}

	value, ok := c[name]
	if !ok {
		return default_value
	}
	return value
}

func (c configVars) GetInt(name string, def ...int) int {
	default_value := 0
	if len(def) > 0 {
		default_value = def[0]
	}

	v := c.Get(name)
	if v == "" {
		return default_value
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return default_value
	}
	return i
}

func (c configVars) GetDuration(name string, def ...string) time.Duration {
	default_value := "0s"
	if len(def) > 0 {
		default_value = def[0]
	}

	dd, err := time.ParseDuration(default_value)
	if err != nil {
		return time.Duration(0)
	}

	v := c.Get(name)
	if v == "" {
		return dd
	}

	d, err := time.ParseDuration(v)
	if err != nil {
		return dd
	}
	return d

}

func (c configVars) Set(name string, value string) {
	c[name] = value
}

func Get(name string, def ...string) string {
	return Values.Get(name, def...)
}

func GetInt(name string, def ...int) int {
	return Values.GetInt(name, def...)
}

func GetDuration(name string, def ...string) time.Duration {
	return Values.GetDuration(name, def...)
}

func Set(name string, value string) {
	Values.Set(name, value)
}

func FromEnv() {
	Values.FromEnv()
}

func FromEnvWithPrefix(pref string, must ...string) error {
	return Values.FromEnvWithPrefix(pref, must...)
}
