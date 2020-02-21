package database

import (
	"database/sql"
	"io"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v1"
)

type Config struct {
	Datasource string `yaml:"datasource"`
}

type Configs map[string]*Config

func (cs Configs) OpenDB(env string) (*sql.DB, error) {
	config, ok := cs[env]
	if !ok {
		return nil, nil
	}
	return config.open()
}

func (c *Config) GetDatasource() string {
	return c.Datasource
}

func (c *Config) open() (*sql.DB, error) {
	return sql.Open("mysql", c.GetDatasource())
}

func NewConfigsFromFile(path string) (Configs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return newConfigs(f)
}

func newConfigs(r io.Reader) (Configs, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var configs Configs
	if err = yaml.Unmarshal(b, &configs); err != nil {
		return nil, err
	}
	return configs, nil
}
