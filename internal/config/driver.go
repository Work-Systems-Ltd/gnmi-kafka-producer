package config

import "fmt"

// Driver is the on-disk shape of configs/driver.yaml.
type Driver struct {
	GNMI  GNMI     `yaml:"gnmi"`
	Hosts []string `yaml:"hosts"`
	Flap  Flap     `yaml:"flap"`
}

func LoadDriver(path string) (*Driver, error) {
	var c Driver
	if err := loadYAML(path, &c); err != nil {
		return nil, err
	}
	c.GNMI.applyDefaults()
	c.Flap.applyDefaults()
	if err := c.validate(); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Driver) validate() error {
	if len(c.Hosts) == 0 {
		return fmt.Errorf("hosts must have at least one entry")
	}
	if c.Flap.Enabled && len(c.Flap.Interfaces) == 0 {
		return fmt.Errorf("flap.enabled=true requires at least one entry in flap.interfaces")
	}
	return nil
}
