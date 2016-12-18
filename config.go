package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Configuration holds configuration data
type Configuration struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type configurationCache struct {
	Configurations map[string]Configuration `json:"configurations"`
}

var cache configurationCache

// AddConfiguration adds a configuration
func AddConfiguration(config Configuration) error {

	if config.Key == "" || config.Value == nil {
		return errors.New("config.key or config.value cannot be nil or empty")
	}

	if cache.Configurations == nil {
		cache.Configurations = make(map[string]Configuration)

		cache.Configurations[config.Key] = config
		return nil
	}

	cache.Configurations[config.Key] = config

	return nil
}

// GetConfiguration gets a configuration by a specified key
func GetConfiguration(key string) (*Configuration, error) {
	config := cache.Configurations[key]

	if config.Key == "" {
		return nil, errors.New("no configuration found")
	}

	return &config, nil
}

// GetConfigurationValue gets a configuration value by a specified key
func GetConfigurationValue(key string) interface{} {
	config, _ := GetConfiguration(key)

	return config.Value
}

// SetConfigurationValue sets an existing configurations value
func SetConfigurationValue(key string, value interface{}) error {
	config, err := GetConfiguration(key)

	if err != nil {
		return err
	}

	config.Value = value

	return nil
}

// Config Functions

// ValueString returns the value of the specified configuration as a string
func (config Configuration) ValueString() string {
	return config.Value.(string)
}

// ValueInt returns the value of the specified configuration as an int
func (config Configuration) ValueInt() int {
	return config.Value.(int)
}

// ValueBool returns the value of the specified configuration as a bool
func (config Configuration) ValueBool() bool {
	return config.Value.(bool)
}

// ValueFloat64 returns the value of the specified configuration as a float64
func (config Configuration) ValueFloat64() float64 {
	return config.Value.(float64)
}

// NewConfiguration creates a new configuration with the specified key and value
func NewConfiguration(key string, value interface{}) *Configuration {
	return &Configuration{key, value}
}

// Save saves the current configuration to the specified filepath, saves to the current working directory if no path is specified
func Save(filename string, path string) error {

	if _, err := os.Stat(path); err != nil {

		if os.IsNotExist(err) {
			// create the file
			os.Mkdir(path, os.ModeDir)
		}
	}

	b, err := json.Marshal(cache)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath.Join(path, filename), b, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Load loads the cache from the specified file
func Load(filepath string) error {

	rawData, err := ioutil.ReadFile(filepath)

	if err != nil {
		return err
	}

	json.Unmarshal(rawData, &cache)

	return nil
}
