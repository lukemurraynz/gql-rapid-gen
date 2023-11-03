package state

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	OutputDirectory string
	TagEnable       []string `json:"TagEnable,omitempty"`
	PluginEnable    []string `json:"PluginEnable,omitempty"`
	PluginDisable   []string `json:"PluginDisable,omitempty"`
	SchemaFiles     []string
}

func (c *Config) Validate() (err error) {
	if c.OutputDirectory == "" {
		return fmt.Errorf("OutputDirectory must be specified")
	}
	if len(c.PluginEnable) > 0 && len(c.PluginDisable) > 0 {
		return fmt.Errorf("PluginEnable and PluginDisable may not both be specified")
	}
	if len(c.SchemaFiles) == 0 {
		return fmt.Errorf("SchemaFiles must have at least one entry")
	}
	return nil
}

func (c *Config) Save(path string) (err error) {
	raw, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return fmt.Errorf("failed marshalling config: %w", err)
	}
	err = os.WriteFile(path, raw, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed writing config file: %w", err)
	}
	return nil
}

func LoadConfig(path string) (c *Config, err error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed reading config file: %w", err)
	}

	c = &Config{}
	err = json.Unmarshal(raw, c)
	if err != nil {
		return nil, fmt.Errorf("failed parsing config file: %w", err)
	}
	return c, nil
}
