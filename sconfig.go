package sconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

// SConfig read config interface
type SConfig interface {
	LoadConfig(configFile string) error
	Get(key string) (any, bool)
	GetString(key string) (string, bool)
	GetBool(key string) (bool, bool)
	GetInt(key string) (int, bool)
	GetUint(key string) (uint, bool)
	GetFloat64(key string) (float64, bool)
	GetStringSlice(key string) ([]string, bool)
	AllSettings() map[string]any
	UnmarshalKey(key string, rawVal any) error
}

// config read config struct
type config struct {
	Data map[string]any
}

// LoadConfig load config file, call it once before call other method.
func (c *config) LoadConfig(configFile string) error {
	content, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	configData := make(map[string]any)
	if err := yaml.Unmarshal(content, configData); err != nil {
		return err
	}
	c.Data = configData
	return nil
}

func (c *config) find(cfg any, key string) (any, bool) {
	parts := strings.Split(key, ".")
	for _, item := range parts {
		switch c := cfg.(type) {
		case map[string]any:
			if value, ok := c[item]; ok {
				cfg = value
			} else {
				return nil, false
			}
		default:
			return nil, false
		}
	}
	return cfg, true
}

// Get get config by key.
func (c *config) Get(key string) (any, bool) {
	if c.Data == nil {
		return nil, false
	}
	if key == "" {
		return nil, false
	}
	return c.find(c.Data, key)
}

// GetString get string by key
func (c *config) GetString(key string) (string, bool) {
	data, ok := c.Get(key)
	if !ok {
		return "", false
	}
	result, ok := data.(string)
	return result, ok
}

// GetBool get bool by key
func (c *config) GetBool(key string) (bool, bool) {
	data, ok := c.Get(key)
	if !ok {
		return false, false
	}
	result, ok := data.(bool)
	return result, ok
}

// GetInt get int by key
func (c *config) GetInt(key string) (int, bool) {
	data, ok := c.Get(key)
	if !ok {
		return 0, false
	}
	result, ok := data.(int)
	return result, ok
}

// GetUint get uint by key
func (c *config) GetUint(key string) (uint, bool) {
	data, ok := c.Get(key)
	if !ok {
		return 0, false
	}
	result, ok := data.(uint)
	return result, ok
}

// GetFloat64 get float64 by key
func (c *config) GetFloat64(key string) (float64, bool) {
	data, ok := c.Get(key)
	if !ok {
		return 0, false
	}
	result, ok := data.(float64)
	return result, ok
}

// GetStringSlice get string[] by key
func (c *config) GetStringSlice(key string) ([]string, bool) {
	data, ok := c.Get(key)
	if !ok {
		return nil, false
	}
	list, ok := data.([]any)
	if !ok {
		return nil, false
	}
	result := make([]string, len(list))
	for index, item := range list {
		result[index], ok = item.(string)
		if !ok {
			return nil, false
		}
	}
	return result, ok
}

// AllSettings get all settings by key
func (c *config) AllSettings() map[string]any {
	return c.Data
}

// UnmarshalKey get special type config item by key
func (c *config) UnmarshalKey(key string, rawVal any) error {
	data, ok := c.Get(key)
	if !ok {
		return fmt.Errorf("key: %s not exist", key)
	}
	return mapstructure.Decode(data, rawVal)
}

// New SConfig instance.
func New() SConfig {
	return &config{}
}
